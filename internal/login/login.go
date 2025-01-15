// Specific HTTP route for handling login since setting cookies
// proved to be difficult
package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kein-1/pawnotes/ent"
	"github.com/kein-1/pawnotes/ent/user"
	auth "github.com/kein-1/pawnotes/internal/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	db *ent.Client
}

func NewAuthHandler(db *ent.Client) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) RegisterRoute(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.handleLogin)
	})
}

func (h *AuthHandler) handleLogin(w http.ResponseWriter, r *http.Request) {

	// have http.HandlerFunc, which wraps
	// if we have the signature w http.ResponseWriter, r *http.Request, it
	// gets converted;
	// also, know diff between http.Handle and HandlerFunc!

	var payload LoginPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, map[string]string{"message": fmt.Errorf("Error in parsing payload %w", err).Error()})
		return
	}

	user, err := h.db.User.Query().Where(user.Email(payload.Email)).First(r.Context())
	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, map[string]string{"message": fmt.Errorf("User does not exist. %w", err).Error()})
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		WriteJSON(w, http.StatusUnauthorized, map[string]string{"message": fmt.Errorf("Incorrect password. %w", err).Error()})
		return
	}

	// generate token, set cookie

	accessToken, err := auth.GenerateToken(user.ID)
	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, map[string]string{"message": fmt.Errorf("Error generating token. %w", err).Error()})
		return
	}

	accessCookie := http.Cookie{
		Name:     "access",
		Value:    accessToken,
		MaxAge:   60 * 15,
		HttpOnly: true,
		Secure:   false, // change in prod
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	http.SetCookie(w, &accessCookie)
	WriteJSON(w, 200, "successfully logged in")

}

func WriteJSON(w http.ResponseWriter, status int, content any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(content)
	return err
}
