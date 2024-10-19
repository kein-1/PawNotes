package middleware

import (
	"context"
	"net/http"
)

func TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user2", "1234")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
