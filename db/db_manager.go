package db

import "github.com/DeveloperBeau/CoffeeTime-Go/db/model"

// Handler interface which allows other packages to use the database
type Handler interface {
	AddUser(model.User) error
	StartSession(model.Session) (string, error)
	EndSession(string) error
	Session() *model.Session
	Order(model.Order) (string, error)
}

// MakeDatabaseHandler factory function
func MakeDatabaseHandler(connection string, isProduction bool) (Handler, error) {
	return NewSQLHandler(connection, isProduction)
}
