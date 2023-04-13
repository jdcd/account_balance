package mock

import (
	"github.com/jdcd/account_balance/internal/domain"
	"github.com/stretchr/testify/mock"
)

// DeliveryMock is a mock implementation of Delivery interface
type DeliveryMock struct {
	mock.Mock
}

func (m *DeliveryMock) SendNotification(notification domain.Notification) error {
	args := m.Called(notification)
	return args.Error(0)
}
