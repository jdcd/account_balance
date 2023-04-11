package notification

// Sender contains the contracts for notifying the result of the balancing process.
type Sender interface {
	SendResult()
}
