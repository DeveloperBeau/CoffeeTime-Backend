package model

//Order data model
type Order struct {
	ID        int
	UserID    int
	SessionID int
	Request   string
}
