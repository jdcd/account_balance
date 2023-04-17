package domain

import "time"

type SuccessReport struct {
	FileName           string
	Date               time.Time
	Summary            Summary
	SendTo             []string
	Transactions       []Transaction
	IgnoredTransaction []IgnoredTransaction
}

type ErrorReport struct {
	FileName string
	Date     time.Time
	Error    string
}
