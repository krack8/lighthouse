package routes

import (
	"github.com/gorilla/mux"
	"github.com/krack8/lighthouse/pkg/auth/controllers"
)

// InitAuthRoutes initializes authentication-related routes
func InitAuthRoutes(router *mux.Router) {

	// Create a sub-router for auth-related endpoints
	authRouter := router.PathPrefix("/auth").Subrouter()

	// Define routes
	authRouter.HandleFunc("/login", controllers.LoginHandler).Methods("POST")                // Login route
	authRouter.HandleFunc("/refresh-token", controllers.RefreshTokenHandler).Methods("POST") // Refresh token route
}
