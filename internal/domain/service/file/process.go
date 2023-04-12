package file

import (
	"time"

	"github.com/jdcd/account_balance/internal/domain"
)

// ProcessService provides functionality related to process transaction files
type ProcessService struct{}

// MakeSummary return a domain.Summary struct, for a domain.Transaction array
func (s *ProcessService) MakeSummary(transactions []domain.Transaction) domain.Summary {
	var total, totalD, totalC, avrD, avrC float32
	var countC, countD int
	trByMonth := make(map[time.Month]int, 0)

	for _, tr := range transactions {
		trByMonth[tr.Date.Month()]++
		if tr.Movement == domain.Credit {
			total += tr.Value
			totalC += tr.Value
			countC++
		}

		if tr.Movement == domain.Debit {
			total -= tr.Value
			totalD += tr.Value
			countD++
		}
	}

	avrD = 0
	if countD != 0 && totalD != 0 {
		avrD = totalD / float32(countD)
	}

	avrC = 0
	if countC != 0 && totalC != 0 {
		avrC = totalC / float32(countC)
	}

	return domain.Summary{Total: total, TransactionByMonth: trByMonth, AvrDebitAmount: avrD, AvrCreditAmount: avrC}
}
