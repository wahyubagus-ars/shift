package authService

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/constant"
	"go-shift/cmd/app/domain/dao"
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

func (svc *GoogleOauthServiceImpl) Login(c *gin.Context) {
	defer pkg.PanicHandler(c)

	/** TODO: need to move the Authorization-Code to query params*/
	code := c.Request.Header["Authorization-Code"]

	clientId := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	redirectUri := os.Getenv("GOOGLE_REDIRECT_URI")
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

	var userAccount dao.UserAccount
	userAccount, err = svc.userRepository.FindUserByEmail(payload.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userAccount, err = svc.userRepository.SaveInitiateUser(payload.Email, 1)
			if err != nil {
				pkg.PanicException(constant.UnknownError)
			}
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

	authToken := &dao.AuthToken{
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

func (svc *GoogleOauthServiceImpl) Callback(c *gin.Context) {
	authorizationCode := c.Query("code")
	c.Request.Header.Add("Authorization-Code", authorizationCode)
	svc.Login(c)
}

//func GetTokenPayload(token string) (dto.JWTClaimsPayload, error) {
//	parts := strings.Split(token, ".")
//	if len(parts) != 3 {
//		log.Error("Invalid token format")
//		return dto.JWTClaimsPayload{}, &dto.AppError{
//			Message: "Invalid token format",
//		}
//	}
//
//	// Decode the second part, which is the payload
//	payload, err := jwt.DecodeSegment(parts[1])
//	if err != nil {
//		log.Error("Error decoding payload:", err)
//		return dto.JWTClaimsPayload{}, err
//	}
//
//	var claims dto.JWTClaimsPayload
//	err = json.Unmarshal(payload, &claims)
//	if err != nil {
//		log.Error("Error when unmarshalling payload:", err)
//		return dto.JWTClaimsPayload{}, nil
//	}
//
//	return claims, nil
//}

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
