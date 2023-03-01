package utils

import (
	uuid "github.com/satori/go.uuid"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID() string {
	id:= uuid.NewV4()
	return id.String()
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func MakeFilters(filters map[string]interface{}) (string, []interface{}) {
    var query string
    var args []interface{}
    for key, value := range filters {
        if value != "" {
            query += fmt.Sprintf(" AND %s = ?", key)
            args = append(args, value)
        }
    }
    return query, args
}