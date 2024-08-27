package service

import (
	"bytes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/domain/dto"
	"go-shift/cmd/app/domain/dto/system"
	"gopkg.in/gomail.v2"
	"html/template"
	"sync"
)

var (
	mailServiceOnce sync.Once
	mailService     *MailServiceImpl
)

type MailService interface {
	SendMail(subject string, payload dto.JWTClaimsPayload) error
	VerifyEmail(c *gin.Context)
}

type MailServiceImpl struct {
	dialer *gomail.Dialer
}

func (svc *MailServiceImpl) SendMail(subject string, payload dto.JWTClaimsPayload) error {
	htmlFile := "template/verify_email_new_user.html"
	tmpl, err := template.ParseFiles(htmlFile)
	if err != nil {
		log.Fatalf("Failed to parse HTML template: %v", err)
	}

	data := system.EmailNewUserTemplate{
		Name:             payload.Name,
		VerificationLink: "http://localhost:8087/api/v1/mail/verification",
	}

	var renderedHTML bytes.Buffer
	if err := tmpl.Execute(&renderedHTML, data); err != nil {
		log.Fatalf("Failed to render HTML template: %v", err)
	}

	mail := gomail.NewMessage()

	mailFrom := "mailtrap-admin@mail.com"
	mailReplyTo := "mailtrap-noreply@mail.com"
	mail.SetHeader("From", mailFrom)
	mail.SetHeader("To", payload.Email)
	mail.SetHeader("Reply-To", mailReplyTo)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", renderedHTML.String())

	if err := svc.dialer.DialAndSend(mail); err != nil {
		return err
	}

	// TODO: need save to mongodb invitation data

	return nil
}

func (svc *MailServiceImpl) VerifyEmail(c *gin.Context) {
	c.JSON(200, "verification email")
}

func ProvideMailService(dialer *gomail.Dialer) *MailServiceImpl {
	mailServiceOnce.Do(func() {
		mailService = &MailServiceImpl{
			dialer: dialer,
		}
	})

	return mailService
}
