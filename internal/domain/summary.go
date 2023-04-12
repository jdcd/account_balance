package domain

import "time"

type Summary struct {
	Total              float32
	TransactionByMonth map[time.Month]int
	AvrDebitAmount     float32
	AvrCreditAmount    float32
}
