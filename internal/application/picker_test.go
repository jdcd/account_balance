package application

import (
	"testing"

	"github.com/jdcd/account_balance/internal/mock/service"
	"github.com/stretchr/testify/assert"
)

func TestPickerUseCase(t *testing.T) {
	FinderMock := &service.FinderMock{}
	ms := FilePickerUseCase{Finder: FinderMock}
	expFileName := "file.csv"

	FinderMock.On("GiveNextFileName").Return(expFileName, nil)

	fileName, err := ms.SelectFileHandler()

	assert.Nil(t, err)
	assert.Equal(t, expFileName, fileName)
	FinderMock.AssertExpectations(t)
}
