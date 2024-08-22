package authService

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/repository"
	"sync"
)

var (
	authService     *AuthServiceImpl
	authServiceOnce sync.Once
)

type AuthService interface {
	Login(c *gin.Context)
	Callback(c *gin.Context)
}

type AuthServiceImpl struct {
	userRepository repository.UserRepository
}

func (svc *AuthServiceImpl) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"response_key": "success",
		"message":      "login api",
	})
}

func (svc *AuthServiceImpl) Callback(c *gin.Context) {

}

func ProvideAuthService(userRepository repository.UserRepository) *AuthServiceImpl {
	authServiceOnce.Do(func() {
		authService = &AuthServiceImpl{
			userRepository: userRepository,
		}
	})

	return authService
}
