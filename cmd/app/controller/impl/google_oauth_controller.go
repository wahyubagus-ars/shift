package impl

import (
	"github.com/gin-gonic/gin"
	auth "go-shift/cmd/app/service/auth"
	"sync"
)

var (
	googleOauthControllerOnce sync.Once
)

type GoogleOauthControllerImpl struct {
	AuthSvc auth.AuthService
}

func (controller *GoogleOauthControllerImpl) LoginHandler(c *gin.Context) {
	controller.AuthSvc.Login(c)
}

func (controller *GoogleOauthControllerImpl) CallbackHandler(c *gin.Context) {
	controller.AuthSvc.Callback(c)
}

func ProvideGoogleOauthController(as auth.AuthService) *GoogleOauthControllerImpl {
	var controller *GoogleOauthControllerImpl
	googleOauthControllerOnce.Do(func() {
		controller = &GoogleOauthControllerImpl{
			AuthSvc: as,
		}
	})

	return controller
}
