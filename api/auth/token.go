package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(user_id uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour + 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	api_secret := os.Getenv("API_SECRET")
	tokenString, err := token.SignedString([]byte(api_secret))
	return tokenString, err
}
