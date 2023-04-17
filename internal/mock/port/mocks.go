package port

import (
	"time"

	"github.com/jdcd/account_balance/internal/domain"
	"github.com/stretchr/testify/mock"
)

// DeliveryMock implements notification.Delivery
type DeliveryMock struct {
	mock.Mock
}

func (m *DeliveryMock) SendNotification(notification domain.Notification) error {
	args := m.Called(notification)
	return args.Error(0)
}

// GeneratorMock implements template.Generator
type GeneratorMock struct {
	mock.Mock
}

func (m *GeneratorMock) FormatSummary(summary domain.Summary) (string, error) {
	args := m.Called(summary)
	return args.String(0), args.Error(1)
}

// RepositoryMock implements report.Repository
type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) SaveSuccessReport(report domain.SuccessReport) error {
	args := m.Called(report)
	return args.Error(0)
}

func (m *RepositoryMock) SaveErrorReport(fileName string, date time.Time, err error) {
	m.Called(fileName, date, err)
}
