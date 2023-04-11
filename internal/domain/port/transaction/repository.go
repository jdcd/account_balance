package transaction

// Repository describes the behavior, for save  account transactions in a persistent repository
type Repository interface {
	Save()
}
