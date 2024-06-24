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
	CallbackHandler(c *gin.Context)
}

type AuthControllerImpl struct {
	AuthSvc service.AuthService
}

func (controller *AuthControllerImpl) LoginHandler(c *gin.Context) {
	controller.AuthSvc.Login(c)
}

func (controller *AuthControllerImpl) CallbackHandler(c *gin.Context) {

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
