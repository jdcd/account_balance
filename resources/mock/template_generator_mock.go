package mock

import (
	"github.com/jdcd/account_balance/internal/domain"
	"github.com/stretchr/testify/mock"
)

type GeneratorMock struct {
	mock.Mock
}

func (m *GeneratorMock) FormatSummary(summary domain.Summary) (string, error) {
	args := m.Called(summary)
	return args.String(0), args.Error(1)
}
