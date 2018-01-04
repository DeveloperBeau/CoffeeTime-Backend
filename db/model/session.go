package model

import "time"

//Session data model
type Session struct {
	ID       int
	UserID   int
	IsActive bool
	Started  time.Time
	Ended    time.Time
	Orders   []Order
}
