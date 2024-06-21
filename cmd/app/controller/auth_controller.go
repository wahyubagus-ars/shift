package controller

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/service"
	"sync"
)

var (
	authControllerOnce sync.Once
)

type AuthController interface {
	LoginHandler(c *gin.Context)
}

type AuthControllerImpl struct {
	AuthSvc service.AuthService
}

func (ac *AuthControllerImpl) LoginHandler(c *gin.Context) {
	ac.AuthSvc.Login(c)
}

func ProvideAuthController(as service.AuthService) *AuthControllerImpl {
	var controller *AuthControllerImpl
	authControllerOnce.Do(func() {
		controller = &AuthControllerImpl{
			AuthSvc: as,
		}
	})

	return controller
}
