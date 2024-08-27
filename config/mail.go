package config

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func InitMail() *gomail.Dialer {

	dialer := gomail.NewDialer(
		"sandbox.smtp.mailtrap.io",
		2525,
		"4c6cd32a2f97ee",
		"ca18752ec66837",
	)

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return dialer

}
