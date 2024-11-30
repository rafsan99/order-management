package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(5 * 24 * time.Hour).Unix(), // expires in 5 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
