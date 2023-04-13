package notification

import "github.com/jdcd/account_balance/internal/domain"

// Delivery contains the contracts for any notification sender.
type Delivery interface {
	SendNotification(notification domain.Notification) error
}
