package custommiddleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	auth "github.com/kein-1/pawnotes/internal/jwt"
	"github.com/kein-1/pawnotes/utils"

	"github.com/rs/zerolog/log"
)

type ContextString string

const UserKey ContextString = "user"

// Parse JWT: either as cookie or as auth header
func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		// cookie for future web
		// var jwtToken string
		// cookieVal, err := r.Cookie("access")

		// auth bearer via mobile apps
		bearerToken := r.Header.Get("Authorization")
		jwtToken, err := splitToken(bearerToken)
		if err != nil {
			log.Info().Err(err).Msg("User does not exist")
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Invalid Token").Error())
			return
		}

		userID, err := auth.ParseToken(jwtToken)
		if err != nil {
			log.Info().Err(err).Msg("Error parsing token; token is invalid")
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Invalid Token").Error())
			return
		}
		ctx := context.WithValue(r.Context(), UserKey, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func splitToken(token string) (string, error) {
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return "", fmt.Errorf("Incorrect token format")
	}

	reqToken := strings.TrimSpace(splitToken[1])

	fmt.Println(reqToken)
	return reqToken, nil
}
