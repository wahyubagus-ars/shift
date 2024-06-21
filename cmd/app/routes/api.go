package routes

import (
	"github.com/gin-gonic/gin"
	"go-shift/config"
)

func Init(init *config.Initialization) *gin.Engine {
	r := gin.New()

	api := r.Group("/api")
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
			auth.GET("/login-with-google", init.GoogleOauthController.LoginHandler)
		}
	}

	return r
}
