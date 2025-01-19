// Specific HTTP route for handling login since setting cookies
// proved to be difficult
package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kein-1/pawnotes/ent"
	"github.com/kein-1/pawnotes/ent/user"
	auth "github.com/kein-1/pawnotes/internal/jwt"
	"github.com/kein-1/pawnotes/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/rs/zerolog/log"
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
		log.Info().Err(err).Msg("Error in parsing payload")
		utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{"message": fmt.Errorf("Incorrect username or password").Error()})
		return
	}

	user, err := h.db.User.Query().Where(user.Email(payload.Email)).First(r.Context())
	if err != nil {
		log.Info().Err(err).Msg("User does not exist")
		utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{"message": fmt.Errorf("Incorrect username or password").Error()})
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		log.Info().Err(err).Msg("Incorrect password; failed to hash properly")
		utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{"message": fmt.Errorf("Incorrect username or password").Error()})
		return
	}

	// generate token, set cookie
	accessToken, err := auth.GenerateToken(user.ID)
	if err != nil {
		log.Info().Err(err).Msg("Error generating token")
		utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{"message": fmt.Errorf("Incorrect username or password").Error()})
		return
	}

	accessCookie := http.Cookie{
		Name:     "access",
		Value:    accessToken,
		MaxAge:   int((time.Hour * 24 * 7).Seconds()), // TODO: change in prod
		HttpOnly: true,
		Secure:   false, // TODO: change in prod
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	http.SetCookie(w, &accessCookie)

	// also set cookie in auth header
	w.Header().Set("Access-Token", accessToken)

	utils.WriteJSON(w, 200, "successfully logged in")

}
