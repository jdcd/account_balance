package report

import (
	"time"

	"github.com/jdcd/account_balance/internal/domain"
)

// Repository describes the behavior, for save reports about account transactions in a persistent repository
type Repository interface {
	SaveSuccessReport(report domain.SuccessReport) error
	SaveErrorReport(fileName string, date time.Time, err error)
}
