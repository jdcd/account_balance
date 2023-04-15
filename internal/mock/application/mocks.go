package mock

import (
	"github.com/stretchr/testify/mock"
)

// FilePickerUseCaseMock implements application.IFilePickerUseCase
type FilePickerUseCaseMock struct {
	mock.Mock
}

func (m *FilePickerUseCaseMock) SelectFileHandler() (filePath string, err error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

// ProcessTransactionsUseCaseMock implements application.IProcessTransactionsUseCase
type ProcessTransactionsUseCaseMock struct {
	mock.Mock
}

func (m *ProcessTransactionsUseCaseMock) TransactionHandler(file string, recipes []string) {
	m.Called(file, recipes)
}
