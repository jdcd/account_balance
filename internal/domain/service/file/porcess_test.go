package file

import (
	"testing"
	"time"

	"github.com/jdcd/account_balance/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestMakeSummaryWithAllData(t *testing.T) {
	transactions := []domain.Transaction{
		{Number: 0, Date: auxTestParseDate("7/12"), Movement: domain.Credit, Value: 50.6},
		{Number: 1, Date: auxTestParseDate("7/13"), Movement: domain.Credit, Value: 49.4},
		{Number: 2, Date: auxTestParseDate("8/20"), Movement: domain.Debit, Value: 10.7},
		{Number: 3, Date: auxTestParseDate("8/25"), Movement: domain.Debit, Value: 9.3},
		{Number: 4, Date: auxTestParseDate("9/30"), Movement: domain.Credit, Value: 50},
	}

	expectedSummary := domain.Summary{
		Total:              130,
		TransactionByMonth: map[time.Month]int{7: 2, 8: 2, 9: 1},
		AvrDebitAmount:     10,
		AvrCreditAmount:    50,
	}

	ms := ProcessService{}
	summary := ms.MakeSummary(transactions)

	assert.EqualValues(t, expectedSummary, summary)
}
