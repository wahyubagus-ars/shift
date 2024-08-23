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

		auth := api.Group("/auth")
		{
			auth.GET("/login", init.AuthController.LoginHandler)
		}

		oauth := api.Group("/oauth")
		{
			googleOauth := oauth.Group("/google")
			{
				googleOauth.GET("/login-with-google", init.GoogleOauthController.LoginHandler)
				googleOauth.GET("/callback", init.GoogleOauthController.CallbackHandler)
			}
		}

	}

	return r
}
