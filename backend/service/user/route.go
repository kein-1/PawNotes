package user

import (
	"backend/types"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service types.UserStore
}

// basically the constructor
func NewUserHandler(s types.UserStore) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterRoutes(r *chi.Mux) {
	r.Post("/api/user/register", h.handleCreateUser)
	r.Get("/api/user/getUser", h.handleRetrieveUser)
	r.Delete("/api/user/deleteUser", h.Delete)
}

func (h *UserHandler) handleRetrieveUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Retrieving some user")
}

func (h *UserHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	var user types.User

	// check format
	if err := utils.ParseJSON(r, &user); err != nil {
		http.Error(w, "Invalid JSON data", 400)
		return
	}

	// check validator
	if err := utils.Validate.Struct(user); err != nil {

		// validation erros is actuall an array of []FieldErrors
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			utils.WriteError(w, validationErrors)
		}
	}
	fmt.Println("This is json user to struct", user)
	h.service.CreateUser(user)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting")
}
