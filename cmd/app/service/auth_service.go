package service

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/repository"
)

type AuthService interface {
	Login(c *gin.Context)
}

type AuthServiceImpl struct {
	UserRepo repository.UserRepository
}

func (as *AuthServiceImpl) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"response_key": "success",
		"message":      "login api",
	})
}
