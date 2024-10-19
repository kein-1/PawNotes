package utils

// mainly used for validator and parsing json etc

import (
	"backend/config"
	"backend/types"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	finalMap["message"] = errorMap
	fmt.Println("Map", errorMap)
	SendJSON(w, status, finalMap)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	errorMsg := make(map[string]string)
	errorMsg["message"] = err.Error() // calling .Error returns a string
	SendJSON(w, status, errorMsg)
}

func GenerateJWT(user *types.User) (string, error) {
	key := []byte(config.ConfigValues.SECRET_KEY)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  user.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenStr, err
}
