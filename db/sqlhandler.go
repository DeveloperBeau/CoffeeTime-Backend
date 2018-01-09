package db

import (
	"CoffeeTime-Go/api"
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
	ErrUserNotFound = errors.New("error - User doesn't exist")
	//ErrUserExists : Generic error for user already existing
	ErrUserExists = errors.New("error - User Already Exists")
	//ErrSessionExists : Generic error for session already existing
	ErrSessionExists = errors.New("error - Session Already Exists")
)

// AddUser : adds a user to the database
func (handler SQLHandler) AddUser(u model.User) error {
	existingUser := handler.getUser(u.Email)
	if existingUser == nil {
		_, err := handler.Exec(fmt.Sprintf("Insert into %s (first_name,last_name,email,auth_token) values ('%s','%s','%s','%s')", userTable(), u.FirstName, u.LastName, u.Email, u.Token))
		return err
	}
	return ErrUserExists
}

// getUser : received e (email) and gets the user based on the email provided.
func (handler *SQLHandler) getUser(e string) *model.User {
	var user model.User
	q := handler.queryManager.getAllFromUserWithEmail(e)
	queryErr := handler.DB.QueryRow(q).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Token, &user.IsEnabled)
	if queryErr == nil {
		return &user
	}
	return nil
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
	currentSession := handler.Session()
	if currentSession == nil {
		return api.ErrSessionNotFound
	}
	if currentSession.UserID == UID {
		return endSession(*currentSession, handler)
	}
	return api.ErrSomethingWentWrong
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
	currentSession := handler.Session()
	currentOrder := getOrder(currentSession.ID, o.UserID, handler)
	o.SessionID = currentSession.ID
	if currentSession == nil {
		return "", api.ErrSessionNotFound
	} else if currentOrder != nil {
		updateOrder(o, handler)
		return currentOrder.ID, nil
	}

	return "", nil
}

func getOrder(SID string, UID string, handler *SQLHandler) *model.Order {
	var currentOrder model.Order
	q := handler.queryManager.getUserCurrentOrder(UID, SID)
	queryErr := handler.DB.QueryRow(q).Scan(&currentOrder.ID, &currentOrder.UserID, &currentOrder.SessionID, &currentOrder.Request)
	if queryErr != nil {
		return nil
	}
	return &currentOrder
}

func createOrder(o model.Order, handler *SQLHandler) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into %s (user_id,session_id,request) values (%s,%s,'%s')", orderTable(), o.UserID, o.SessionID, o.Request))
	return err
}

func updateOrder(o model.Order, handler *SQLHandler) error {
	_, err := handler.Exec(fmt.Sprintf("UPDATE %s SET request = \"%s\" WHERE session_id = %s AND user_id = %s", orderTable(), o.Request, o.SessionID, o.UserID))
	return err
}

func createSession(s model.Session, handler *SQLHandler) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into %s (user_id) values (%s)", sessionTable(), s.UserID))
	return err
}

func endSession(s model.Session, handler *SQLHandler) error {
	_, err := handler.Exec(fmt.Sprintf("UPDATE %s SET is_active = FALSE WHERE id = %s", sessionTable(), s.ID))
	return err
}
