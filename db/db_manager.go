package db

import "CoffeeTime-Go/db/model"

type Handler interface {
	AddUser(model.User) error
}

//Database Handler factory function
func MakeDatabaseHandler(connection string, isProduction bool) (Handler, error) {
	return NewSQLHandler(connection, isProduction)
}
