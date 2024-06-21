package service

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/repository"
	"sync"
)

var (
	googleService     *GoogleOauthServiceImpl
	googleServiceOnce sync.Once
)

type GoogleOauthServiceImpl struct {
	userRepository repository.UserRepository
}

func (svc *GoogleOauthServiceImpl) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"response_key": "success",
		"message":      "login api google",
	})
}

func ProvideGoogleOauthService(ur repository.UserRepository) *GoogleOauthServiceImpl {
	googleServiceOnce.Do(func() {
		googleService = &GoogleOauthServiceImpl{
			userRepository: ur,
		}
	})

	return googleService
}
