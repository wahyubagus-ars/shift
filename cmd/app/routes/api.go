package routes

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/provider"
)

func Init(init *provider.Initialization) *gin.Engine {
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
			auth.GET("/login", init.AuthCtrl.Login)
		}
	}

	return r
}
