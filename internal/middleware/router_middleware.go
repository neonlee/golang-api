package middleware

import (
	"net/http"
	"strings"
)

// AuthMiddleware checks for a valid Authorization header (e.g., Bearer token).
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		// TODO: Validate the token (e.g., JWT verification)
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
