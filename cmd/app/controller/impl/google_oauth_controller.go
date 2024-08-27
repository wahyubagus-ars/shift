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
	AuthSvc auth.OAuthService
}

func (controller *GoogleOauthControllerImpl) SignInHandler(c *gin.Context) {
	controller.AuthSvc.SignIn(c)
}

func (controller *GoogleOauthControllerImpl) SignInCallbackHandler(c *gin.Context) {
	controller.AuthSvc.SignInCallback(c)
}

func (controller *GoogleOauthControllerImpl) SignUpHandler(c *gin.Context) {
	controller.AuthSvc.SignIn(c)
}

func (controller *GoogleOauthControllerImpl) SignUpCallbackHandler(c *gin.Context) {
	controller.AuthSvc.SignUpCallback(c)
}

func ProvideGoogleOauthController(as auth.OAuthService) *GoogleOauthControllerImpl {
	var controller *GoogleOauthControllerImpl
	googleOauthControllerOnce.Do(func() {
		controller = &GoogleOauthControllerImpl{
			AuthSvc: as,
		}
	})

	return controller
}
