package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kein-1/pawnotes/config"
)

func GenerateToken(userID int) (string, error) {

	claims := jwt.MapClaims{}
	claims["id"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.ConfigValues.JWT_SECRET))
	if err != nil {
		return "", fmt.Errorf("Error creating jwt token; %w", err)
	}
	return tokenString, nil
}

func ParseToken(jwt string) (int, error) {

	return 0, nil
}
