package model

import "time"

//Session data model
type Session struct {
	ID       string
	UserID   string
	IsActive bool
	Started  time.Time
	Ended    time.Time
	Orders   []Order
}
