package utils

// mainly used for validator and parsing json etc

import (
	"backend/config"
	"backend/types"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

var Validate = validator.New(validator.WithRequiredStructEnabled())

func ParseJSON(r *http.Request, object any) error {
	err := json.NewDecoder(r.Body).Decode(object)
	if err != nil {
		return err
	}
	return nil
}

func SendJSON(w http.ResponseWriter, status int, content any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(content)
}

func WriteValidationError(w http.ResponseWriter, status int, err validator.ValidationErrors) {
	errorMap := make(map[string]string)

	for _, val := range err {
		field := val.Field()
		tag := val.Tag()

		var message string

		switch tag {
		case "required":
			message = fmt.Sprintf("%s is required", field)
		case "email":
			message = fmt.Sprintf("%s is incorrect email format", field)
		case "min":
			message = fmt.Sprintf("%s must between 3 and 15 characters", field)
		case "max":
			message = fmt.Sprintf("%s must between 3 and 15 characters", field)
		default:
			message = fmt.Sprintf("%s is invalid", field)
		}
		errorMap[field] = message
	}

	finalMap := make(map[string]map[string]string)
	finalMap["errors"] = errorMap
	fmt.Println("Map", errorMap)
	SendJSON(w, status, finalMap)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	errorMsg := map[string]string{
		"error": err.Error(),
	}
	SendJSON(w, status, errorMsg)
}

func GenerateJWT(user *types.User) (string, error) {
	key := []byte(config.ConfigValues.SECRET_KEY)

	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	jwtauth.SetExpiry(claims, time.Now().Add(time.Hour*24))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenStr, err
}

func ExtractIDFromJWT(r *http.Request) (int, error) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	userID, ok := claims["user_id"]
	userIDFloat, ok := userID.(float64)
	if !ok {
		return -1, fmt.Errorf("Failed to cast")
	}
	userIDInt := int(userIDFloat)
	return userIDInt, nil
}
