package middlewares

import (
	"os"

	"github.com/go-chi/jwtauth/v5"
)

func GenerateJwtToken(payload map[string]interface{}) string {
	secretKey := os.Getenv("SECRET_KEY")
	tokenAuth := jwtauth.New("HS256", []byte(secretKey), nil) // will be move to env

	_, tokenString, _ := tokenAuth.Encode(payload)
	return tokenString
}
