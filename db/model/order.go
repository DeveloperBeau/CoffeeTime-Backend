package model

//Order data model
type Order struct {
	ID        int `json:"orderID"`
	UserID    int `json:"userID"`
	SessionID int `json:"sessionID"`
	Request   string `json:"request"`
}