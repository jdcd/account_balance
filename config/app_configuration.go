package config

import (
	"fmt"
	"os"

	"github.com/jdcd/account_balance/pkg"
)

// GetConfigurations return *AppConfiguration instance, with configured data
func GetConfigurations() *AppConfiguration {
	return &AppConfiguration{
		EmailSender:     os.Getenv("EMAIL_SENDER"),
		EmailPwd:        os.Getenv("EMAIL_PWD"),
		EmailSenderName: os.Getenv("EMAIL_SENDER_NAME"),
		SmtpServer:      os.Getenv("SMTP_SERVER"),
		SmtpPort:        os.Getenv("SMTP_PORT"),
		DbUrl:           os.Getenv("DB_URL"),
	}
}

// AppConfiguration contains data from env variables, to works inside the app
type AppConfiguration struct {
	EmailSender     string
	EmailPwd        string
	EmailSenderName string
	SmtpServer      string
	SmtpPort        string
	DbUrl           string
}

func (r *AppConfiguration) CheckData() error {

	if !pkg.IsValidateEmail(r.EmailSender) {
		return fmt.Errorf("invalid EMAIL_SENDER")
	}

	if !pkg.IsValidSMTPServer(r.SmtpServer) {
		return fmt.Errorf("invalid SMTP_SERVER")
	}

	if !pkg.IsValidStringOfNumbers(r.SmtpPort) {
		return fmt.Errorf("invalid SMTP_PORT")
	}

	return nil
}
