package db

import (
	"database/sql"

	//Importing postgresql framework
	_ "github.com/lib/pq"
)

func NewSQLHandler(connection string, isProduction bool) (*SQLHandler, error) {
	db, err := sql.Open("postgres", connection)
	qm := GetQueryManager(isProduction)
	return &SQLHandler{
		DB: db,
		QueryManager: qm,
	}, err
}
