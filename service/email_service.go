package service

import (
	"authentication-service/config"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

func InitEmailService(config config.Config) smtp.Auth {

	auth := smtp.PlainAuth("", config.SMTP_SENDER, config.SMTP_AUTH_PASSWORD, config.SMTP_SERVER)

	log.Printf("SMTP server: %s, port: %s, sender: %s", config.SMTP_SERVER, config.SMTP_PORT, config.SMTP_SENDER)

	return auth

}

func SendEmail(to []string, subject string, body string, message string) ([]byte, error) {
	config := config.Load()

	auth := InitEmailService(config)

	msg := []byte("To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n" +
		"\r\n" +
		message)
	addr := config.SMTP_SERVER + ":" + config.SMTP_PORT

	if err := smtp.SendMail(addr, auth, config.SMTP_SENDER, to, msg); err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("Email sent to: %s", to)), nil
}
