package notification

import (
	"github.com/jdcd/account_balance/internal/domain"
	"github.com/jdcd/account_balance/internal/domain/port/notification"
	"github.com/jdcd/account_balance/internal/domain/port/template"
)

// ISender provides contracts related to send summary notifications
type ISender interface {
	SendSummaryNotification(summary domain.Summary, recipients []string) error
}

// SenderService provides implements ISender with bridge between port.notification.Delivery and template.Generator
type SenderService struct {
	NotificationDeliveryService notification.Delivery
	SummaryTemplateGenerator    template.Generator
}

const (
	notificationSubject  = "Account summary notification"
	notificationBranding = "logo.png"
)

func (s *SenderService) SendSummaryNotification(summary domain.Summary, recipients []string) error {
	content, err := s.SummaryTemplateGenerator.FormatSummary(summary)
	if err != nil {
		return err
	}

	n := domain.Notification{
		Content:    content,
		Subject:    notificationSubject,
		Recipients: recipients,
		Branding:   notificationBranding,
	}

	return s.NotificationDeliveryService.SendNotification(n)

}
