package controller

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/service"
	"sync"
)

var (
	mailControllerOnce sync.Once
	mailController     *MailControllerImpl
)

type MailController interface {
	VerifyEmail(c *gin.Context)
}

type MailControllerImpl struct {
	mailService service.MailService
}

func (controller *MailControllerImpl) VerifyEmail(c *gin.Context) {
	//controller.mailService.VerifyEmail(c)
}

func ProvideMailController(mailService service.MailService) *MailControllerImpl {
	mailControllerOnce.Do(func() {
		mailController = &MailControllerImpl{
			mailService: mailService,
		}
	})

	return mailController
}
