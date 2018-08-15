package tx

// Transaction interface
type Transaction interface {
	Begin()
	Commit()
	Rollback()
}
