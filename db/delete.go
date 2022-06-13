package db

type dbDelete struct{}

func NewDelete() *dbDelete {
	return &dbDelete{}
}
