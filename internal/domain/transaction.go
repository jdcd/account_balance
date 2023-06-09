package domain

import "time"

// MovementType groups the valid types of movements
type MovementType string

const (
	Credit MovementType = "credit"
	Debit  MovementType = "debit"
)

// Transaction represent a formatted transaction, with its type of movement and value separate
type Transaction struct {
	Number   int
	Date     time.Time
	Movement MovementType
	Value    float32
}

// IgnoredTransaction represent a transaction without valid format, that was ignored
type IgnoredTransaction struct {
	ID          string
	Date        string
	Transaction string
	Reason      string
}
