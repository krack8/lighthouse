package middleware

import (
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"net/http"
	"os"
	"strings"
)

// AuthMiddleware is used to verify if a user is authenticated and authorized to access a route
func AuthMiddleware(route string, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization token required", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix from the token if present
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
			return
		}

		// Validate token and extract claims
		claims, err := utils.ValidateToken(token, os.Getenv("JWT_SECRET"))
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if user has permission for the requested route and method
		if !services.CheckPermission(claims.Permissions, route, method) {
			http.Error(w, "Permission denied", http.StatusForbidden)
			return
		}

		// Proceed with the next handler if authentication and authorization pass
		// Here `next.ServeHTTP(w, r)` is skipped because it's implicitly managed
		// through the `http.HandlerFunc`.
	})
}
