package application

import (
	"errors"
	"testing"
	"time"

	"github.com/jdcd/account_balance/internal/domain"
	"github.com/jdcd/account_balance/internal/domain/service/file"
	"github.com/jdcd/account_balance/internal/mock/port"
	"github.com/jdcd/account_balance/internal/mock/service"
	"github.com/stretchr/testify/mock"
)

func TestTransactionHandlerAllFine(t *testing.T) {
	fileName := "file.csv"
	recipes := []string{"cheems@gmail.com", "doge@gmial.com"}
	senderMock := &service.SenderMock{}
	readerMock := &service.ReaderMock{}
	processMock := &service.ProcessMock{}
	finderMock := &service.FinderMock{}
	repositoryMock := &port.RepositoryMock{}
	ms := ProcessTransactionsUseCase{Sender: senderMock, Finder: finderMock, Reader: readerMock, Processor: processMock,
		Repository: repositoryMock}
	trs := []domain.Transaction{{Number: 1, Date: time.Time{}, Movement: domain.Credit, Value: 32}}
	igTrs := []domain.IgnoredTransaction{{ID: "2", Date: "8/56", Transaction: "+10.6", Reason: "date is wrong"}}
	summary := domain.Summary{Total: 12, TransactionByMonth: nil, AvrDebitAmount: 21, AvrCreditAmount: 12}
	readerMock.On("ReadFile", fileName).Return(trs, igTrs, nil).Once()
	processMock.On("MakeSummary", trs).Return(summary).Once()
	senderMock.On("SendSummaryNotification", summary, recipes).Return(nil).Once()
	finderMock.On("Relocate", fileName, mock.Anything, file.SuccessFolder).Once()
	repositoryMock.On("SaveSuccessReport", mock.Anything).Return(nil).Once()

	ms.TransactionHandler(fileName, recipes)

	readerMock.AssertExpectations(t)
	processMock.AssertExpectations(t)
	senderMock.AssertExpectations(t)
	finderMock.AssertExpectations(t)
	repositoryMock.AssertExpectations(t)
}

func TestTransactionHandlerReadFileFails(t *testing.T) {
	fileName := "file.csv"
	recipes := []string{"cheems@gmail.com", "doge@gmial.com"}
	readerMock := &service.ReaderMock{}
	finderMock := &service.FinderMock{}
	repositoryMock := &port.RepositoryMock{}
	ms := ProcessTransactionsUseCase{Reader: readerMock, Finder: finderMock, Repository: repositoryMock}
	var trs []domain.Transaction
	var igTrs []domain.IgnoredTransaction
	errorExpected := errors.New("error reading")
	readerMock.On("ReadFile", fileName).Return(trs, igTrs, errorExpected).Once()
	finderMock.On("Relocate", fileName, mock.Anything, file.ErrorFolder).Once()
	repositoryMock.On("SaveErrorReport", fileName, mock.Anything, errorExpected).Once()

	ms.TransactionHandler(fileName, recipes)

	readerMock.AssertExpectations(t)
	finderMock.AssertExpectations(t)
	repositoryMock.AssertExpectations(t)
}

func TestTransactionHandlerSendNotificationFails(t *testing.T) {
	fileName := "file.csv"
	recipes := []string{"cheems@gmail.com", "doge@gmial.com"}
	readerMock := &service.ReaderMock{}
	senderMock := &service.SenderMock{}
	processMock := &service.ProcessMock{}
	repositoryMock := &port.RepositoryMock{}
	finderMock := &service.FinderMock{}
	ms := ProcessTransactionsUseCase{Reader: readerMock, Sender: senderMock, Processor: processMock,
		Repository: repositoryMock, Finder: finderMock}
	trs := []domain.Transaction{{Number: 1, Date: time.Time{}, Movement: domain.Credit, Value: 32}}
	igTrs := []domain.IgnoredTransaction{{ID: "2", Date: "8/56", Transaction: "+10.6", Reason: "date is wrong"}}
	sendingError := errors.New("error reading")
	summary := domain.Summary{Total: 12, TransactionByMonth: nil, AvrDebitAmount: 21, AvrCreditAmount: 12}
	readerMock.On("ReadFile", fileName).Return(trs, igTrs, nil).Once()
	senderMock.On("SendSummaryNotification", summary, recipes).Return(sendingError).Once()
	processMock.On("MakeSummary", trs).Return(summary).Once()
	repositoryMock.On("SaveErrorReport", fileName, mock.Anything, sendingError).Once()
	finderMock.On("Relocate", fileName, mock.Anything, file.ErrorFolder).Once()

	ms.TransactionHandler(fileName, recipes)

	readerMock.AssertExpectations(t)
	senderMock.AssertExpectations(t)
	repositoryMock.AssertExpectations(t)
	processMock.AssertExpectations(t)
}
