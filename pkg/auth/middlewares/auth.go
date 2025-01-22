package middleware

import (
	"context"
	"errors"
	db "github.com/krack8/lighthouse/pkg/auth/config"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"strings"
)

// AuthMiddleware is used to verify if a user is authenticated and authorized to access a route
func AuthMiddleware(route string, method string, next http.Handler) http.Handler {
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

		filter := bson.M{"username": claims.Username}
		// Check if filter is not nil
		if filter == nil {
			// Handle the error
			http.Error(w, "User not Found", http.StatusUnauthorized)
			return
		}

		// FindOne with error handling
		result := db.UserCollection.FindOne(context.Background(), filter)

		var user models.User
		if err := result.Decode(&user); err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				http.Error(w, "User not Found", http.StatusUnauthorized)
				return
			}
			http.Error(w, "User not Found", http.StatusUnauthorized)
			return
		}

		if user.UserType == models.RegularUser {
			// Collect all permissions of the user's roles
			var permissions []string
			for _, role := range user.Roles {
				for _, perm := range role.Permissions {
					permissions = append(permissions, perm.Route+":"+perm.Method)
				}
			}

			// Check if user has permission for the requested route and method
			if !services.CheckPermission(permissions, route, method) {
				http.Error(w, "Permission denied", http.StatusForbidden)
				return
			}
		}

		// Forward to the next handler if authentication and authorization pass
		next.ServeHTTP(w, r)
	})
}
