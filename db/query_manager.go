package db

import (
	"bytes"
	"fmt"
	"sync"
	"CoffeeTime-Go/db/model"
)

type queryManager struct {
	isProduction bool
}

var instance *queryManager
var once sync.Once

func getQueryManager(isProduction bool) *queryManager {
	once.Do(func() {
		instance = &queryManager{isProduction: isProduction}
	})
	return instance
}

func (q queryManager) getAllFromUserWithEmail(e string) string {
	return getAllFromUser() + whereClause(false, where{field: "email", value: e})
}

func (q queryManager) getCurrentSession() string {
	return getAllFromSession() + whereClause(false, where{field: "is_active", value: true})
}

func (q queryManager) getUserCurrentOrder(UID string, SID string) string {
	return getAllFromOrder() + whereClause(false, where{field: "user_id", value: UID}, where{field: "session_id", value: SID})
}

type where struct {
	field string
	value interface{}
}

// whereClause:
// variables: o - Optional, w - splice of where structs
// this function returns a structured where extension for queries
func whereClause(o bool, w ...where) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(" where "))
	for c, wc := range w {
		if c == 0 {
			buffer.WriteString(fmt.Sprintf("%s='%v' ", wc.field, wc.value))
		} else {
			if o {
				buffer.WriteString(fmt.Sprintf("OR %s='%v' ", wc.field, wc.value))
			} else {
				buffer.WriteString(fmt.Sprintf("AND %s='%v' ", wc.field, wc.value))
			}
		}
	}
	return buffer.String()
}
