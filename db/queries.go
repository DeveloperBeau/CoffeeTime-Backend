package db

import (
	"fmt"
)

func GetAllFromUser() string {
	return fmt.Sprintf("select * from %s", UserTable())
}