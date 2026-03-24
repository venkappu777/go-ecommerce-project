package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, email string) (string, error) {

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
	})

	// token header is HS256 and JWT and payload is "user_id and email"

	// Get secret from .env
	secret := os.Getenv("JWT_SECRET")

	// sign token
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	// after signing it will be like header.payload.signature

	return tokenString, nil

}
