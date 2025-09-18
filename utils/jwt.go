package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey string

func GenerateToken(email string, userid int64) (string, error) {

	secretKey = os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userid,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))

}

func VerifyToken(tokenString string) error {
	secretKey = os.Getenv("JWT_SECRET")

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Make sure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return err // parsing failed
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	return nil

}
