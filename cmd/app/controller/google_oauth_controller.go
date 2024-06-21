package controller

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/service"
	"sync"
)

var (
	googleOauthControllerOnce sync.Once
)

type GoogleOauthControllerImpl struct {
	AuthSvc service.AuthService
}

func (ac *GoogleOauthControllerImpl) LoginHandler(c *gin.Context) {
	ac.AuthSvc.Login(c)
}

func ProvideGoogleOauthController(as service.AuthService) *GoogleOauthControllerImpl {
	var controller *GoogleOauthControllerImpl
	googleOauthControllerOnce.Do(func() {
		controller = &GoogleOauthControllerImpl{
			AuthSvc: as,
		}
	})

	return controller
}
