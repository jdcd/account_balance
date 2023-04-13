package template

import (
	"github.com/jdcd/account_balance/internal/domain"
)

// Generator contains the contracts to create any summary template generator, to send it
type Generator interface {
	FormatSummary(summary domain.Summary) (string, error)
}
