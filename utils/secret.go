package utils

import (
	"time"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var SecretKey []byte

func init() {
	secret := os.Getenv("JWT_SECRET")
	SecretKey = []byte(secret)
}
func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}
