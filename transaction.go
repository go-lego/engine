package engine

// Transaction interface
type Transaction interface {
	Begin()
	Commit()
	Rollback()
}
