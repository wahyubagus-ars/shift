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
	SignInHandler(c *gin.Context)
	SignInCallbackHandler(c *gin.Context)

	SignUpHandler(c *gin.Context)
	SignUpCallbackHandler(c *gin.Context)
}

type AuthControllerImpl struct {
	AuthSvc auth.OAuthService
}

func (controller *AuthControllerImpl) SignInHandler(c *gin.Context) {
	controller.AuthSvc.SignIn(c)
}

func (controller *AuthControllerImpl) SignInCallbackHandler(c *gin.Context) {

}

func (controller *AuthControllerImpl) SignUpHandler(c *gin.Context) {
	controller.AuthSvc.SignIn(c)
}

func (controller *AuthControllerImpl) SignUpCallbackHandler(c *gin.Context) {

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
