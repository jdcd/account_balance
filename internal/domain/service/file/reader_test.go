package file

import (
	"fmt"
	"testing"
	"time"

	"github.com/jdcd/account_balance/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestReadFileWithErrorsAndValidData(t *testing.T) {
	path := "test.csv"
	ignoredExpected := []domain.IgnoredTransaction{
		{ID: "5", Date: "8/16", Transaction: "10.5", Reason: fmt.Sprintf(invalidTransactionFormatError, "10.5")},
		{ID: "FAIL", Date: "8/17", Transaction: "+9.5", Reason: "strconv.Atoi: parsing \"FAIL\": invalid syntax"},
		{ID: "7", Date: "14/20", Transaction: "+16.5", Reason: "parsing time \"14/20/2023\": month out of range"},
		{ID: "8", Date: "8/19", Transaction: "+FAIL", Reason: fmt.Sprintf(invalidTransactionFormatError, "+FAIL")},
	}
	transactionsExpected := []domain.Transaction{
		{Number: 0, Date: auxTestParseDate("7/11"), Movement: domain.Credit, Value: 10.5},
		{Number: 1, Date: auxTestParseDate("7/12"), Movement: domain.Credit, Value: 15.5},
		{Number: 2, Date: auxTestParseDate("8/13"), Movement: domain.Debit, Value: 8.5},
		{Number: 3, Date: auxTestParseDate("8/14"), Movement: domain.Debit, Value: 10.5},
		{Number: 4, Date: auxTestParseDate("8/15"), Movement: domain.Credit, Value: 5.5},
	}

	ms := Service{}
	tr, ig, err := ms.ReadFile(path)

	assert.Nil(t, err)
	assert.EqualValues(t, transactionsExpected, tr)
	assert.EqualValues(t, ignoredExpected, ig)
}

func TestWhenNotFoundFileThenReadFileShouldReturnError(t *testing.T) {
	path := "notExists.csv"
	var ignoredExpected []domain.IgnoredTransaction
	var transactionsExpected []domain.Transaction
	errorExpected := fmt.Errorf("open %s: no such file or directory", path)

	ms := Service{}
	tr, ig, err := ms.ReadFile(path)

	assert.Equal(t, errorExpected.Error(), err.Error())
	assert.EqualValues(t, transactionsExpected, tr)
	assert.EqualValues(t, ignoredExpected, ig)
}

func auxTestParseDate(date string) time.Time {
	fDate, _ := time.Parse(dateParserLayout, fmt.Sprintf("%s/%s", date, currentYear))
	return fDate
}
