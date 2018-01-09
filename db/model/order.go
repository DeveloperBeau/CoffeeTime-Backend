package model

//Order data model
type Order struct {
	ID        string `json:"orderID"`
	UserID    string `json:"userID"`
	SessionID string `json:"sessionID"`
	Request   string `json:"request"`
}