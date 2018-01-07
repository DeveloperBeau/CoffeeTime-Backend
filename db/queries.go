package db

import (
	"fmt"
)

func getAllFromUser() string {
	return fmt.Sprintf("select * from %s", userTable())
}
