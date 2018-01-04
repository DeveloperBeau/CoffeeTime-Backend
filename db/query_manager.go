package db

import (
	"bytes"
	"fmt"
	"sync"
)

type QueryManager struct {
	isProduction bool
}

var instance *QueryManager
var once sync.Once

func GetQueryManager(isProduction bool) *QueryManager {
	once.Do(func() {
		instance = &QueryManager{isProduction: isProduction}
	})
	return instance
}

func (q QueryManager) getAllFromUserWithEmail(e string) string {
	return GetAllFromUser() + whereClause(false, Where{field: "email", value: e})
}

type Where struct {
	field string
	value interface{}
}

// whereClause:
// variables: o - Optional, w - splice of where structs
// this function returns a structured where extension for queries
func whereClause(o bool, w ...Where) string {
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
