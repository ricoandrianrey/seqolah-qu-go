package middlewares

import (
	"net/http"
	"os"

	"github.com/go-chi/jwtauth/v5"
)

type ctxKey struct {
	key string
}

func Jwt() *jwtauth.JWTAuth {
	secretKey := os.Getenv("SECRET_KEY")
	jwt := jwtauth.New("HS256", []byte(secretKey), nil) // will be move to env
	return jwt
}

func GenerateJwtToken(payload map[string]interface{}) string {
	jwt := Jwt()
	_, tokenString, _ := jwt.Encode(payload)
	return tokenString
}

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt := Jwt()
		tokenString := jwtauth.TokenFromHeader(r)

		if tokenString == "" {
			http.Error(w, "Token not provided", http.StatusUnauthorized)
			return
		}

		_, err := jwtauth.VerifyToken(jwt, tokenString)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
