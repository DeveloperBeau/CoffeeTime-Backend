package db

// userTable: returns string to user table, whether it is production or development
func userTable() string {
	// TODO: implement production and dev table branching
	return "dev.\"user\""
}
