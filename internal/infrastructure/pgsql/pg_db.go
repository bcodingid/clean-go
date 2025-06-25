package pgsql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=clean-db host=localhost sslmode=disable")

	if err != nil {
		return nil, err
	}

	println("PG Database connection established")
	return db, nil
}
