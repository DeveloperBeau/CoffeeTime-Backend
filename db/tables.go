package db

// UserTable: returns string to user table, whether it is production or development
func UserTable() string {
	// TODO: implement production and dev table branching
	return "dev.\"user\""
}
