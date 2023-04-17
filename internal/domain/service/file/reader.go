package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jdcd/account_balance/internal/domain"
	"github.com/jdcd/account_balance/pkg"
)

// IReader contracts related to read report files
type IReader interface {
	ReadFile(fileName string) ([]domain.Transaction, []domain.IgnoredTransaction, error)
}

const (
	currentYear                   = "2023" // To complete the challenge, a date needs a year
	dateParserLayout              = "1/2/2006"
	emptyFileError                = "the file %s is empty"
	invalidTransactionFormatError = "report value \"%s\" has incorrect format"
	openingFileError              = "the file %s could not be opened: %s\n"
	readingFileError              = "the file %s could not be reading: %s\n"
	unclosedFileError             = "the file %s could not be closed: %s\n"
)

// ReaderService implements IReader reading for a csv file on current file system.
type ReaderService struct{}

// ReadFile reads a csv file of transactions and discriminates them by admitted or discarded
func (s *ReaderService) ReadFile(fileName string) ([]domain.Transaction, []domain.IgnoredTransaction, error) {
	records, err := s.getRecords(fileName)
	if err != nil {
		return nil, nil, err
	}

	if len(records) < 2 {
		return nil, nil, fmt.Errorf(emptyFileError, fileName)
	}

	transactions := make([]domain.Transaction, 0)
	ignored := make([]domain.IgnoredTransaction, 0)

	for i, record := range records { // for large data, implementing two traditional "for" can improve performance
		if i == 0 { // avoid csv header
			continue
		}

		tr, err := s.formatTransaction(record)
		if err != nil {
			ig := domain.IgnoredTransaction{
				ID:          record[0],
				Date:        record[1],
				Transaction: record[2],
				Reason:      err.Error(),
			}
			ignored = append(ignored, ig)
			continue
		}
		transactions = append(transactions, tr)
	}

	return transactions, ignored, nil
}

func (s *ReaderService) formatTransaction(record []string) (domain.Transaction, error) {
	number, err := strconv.Atoi(record[0])
	if err != nil {
		return domain.Transaction{}, err
	}

	date, err := time.Parse(dateParserLayout, fmt.Sprintf("%s/%s", record[1], currentYear))
	if err != nil {
		return domain.Transaction{}, err
	}

	movementType, value, err := s.parseValue(record[2])
	if err != nil {
		return domain.Transaction{}, err
	}

	tr := domain.Transaction{
		Number:   number,
		Date:     date,
		Movement: movementType,
		Value:    value,
	}

	return tr, nil
}

func (s *ReaderService) getRecords(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			pkg.WarningLogger().Printf(unclosedFileError, fileName, err)
		}
	}(file)

	if err != nil {
		pkg.ErrorLogger().Printf(openingFileError, fileName, err)
		return nil, err
	}

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.Comment = '#'
	records, err := reader.ReadAll()
	if err != nil {
		pkg.ErrorLogger().Printf(readingFileError, fileName, err)
		return nil, err
	}

	return records, nil
}

func (s *ReaderService) parseValue(tr string) (domain.MovementType, float32, error) {
	if len(tr) < 2 {
		return "", 0, fmt.Errorf(invalidTransactionFormatError, tr)
	}

	var movementType domain.MovementType
	switch string(tr[0]) {
	case "+":
		movementType = domain.Credit
	case "-":
		movementType = domain.Debit
	default:
		return "", 0, fmt.Errorf(invalidTransactionFormatError, tr)
	}

	sValue := tr[1:]
	value, err := strconv.ParseFloat(sValue, 32)
	if err != nil {
		return "", 0, fmt.Errorf(invalidTransactionFormatError, tr)
	}

	return movementType, float32(value), nil
}
