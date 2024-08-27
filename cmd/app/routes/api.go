package routes

import (
	"github.com/gin-gonic/gin"
	"go-shift/config"
)

func Init(init *config.Initialization) *gin.Engine {
	r := gin.New()

	api := r.Group("/api/v1")
	{
		api.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"response_key": "success",
				"message":      "this is shitf-service auth",
			})
		})

		/** auth api endpoint */
		auth := api.Group("/auth")
		{
			auth.GET("/login", init.AuthController.SignInHandler)
		}

		oauth := api.Group("/oauth")
		{
			googleOauth := oauth.Group("/google")
			{
				googleOauth.GET("/sign-in", init.GoogleOauthController.SignInHandler)
				googleOauth.GET("/sign-in-callback", init.GoogleOauthController.SignInCallbackHandler)
				googleOauth.GET("/sign-up", init.GoogleOauthController.SignUpHandler)
				googleOauth.GET("/sign-up-callback", init.GoogleOauthController.SignUpCallbackHandler)
			}
		}

		/** mail api endpoint */
		mail := api.Group("/mail")
		{
			mail.GET("/verification", init.GoogleOauthController.VerifyEmail)
		}

		user := api.Group("/user")
		{
			user.GET("/profile", init.UserProfileController.GetUserProfile)
		}

		tracking := api.Group("/tracking")
		{
			tracking.GET("/time-entries", init.TimeTrackingController.GetTimeEntries)
			tracking.POST("/time-entries", init.TimeTrackingController.SubmitTimeEntry)
		}

	}

	return r
}
