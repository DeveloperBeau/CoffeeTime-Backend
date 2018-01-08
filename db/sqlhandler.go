package db

import (
	"CoffeeTime-Go/db/model"
	"database/sql"
	"errors"
	"fmt"
	"CoffeeTime-Go/api"
)

//SQLHandler : Generic DB handler struct
type SQLHandler struct {
	*sql.DB
	*queryManager
}

var (
	//ErrUserExists : Generic error for user already existing
	ErrUserExists = errors.New("error - User Already Exists")
	//ErrSessionExists : Generic error for session already existing
	ErrSessionExists = errors.New("error - Session Already Exists")

)

// AddUser : adds a user to the database
func (handler SQLHandler) AddUser(u model.User) error {
	existingUser, err := handler.getUser(u.Email)
	if existingUser == nil {
		_, err = handler.Exec(fmt.Sprintf("Insert into %s (first_name,last_name,email,auth_token) values ('%s','%s','%s','%s')", userTable(), u.FirstName, u.LastName, u.Email, u.Token))
	}
	return err
}

// getUser : received e (email) and gets the user based on the email provided.
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
	currentSession := handler.Session()
	if currentSession != nil {
		return currentSession.ID, ErrSessionExists
	}
	err := createSession(s, handler)
	if err != nil {
		currentSession := handler.Session()
		if currentSession != nil {
			return currentSession.ID, nil
		}
	}
	return "", api.ErrSomethingWentWrong
}

// EndSession : Ends the session if the UID firing the ending of the session is the one that created it.
func (handler *SQLHandler) EndSession(UID string) error {
	return nil
}

// Session : Get current session information for the current open session if there is one else return nil
func (handler *SQLHandler) Session() *model.Session {
	var currentSession model.Session
	q := handler.queryManager.getCurrentSession()
	queryErr := handler.DB.QueryRow(q).Scan(&currentSession.ID, &currentSession.UserID, &currentSession.IsActive, &currentSession.Started, &currentSession.Ended)
	if currentSession.ID != "" && queryErr == nil {
		return &currentSession
	}
	return nil
}

// Order : Place an order for the current session and return its order id else return an error
func (handler *SQLHandler) Order(o model.Order) (string, error) {
	return "", nil
}

func createSession(s model.Session, handler *SQLHandler) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into %s (user_id) values ('%s')", sessionTable(), s.UserID))
	return err
}