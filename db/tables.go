package db

// userTable: returns string to user table, whether it is production or development
func userTable() string {
	// TODO: implement production and dev table branching
	return "dev.\"user\""
}

func sessionTable() string {
	// TODO: implement production and dev table branching
	return "dev.session"
}

func orderTable() string {
	// TODO: implement production and dev table branching
	return "dev.order"
}
