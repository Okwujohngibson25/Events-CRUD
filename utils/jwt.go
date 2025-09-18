package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var secretKey string

func GenerateToken(email string, userid int64) (string, error) {

	// secretKey := os.Getenv("JWT_SECRET")

	secretKey := "Johndev"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userid,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))

}

func VerifyToken(tokenString string) (int64, error) {
	// secretKey = os.Getenv("JWT_SECRET")
	secretKey := "Johndev"

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Make sure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err // parsing failed
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	rawUserID, ok := claims["userId"]
	if !ok || rawUserID == nil {
		return 0, errors.New("userID not found in token claims")
	}

	fmt.Printf("rawUserID value: %#v (type %T)\n", claims, rawUserID)

	userIDFloat, ok := rawUserID.(float64)
	if !ok {
		return 0, fmt.Errorf("userID claim is not a number (got type %T)", rawUserID)
	}

	return int64(userIDFloat), nil

}
