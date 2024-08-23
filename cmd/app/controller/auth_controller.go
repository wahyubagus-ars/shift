package controller

import (
	"github.com/gin-gonic/gin"
	auth "go-shift/cmd/app/service/auth"
	"sync"
)

var (
	authControllerOnce sync.Once
)

type AuthController interface {
	LoginHandler(c *gin.Context)
	CallbackHandler(c *gin.Context)
}

type AuthControllerImpl struct {
	AuthSvc auth.OAuthService
}

func (controller *AuthControllerImpl) LoginHandler(c *gin.Context) {
	controller.AuthSvc.Login(c)
}

func (controller *AuthControllerImpl) CallbackHandler(c *gin.Context) {

}

func ProvideAuthController(as auth.OAuthService) *AuthControllerImpl {
	var controller *AuthControllerImpl
	authControllerOnce.Do(func() {
		controller = &AuthControllerImpl{
			AuthSvc: as,
		}
	})

	return controller
}
