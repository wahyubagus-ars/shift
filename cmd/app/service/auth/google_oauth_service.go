package authService

import (
	"errors"
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
	"os"
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

	data := system.ApiResponse[any]{
		ResponseKey:     constant.Success.GetResponseStatus(),
		ResponseMessage: constant.Success.GetResponseMessage(),
		Data:            nil,
	}

	c.Header("Authorization-Token", token.IdToken)

	c.JSON(200, data)
}

func (svc *GoogleOauthServiceImpl) SignInCallback(c *gin.Context) {
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

	c.Header("Authorization-Token", token.IdToken)

	c.JSON(200, data)
}

func (svc *GoogleOauthServiceImpl) SignUpCallback(c *gin.Context) {
	svc.SignUp(c)
}

func (svc *GoogleOauthServiceImpl) oauthProcess(c *gin.Context, processType string, redirectUri string) (dto.AccessTokenDto, dto.JWTClaimsPayload) {
	defer pkg.PanicHandler(c)

	/** TODO: need to move the Authorization-Code to query params*/
	code := c.Query("code")

	clientId := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	getTokenUrl := os.Getenv("GOOGLE_ACCESS_TOKEN_URL")

	token, err := svc.oauthApiService.GetAccessToken(code, clientId, clientSecret, getTokenUrl, redirectUri)
	if err != nil {
		log.Error("Error when get access token :: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	payload, err := util.GetTokenPayload(token.IdToken)

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

	var marshaledData string
	marshaledData, err = pkg.MarshalToString(userAccount)
	if err != nil {
		log.Error("Error when try to marshalling userAccount data")
		pkg.PanicException(constant.UnknownError)
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

	hashedEmail := pkg.HashData(payload.Email)

	batchDataToken := make(map[string]interface{})
	batchDataToken[constant.UserAccountData.GetRedisKey()+":"+hashedEmail] = marshaledData
	batchDataToken[constant.AccessToken.GetRedisKey()+":"+hashedEmail] = token.AccessToken
	batchDataToken[constant.RefreshToken.GetRedisKey()+":"+hashedEmail] = token.RefreshToken
	batchDataToken[constant.IdToken.GetRedisKey()+":"+hashedEmail] = token.IdToken

	//err = redisService.PutCache(constant.UserAccountData.GetRedisKey()+" : "+pkg.HashData(payload.Email), marshaledData, c)
	err = svc.redisService.PutCacheBatch(batchDataToken, c)
	if err != nil {
		log.Error("Error when try to put userAccount's data in redis")
		pkg.PanicException(constant.UnknownError)
	}

	return token, payload
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
