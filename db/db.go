package db

type db struct {
	debug bool
}

func NEW() *db {
	return &db{
		debug: false,
	}
}

func (r *db) Select(query string) *dbSelect {
	return NewSelect(query)
}

func (r *db) Insert() *dbInsert {
	return NewInsert()
}

func (r *db) Update() *dbUpdate {
	return NewUpdate()
}

func (r *db) Delete() *dbDelete {
	return NewDelete()
}
