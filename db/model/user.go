package model

// User Struct
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	IsEnabled bool   `json:"enabled"`
}
