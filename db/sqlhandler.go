package db

import (
	"CoffeeTime-Go/db/model"
	"database/sql"
)

type SQLHandler struct {
	*sql.DB
}

// AddUser: adds a user to the database
func (handler SQLHandler) AddUser(u model.User) error {
	return nil
}
