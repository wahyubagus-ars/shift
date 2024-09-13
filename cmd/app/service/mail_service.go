package service

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/domain/dao/collection"
	"go-shift/cmd/app/domain/dto"
	"go-shift/cmd/app/domain/dto/system"
	"go-shift/cmd/app/repository"
	"go-shift/cmd/app/util"
	"gopkg.in/gomail.v2"
	"html/template"
	"sync"
	"time"
)

var (
	mailServiceOnce sync.Once
	mailService     *MailServiceImpl
)

type MailService interface {
	SendMail(subject string, payload dto.JWTClaimsPayloadGoogle) error
}

type MailServiceImpl struct {
	mailRepository repository.MailRepository
	dialer         *gomail.Dialer
}

func (svc *MailServiceImpl) SendMail(subject string, payload dto.JWTClaimsPayloadGoogle) error {
	htmlFile := "template/verify_email_new_user.html"
	tmpl, err := template.ParseFiles(htmlFile)
	if err != nil {
		log.Fatalf("Failed to parse HTML template: %v", err)
	}

	expiredAt := time.Now()

	verification := &collection.EmailVerification{
		Email:            payload.Email,
		VerificationType: "verify-new-user",
		ExpiredAt:        expiredAt,
	}

	verificationEmail, err := svc.mailRepository.SaveVerificationEmail(verification)

	if err != nil {
		return err
	}

	/** TODO: need to be encrypt */
	token := verificationEmail.ID.Hex()

	data := system.EmailNewUserTemplate{
		Name:             util.CapitalizeFirstLetter(payload.Name),
		VerificationLink: fmt.Sprintf("http://localhost:8087/api/v1/mail/verification?verificationToken=%s", token),
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

	return nil
}

func ProvideMailService(mailRepository repository.MailRepository, dialer *gomail.Dialer) *MailServiceImpl {
	mailServiceOnce.Do(func() {
		mailService = &MailServiceImpl{
			mailRepository: mailRepository,
			dialer:         dialer,
		}
	})

	return mailService
}
