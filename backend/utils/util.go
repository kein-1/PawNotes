package utils

// mainly used for validator and parsing json etc

import (
	"backend/types"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New(validator.WithRequiredStructEnabled())

func ParseJSON(r *http.Request, user *types.User) error {
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return err
	}
	return nil
}

func SendJSON(w http.ResponseWriter, status int, content any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(content)
	return
}

func WriteError(w http.ResponseWriter, err validator.ValidationErrors) {
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

	fmt.Println("Map", errorMap)
	SendJSON(w, http.StatusBadRequest, errorMap)

	return
}
