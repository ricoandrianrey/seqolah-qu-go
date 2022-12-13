package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) string {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashPassword)
}

func ComparePassword(password string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
