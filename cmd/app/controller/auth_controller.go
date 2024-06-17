package controller

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/service"
)

type AuthController interface {
	Login(c *gin.Context)
}

type AuthControllerImpl struct {
	AuthSvc service.AuthService
}

func (ac *AuthControllerImpl) Login(c *gin.Context) {
	ac.AuthSvc.Login(c)
}
