package worker

import "database/sql"

type ExpiredUrlCleaner struct {
	db *sql.DB
}

func NewExpiredUrlCleaner(db *sql.DB) Worker {
	return &ExpiredUrlCleaner{db: db}
}

// process implements IExpiredUrlCleaner.
func (w *ExpiredUrlCleaner) Process() {
	panic("unimplemented")
}
