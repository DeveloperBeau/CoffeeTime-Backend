package db

import (
	"CoffeeTime-Go/db/model"
	"database/sql"
	"errors"
	"fmt"
)

type SQLHandler struct {
	*sql.DB
	*QueryManager
}

var (
	ErrUserExists = errors.New("error - User Already Exists")
)

// AddUser : adds a user to the database
func (handler SQLHandler) AddUser(u model.User) error {
	existingUser, err := handler.getUser(u.Email)
	if existingUser == nil {
		_, err = handler.Exec(fmt.Sprintf("Insert into %s (first_name,last_name,email,auth_token) values ('%s','%s','%s','%s')", UserTable(), u.FirstName, u.LastName, u.Email, u.Token))
	}
	return err
}

func (handler *SQLHandler) getUser(e string) (*model.User, error) {
	var user model.User
	q := handler.QueryManager.getAllFromUserWithEmail(e)
	queryErr := handler.DB.QueryRow(q).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Token, &user.IsEnabled)
	if queryErr == nil {
		return &user, ErrUserExists
	} else {
		return nil, nil
	}
}

func (handler *SQLHandler) StartSession(model.Session) (string, error) {
	return "", nil
}
func (handler *SQLHandler) EndSession() error {
	return nil
}
func (handler *SQLHandler) Session() *model.Session {
	return nil
}
func (handler *SQLHandler) Order(model.Order) error {
	return nil
}
