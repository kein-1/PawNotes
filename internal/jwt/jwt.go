package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kein-1/pawnotes/config"
)

func GenerateToken(userID int) (string, error) {

	claims := &jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userID), // store it as a string; convert to int later in parsing
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		Issuer:    "pawnotes",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.ConfigValues.JWT_SECRET))
	if err != nil {
		return "", fmt.Errorf("Error creating jwt token; %w", err)
	}
	return tokenString, nil
}

func ParseToken(token string) (int, error) {

	claims := &jwt.RegisteredClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ConfigValues.JWT_SECRET), nil
	})

	if err != nil {
		return -1, fmt.Errorf("failed to parse token: %w", err)
	}

	if !jwtToken.Valid {
		return -1, fmt.Errorf("Invalid token %w", err)
	}
	fmt.Println("claims: after parsing", claims)

	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, err
	}
	return id, nil
}
