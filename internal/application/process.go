package application

import (
	"time"

	"github.com/jdcd/account_balance/internal/domain"
	"github.com/jdcd/account_balance/internal/domain/port/report"
	"github.com/jdcd/account_balance/internal/domain/service/file"
	"github.com/jdcd/account_balance/internal/domain/service/notification"
	"github.com/jdcd/account_balance/pkg"
)

type IProcessTransactionsUseCase interface {
	TransactionHandler(file string, recipes []string)
}

type ProcessTransactionsUseCase struct {
	Sender     notification.ISender
	Finder     file.IFinder
	Reader     file.IReader
	Processor  file.IProcess
	Repository report.Repository
}

func (c *ProcessTransactionsUseCase) TransactionHandler(fileName string, recipes []string) {

	transactions, igTransactions, err := c.Reader.ReadFile(fileName)
	date := time.Now()
	if err != nil {
		c.Finder.Relocate(fileName, date, file.ErrorFolder)
		c.Repository.SaveErrorReport(fileName, date, err)
		return
	}

	summary := c.Processor.MakeSummary(transactions)

	successReport := domain.SuccessReport{
		FileName:           fileName,
		Date:               date,
		Summary:            summary,
		SendTo:             recipes,
		Transactions:       transactions,
		IgnoredTransaction: igTransactions,
	}

	err = c.Sender.SendSummaryNotification(summary, recipes)
	if err != nil {
		c.Finder.Relocate(fileName, date, file.ErrorFolder)
		c.Repository.SaveErrorReport(fileName, date, err)
		return
	}

	err = c.Repository.SaveSuccessReport(successReport)
	if err != nil {
		return
	}

	c.Finder.Relocate(fileName, date, file.SuccessFolder)
	pkg.InfoLogger().Printf("file %s processed successfully", fileName)
}
