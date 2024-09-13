package authService

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/constant"
	"go-shift/cmd/app/domain/dao/table"
	"go-shift/cmd/app/domain/dto"
	"go-shift/cmd/app/domain/dto/system"
	"go-shift/cmd/app/repository"
	"go-shift/cmd/app/service"
	apiService "go-shift/cmd/app/service/auth/api_service"
	"go-shift/cmd/app/util"
	"go-shift/pkg"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	googleService     *GoogleOauthServiceImpl
	googleServiceOnce sync.Once
)

type GoogleOauthServiceImpl struct {
	mailService     service.MailService
	redisService    service.RedisService
	oauthApiService apiService.OauthApiService

	mailRepository        repository.MailRepository
	userRepository        repository.UserRepository
	userProfileRepository repository.UserProfileRepository
	authTokenRepository   repository.AuthTokenRepository
}

func (svc *GoogleOauthServiceImpl) SignIn(c *gin.Context) {
	redirectUri := os.Getenv("GOOGLE_SIGN_IN_REDIRECT_URI")
	token, _ := svc.oauthProcess(c, "sign-in", redirectUri)

	//data := system.ApiResponse[any]{
	//	ResponseKey:     constant.Success.GetResponseStatus(),
	//	ResponseMessage: constant.Success.GetResponseMessage(),
	//	Data:            nil,
	//}

	// Define the callback URL with the query parameters
	callbackUrl := "http://localhost:5173/oauth/google/sign-in-callback"
	callbackType := "sign-in-google"

	// Construct the redirect URL with query parameters
	redirectUrl := fmt.Sprintf("%s?callbackType=%s&token=%s", callbackUrl, callbackType, token)

	c.Redirect(http.StatusFound, redirectUrl)
}

func (svc *GoogleOauthServiceImpl) SignInCallback(c *gin.Context) {
	log.Info("execute sign in callback")
	svc.SignIn(c)
}

func (svc *GoogleOauthServiceImpl) SignUp(c *gin.Context) {
	redirectUri := os.Getenv("GOOGLE_SIGN_UP_REDIRECT_URI")
	token, payload := svc.oauthProcess(c, "sign-up", redirectUri)

	if err := svc.mailService.SendMail("Action Required: Verify Your Email to Activate Your Shift Account",
		payload); err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	data := system.ApiResponse[any]{
		ResponseKey:     constant.Success.GetResponseStatus(),
		ResponseMessage: constant.Success.GetResponseMessage(),
		Data:            nil,
	}

	c.Header("Authorization-Token", token)

	c.JSON(200, data)
}

func (svc *GoogleOauthServiceImpl) SignUpCallback(c *gin.Context) {
	svc.SignUp(c)
}

func (svc *GoogleOauthServiceImpl) oauthProcess(c *gin.Context, processType string, redirectUri string) (string, dto.JWTClaimsPayloadGoogle) {
	defer pkg.PanicHandler(c)

	code := c.Query("code")

	clientId := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	getTokenUrl := os.Getenv("GOOGLE_ACCESS_TOKEN_URL")

	token, err := svc.oauthApiService.GetAccessToken(code, clientId, clientSecret, getTokenUrl, redirectUri)
	if err != nil {
		log.Error("Error when get access token :: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	payload, err := util.GetTokenPayload(dto.JWTClaimsPayloadGoogle{}, token.IdToken)

	if err != nil {
		log.Error("Error when try decode token")
		pkg.PanicException(constant.UnknownError)
	}

	var userAccount table.UserAccount
	userAccount, err = svc.userRepository.FindUserByEmail(payload.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) && processType == "sign-up" {
			userAccount, err = svc.userRepository.SaveInitiateUser(payload.Email, 1)
			if err != nil {
				pkg.PanicException(constant.UnknownError)
			}
		} else if errors.Is(err, gorm.ErrRecordNotFound) && processType == "sign-in" {
			pkg.PanicException(constant.DataNotFound)
		} else {
			pkg.PanicException(constant.UnknownError)
		}
	}

	authToken := &table.AuthToken{
		UserAccountID: userAccount.ID,
		AccessToken:   token.AccessToken,
		RefreshToken:  token.RefreshToken,
		ExpiresIn:     time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
		IsActive:      true,
	}

	_, err = svc.authTokenRepository.SaveUserAuth(authToken)

	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	accessToken := svc.generateToken(c, &userAccount, token, payload)

	return accessToken, *payload
}

func (svc *GoogleOauthServiceImpl) VerifyEmail(c *gin.Context) {
	defer pkg.PanicHandler(c)
	token := c.Query("verificationToken")
	verificationEmail, err := svc.mailRepository.FindVerificationEmailById(token)

	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	err = svc.userRepository.UpdateEmailVerifiedAt(verificationEmail.Email)
	if err != nil {
		return
	}

	c.Redirect(302, "http://localhost:8087/api/v1")
}

func (svc *GoogleOauthServiceImpl) generateToken(c *gin.Context, userAccount *table.UserAccount, token dto.AccessTokenDto, payload *dto.JWTClaimsPayloadGoogle) string {
	var marshaledData string
	marshaledData, err := pkg.MarshalToString(userAccount)
	if err != nil {
		log.Error("Error when try to marshalling userAccount data")
		pkg.PanicException(constant.UnknownError)
	}

	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	accessTokenExpiry, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
	refreshTokenSecret := os.Getenv("REFRESH_TOKEN_SECRET")
	refreshTokenExpiry, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR"))

	accessToken, err := util.CreateAccessToken(accessTokenSecret, accessTokenExpiry)
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	refreshToken, err := util.CreateRefreshToken(refreshTokenSecret, refreshTokenExpiry)
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	hashedEmail := pkg.HashData(payload.Email)

	batchDataToken := make(map[string]interface{})
	batchDataToken[constant.UserAccountData.GetRedisKey()+":"+hashedEmail] = marshaledData
	batchDataToken[constant.GoogleAccessToken.GetRedisKey()+":"+hashedEmail] = token.AccessToken
	batchDataToken[constant.GoogleRefreshToken.GetRedisKey()+":"+hashedEmail] = token.RefreshToken
	batchDataToken[constant.GoogleIdToken.GetRedisKey()+":"+hashedEmail] = token.IdToken
	batchDataToken[constant.AccessToken.GetRedisKey()+":"+hashedEmail] = accessToken
	batchDataToken[constant.RefreshToken.GetRedisKey()+":"+hashedEmail] = refreshToken

	//err = redisService.PutCache(constant.UserAccountData.GetRedisKey()+" : "+pkg.HashData(payload.Email), marshaledData, c)
	err = svc.redisService.PutCacheBatch(batchDataToken, c)
	if err != nil {
		log.Error("Error when try to put userAccount's data in redis")
		pkg.PanicException(constant.UnknownError)
	}

	return accessToken
}

func ProvideGoogleOauthService(mailService service.MailService, redisService service.RedisService, oauthApiService apiService.OauthApiService,
	mailRepository repository.MailRepository, userRepository repository.UserRepository,
	userProfileRepository repository.UserProfileRepository, tokenRepository repository.AuthTokenRepository) *GoogleOauthServiceImpl {
	googleServiceOnce.Do(func() {
		googleService = &GoogleOauthServiceImpl{
			mailService:           mailService,
			redisService:          redisService,
			oauthApiService:       oauthApiService,
			mailRepository:        mailRepository,
			userRepository:        userRepository,
			userProfileRepository: userProfileRepository,
			authTokenRepository:   tokenRepository,
		}
	})

	return googleService
}
