package db

import (
	"CoffeeTime-Go/db/model"
	"database/sql"
	"fmt"
)

type SQLHandler struct {
	*sql.DB
}

// AddUser: adds a user to the database
func (handler SQLHandler) Add(u model.User) error {
	existingUser, err := handler.getUser(u.Email)
	if existingUser == nil {
		_, err = handler.Exec(fmt.Sprintf("Insert into user (first_name, last_name, email, auth_token) values ('%s','%s','%s','%s')", u.FirstName, u.LastName, u.Email, u.Token))
	}
	return err
}

func (handler *SQLHandler) getUser(e string) (*model.User, error) {
	var user model.User
	queryErr := handler.DB.QueryRow("select * from user where email = $1", e).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Token, &user.IsEnabled)
	if queryErr == nil {
		return &user, nil
	} else {
		return nil, queryErr
	}
}
