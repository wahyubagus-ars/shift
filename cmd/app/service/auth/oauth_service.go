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
	SignIn(c *gin.Context)
	SignInCallback(c *gin.Context)
	SignUp(c *gin.Context)
	SignUpCallback(c *gin.Context)
}

type OAuthServiceImpl struct {
	userRepository repository.UserRepository
}

func (svc *OAuthServiceImpl) SignIn(c *gin.Context) {
	c.JSON(200, gin.H{
		"response_key": "success",
		"message":      "login api",
	})
}

func (svc *OAuthServiceImpl) SignInCallback(c *gin.Context) {

}

func (svc *OAuthServiceImpl) SignUp(c *gin.Context) {
	c.JSON(200, gin.H{
		"response_key": "success",
		"message":      "login api",
	})
}

func (svc *OAuthServiceImpl) SignUpCallback(c *gin.Context) {

}

func ProvideAuthService(userRepository repository.UserRepository) *OAuthServiceImpl {
	oAuthServiceOnce.Do(func() {
		oAuthService = &OAuthServiceImpl{
			userRepository: userRepository,
		}
	})

	return oAuthService
}
