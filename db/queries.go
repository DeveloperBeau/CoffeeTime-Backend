package db

import (
	"fmt"
)

func getAllFromUser() string {
	return fmt.Sprintf("select * from %s", userTable())
}

func getAllFromSession() string {
	return fmt.Sprintf("select * from %s", sessionTable())
}

func getAllFromOrder() string {
	return fmt.Sprintf("select * from %s", orderTable())
}