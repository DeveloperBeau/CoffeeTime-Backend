package db

import (
	"CoffeeTime-Go/db/model"
	"database/sql"
	"errors"
	"fmt"
)

//SQLHandler : Generic DB handler struct
type SQLHandler struct {
	*sql.DB
	*queryManager
}

var (
	//ErrUserExists : Generic error for user already existing
	ErrUserExists = errors.New("error - User Already Exists")
)

// AddUser : adds a user to the database
func (handler SQLHandler) AddUser(u model.User) error {
	existingUser, err := handler.getUser(u.Email)
	if existingUser == nil {
		_, err = handler.Exec(fmt.Sprintf("Insert into %s (first_name,last_name,email,auth_token) values ('%s','%s','%s','%s')", userTable(), u.FirstName, u.LastName, u.Email, u.Token))
	}
	return err
}

func (handler *SQLHandler) getUser(e string) (*model.User, error) {
	var user model.User
	q := handler.queryManager.getAllFromUserWithEmail(e)
	queryErr := handler.DB.QueryRow(q).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Token, &user.IsEnabled)
	if queryErr == nil {
		return &user, ErrUserExists
	}
	return nil, nil
}

// StartSession : Starts a new session based on a session model
func (handler *SQLHandler) StartSession(s model.Session) (string, error) {
	return "", nil
}

// EndSession : Ends the session if the UID firing the ending of the session is the one that created it.
func (handler *SQLHandler) EndSession(UID string) error {
	return nil
}

// Session : Get current session information for the current open session if there is one else return nil
func (handler *SQLHandler) Session(UID string) *model.Session {
	return nil
}

// Order : Place an order for the current session and return its order id else return an error
func (handler *SQLHandler) Order(o model.Order) (string, error) {
	return "", nil
}
