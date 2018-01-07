package db

import (
	"database/sql"

	//Importing postgresql framework
	_ "github.com/lib/pq"
)

// NewSQLHandler : Factory method generating SQLHandler struct
func NewSQLHandler(connection string, isProduction bool) (*SQLHandler, error) {
	db, err := sql.Open("postgres", connection)
	qm := getQueryManager(isProduction)
	return &SQLHandler{
		DB:           db,
		queryManager: qm,
	}, err
}
