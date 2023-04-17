package service

import (
	"time"

	"github.com/jdcd/account_balance/internal/domain"
	"github.com/jdcd/account_balance/internal/domain/service/file"
	"github.com/stretchr/testify/mock"
)

// ProcessMock implements file.IProcess
type ProcessMock struct {
	mock.Mock
}

func (m *ProcessMock) MakeSummary(transactions []domain.Transaction) domain.Summary {
	args := m.Called(transactions)
	return args.Get(0).(domain.Summary)
}

// ReaderMock implements file.IReader
type ReaderMock struct {
	mock.Mock
}

func (m *ReaderMock) ReadFile(fileName string) ([]domain.Transaction, []domain.IgnoredTransaction, error) {
	args := m.Called(fileName)
	return args.Get(0).([]domain.Transaction), args.Get(1).([]domain.IgnoredTransaction), args.Error(2)
}

// SenderMock implements notification.ISender
type SenderMock struct {
	mock.Mock
}

func (m *SenderMock) SendSummaryNotification(summary domain.Summary, recipients []string) error {
	args := m.Called(summary, recipients)
	return args.Error(0)
}

// FinderMock implements file.IFinder
type FinderMock struct {
	mock.Mock
}

func (m *FinderMock) GiveNextFileName() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *FinderMock) Relocate(filePath string, date time.Time, newFolder file.Directory) {
	m.Called(filePath, date, newFolder)
}
