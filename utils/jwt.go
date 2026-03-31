package utils

import (
	"errors"
	"time"

	"github.com/Harsh00067/harshad-api/config"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte(config.GetJWTSecret())

func GenerateToken(username string, expiry time.Duration) (string, error) {

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(expiry).Unix(),
		"role":     "admin",
		"iss":      "harshad-api",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(SecretKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return SecretKey, nil
	})
}
