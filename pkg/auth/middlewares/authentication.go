package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/config"
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
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix from the token if present
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}

		// Validate token and extract claims
		claims, err := utils.ValidateToken(token, os.Getenv("JWT_SECRET"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization token"})
			c.Abort()
			return
		}

		filter := bson.M{"username": claims.Username}
		// Check if filter is not nil
		if filter == nil {
			// Handle the error
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// FindOne with error handling
		result := config.UserCollection.FindOne(context.Background(), filter)

		var user models.User
		if err := result.Decode(&user); err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}
		// store username
		c.Set("username", claims.Username)
		if user.UserType == models.RegularUser {
			// Collect all permissions of the user's roles
			var permissionEndpoints []string
			for _, role := range user.Roles {
				for _, perm := range role.Permissions {
					// Iterate through each endpoint in the EndpointList
					for _, endpoint := range perm.EndpointList {
						permissionEndpoints = append(permissionEndpoints,
							endpoint.Route+":"+endpoint.Method)
					}
				}
			}

			// Check if user has permission for the requested route and method
			if !services.CheckPermission(permissionEndpoints, c.FullPath(), c.Request.Method) {
				c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
				c.Abort()
				return
			}
		}

		// Proceed to the next handler if authentication and authorization pass
		c.Next()
	}
}
