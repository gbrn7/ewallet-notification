package external

import (
	"ewallet-notification/helpers"
	"strconv"

	"gopkg.in/gomail.v2"
)

type Email struct {
	To      string
	Subject string
	Body    string
}

func (e *Email) SendEmail() error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("To", e.To)
	mailer.SetHeader("From", helpers.GetEnv("SMTP_AUTH_EMAIL", ""))
	mailer.SetHeader("Subject", e.Subject)
	mailer.SetBody("text/html", e.Body)

	smtpPort := helpers.GetEnv("SMTP_PORT", "")
	intSmtPort, _ := strconv.Atoi(smtpPort)

	var (
		configHost         = helpers.GetEnv("SMTP_HOST", "")
		configAuthEmail    = helpers.GetEnv("SMTP_AUTH_EMAIL", "")
		configAuthPassword = helpers.GetEnv("SMTP_AUTH_PASSWORD", "")
	)

	dialer := gomail.NewDialer(
		configHost,
		intSmtPort,
		configAuthEmail,
		configAuthPassword,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
