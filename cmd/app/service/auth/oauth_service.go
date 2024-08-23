package authService

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/repository"
	"sync"
)

var (
	oAuthService     *OAuthServiceImpl
	oAuthServiceOnce sync.Once
)

type OAuthService interface {
	Login(c *gin.Context)
	Callback(c *gin.Context)
}

type OAuthServiceImpl struct {
	userRepository repository.UserRepository
}

func (svc *OAuthServiceImpl) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"response_key": "success",
		"message":      "login api",
	})
}

func (svc *OAuthServiceImpl) Callback(c *gin.Context) {

}

func ProvideAuthService(userRepository repository.UserRepository) *OAuthServiceImpl {
	oAuthServiceOnce.Do(func() {
		oAuthService = &OAuthServiceImpl{
			userRepository: userRepository,
		}
	})

	return oAuthService
}
