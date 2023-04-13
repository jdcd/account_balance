package notification

import (
	"fmt"
	"net/smtp"
	"path/filepath"

	"github.com/jdcd/account_balance/internal/domain"
	"github.com/jdcd/account_balance/pkg"
	"github.com/jordan-wright/email"
)

const (
	imageNotFoundWarning = "branding image not found: %s"
	emailDidNotSend      = "error sending notification: %s"
	htmlResourcePath     = "../../../../resources/html/"
)

type SmtpEmailSender struct {
	EmailSender     string
	EmailPwd        string
	EmailSenderName string
	SmtpServer      string
	SmtpPort        string
	SmtpIdentity    string
}

func (s *SmtpEmailSender) SendNotification(n domain.Notification) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", s.EmailSenderName, s.EmailSender)
	e.To = n.Recipients
	e.Subject = n.Subject
	logoPath := filepath.Join(htmlResourcePath, n.Branding)

	_, err := e.AttachFile(logoPath)
	if err != nil {
		pkg.WarningLogger().Printf(imageNotFoundWarning, err)
	}

	e.HTML = []byte(n.Content)
	smtpHost := fmt.Sprintf("%s:%s", s.SmtpServer, s.SmtpPort)

	err = e.Send(s.SmtpServer, smtp.PlainAuth(s.SmtpIdentity, s.EmailSender, s.EmailPwd, smtpHost))
	if err != nil {
		return fmt.Errorf(emailDidNotSend, err)
	}

	return nil
}
