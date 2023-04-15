package application

import (
	"github.com/jdcd/account_balance/internal/domain/service/file"
	"github.com/jdcd/account_balance/internal/domain/service/notification"
	"github.com/jdcd/account_balance/pkg"
)

type IProcessTransactionsUseCase interface {
	TransactionHandler(file string, recipes []string)
}

type ProcessTransactionsUseCase struct {
	Sender    notification.ISender
	Finder    file.IFinder
	Reader    file.IReader
	Processor file.IProcess
}

func (c *ProcessTransactionsUseCase) TransactionHandler(fileName string, recipes []string) {

	transactions, _, err := c.Reader.ReadFile(fileName)
	if err != nil {
		c.Finder.Relocate(fileName, file.ErrorFolder)
		return
	}

	summary := c.Processor.MakeSummary(transactions)

	err = c.Sender.SendSummaryNotification(summary, recipes)
	if err != nil {
		return
	}

	c.Finder.Relocate(fileName, file.SuccessFolder)
	pkg.InfoLogger().Printf("file %s processed successfully", fileName)
}
