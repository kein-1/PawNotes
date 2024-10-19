package user

import (
	"backend/config"
	"backend/types"
	"backend/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService types.UserServiceInterface
}

// basically the constructor
func NewUserHandler(s types.UserServiceInterface) *UserHandler {
	return &UserHandler{userService: s}
}

func (h *UserHandler) RegisterRoutes(r *chi.Mux) {
	// setup main router + subrouter routes

	tokenAuth := jwtauth.New("HS256", []byte(config.ConfigValues.SECRET_KEY), nil)

	r.Route("/api/user", func(r chi.Router) {
		r.Post("/register", h.handleCreateUser)
		r.Post("/login", h.handleUserLogin)

		r.Group(func(r chi.Router) {

			// protected routes
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Delete("/deleteUser", h.handleDeleteUser)
		})
	})
}

func (h *UserHandler) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.UserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, 400, err)
		return
	}

	jwt, err := h.userService.LoginUser(payload)
	if err != nil {
		if err == types.ErrInvalidUserCredentials {
			utils.WriteError(w, http.StatusUnauthorized, err)
			return
		}
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    jwt,
		MaxAge:   3600 * 24,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	utils.SendJSON(w, 200, map[string]string{"message": "Successful login!"})

}

func (h *UserHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	// Basic data parsing
	var user types.User

	// check format
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, 400, err)
		return
	}
	// check validator
	if err := utils.Validate.Struct(user); err != nil {
		// validation errors is actually an array of []FieldErrors
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			utils.WriteValidationError(w, http.StatusBadRequest, validationErrors)
			return
		}
		return
	}

	// Hand off to service layer
	if err := h.userService.RegisterUser(user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.SendJSON(w, 200, map[string]string{"message": "Account successfully created"})
}

func (h *UserHandler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	// pass jwt to it

	var random_id string
	err := json.NewDecoder(r.Body).Decode(&random_id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	id, _ := strconv.Atoi(random_id)
	err = h.userService.RemoveUser(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.SendJSON(w, 200, map[string]string{"message": "Successfully deleted account"})

}
