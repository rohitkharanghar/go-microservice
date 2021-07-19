package dataaccess

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func (s *Store) NewStore() *Store {
	db, err := sqlx.Open("postgres", "postgres://user:pass@localhost/db")
	if err != nil {
		log.Fatal(err)
	}

	// Set the maximum number of concurrently open connections (in-use + idle)
	// to 5. Setting this to less than or equal to 0 will mean there is no
	// maximum limit (which is also the default setting).
	db.SetMaxOpenConns(5)
	s.db = db
	return s
}
