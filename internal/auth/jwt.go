package auth

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))