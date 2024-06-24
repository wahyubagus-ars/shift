package service

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

func (as *AuthServiceImpl) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"response_key": "success",
		"message":      "login api",
	})
}

func (as *AuthServiceImpl) Callback(c *gin.Context) {

}

func ProvideAuthService(ur repository.UserRepository) *AuthServiceImpl {
	authServiceOnce.Do(func() {
		authService = &AuthServiceImpl{
			userRepository: ur,
		}
	})

	return authService
}
