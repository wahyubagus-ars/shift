package authService

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/constant"
	"go-shift/cmd/app/domain/dao/table"
	"go-shift/cmd/app/domain/dto"
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
	redisService          service.RedisService
	oauthApiService       apiService.OauthApiService
	userRepository        repository.UserRepository
	userProfileRepository repository.UserProfileRepository
	authTokenRepository   repository.AuthTokenRepository
}

func (svc *GoogleOauthServiceImpl) SignIn(c *gin.Context) {
	redirectUri := os.Getenv("GOOGLE_SIGN_IN_REDIRECT_URI")
	svc.oauthProcess(c, "sign-in", redirectUri)
}

func (svc *GoogleOauthServiceImpl) SignInCallback(c *gin.Context) {
	authorizationCode := c.Query("code")
	c.Request.Header.Add("Authorization-Code", authorizationCode)
	svc.SignIn(c)
}

func (svc *GoogleOauthServiceImpl) SignUp(c *gin.Context) {
	redirectUri := os.Getenv("GOOGLE_SIGN_UP_REDIRECT_URI")
	svc.oauthProcess(c, "sign-up", redirectUri)
}

func (svc *GoogleOauthServiceImpl) SignUpCallback(c *gin.Context) {
	authorizationCode := c.Query("code")
	c.Request.Header.Add("Authorization-Code", authorizationCode)
	svc.SignUp(c)
}

func (svc *GoogleOauthServiceImpl) oauthProcess(c *gin.Context, processType string, redirectUri string) {
	defer pkg.PanicHandler(c)

	/** TODO: need to move the Authorization-Code to query params*/
	code := c.Request.Header["Authorization-Code"]

	clientId := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	getTokenUrl := os.Getenv("GOOGLE_ACCESS_TOKEN_URL")

	token, err := svc.oauthApiService.GetAccessToken(code[0], clientId, clientSecret, getTokenUrl, redirectUri)
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

	data := dto.ApiResponse[any]{
		ResponseKey:     constant.Success.GetResponseStatus(),
		ResponseMessage: constant.Success.GetResponseMessage(),
		Data:            nil,
	}

	c.Header("Authorization-Token", token.IdToken)

	c.JSON(200, data)
}

func ProvideGoogleOauthService(redisService service.RedisService, oauthApiService apiService.OauthApiService,
	userRepository repository.UserRepository,
	userProfileRepository repository.UserProfileRepository,
	tokenRepository repository.AuthTokenRepository) *GoogleOauthServiceImpl {
	googleServiceOnce.Do(func() {
		googleService = &GoogleOauthServiceImpl{
			redisService:          redisService,
			oauthApiService:       oauthApiService,
			userRepository:        userRepository,
			userProfileRepository: userProfileRepository,
			authTokenRepository:   tokenRepository,
		}
	})

	return googleService
}
