package db

import (
	"database/sql"

	//Importing postgresql framework
	_ "github.com/lib/pq"
)

func NewSQLHandler(connection string) (*SQLHandler, error) {
	db, err := sql.Open("postgres", connection)
	return &SQLHandler{
		DB: db,
	}, err
}
