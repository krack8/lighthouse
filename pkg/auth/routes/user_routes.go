package routes

import (
	"github.com/gorilla/mux"
	"github.com/krack8/lighthouse/pkg/auth/controllers"
	middleware "github.com/krack8/lighthouse/pkg/auth/middlewares"
)

// InitUserRoutes initializes user-related routes
func InitUserRoutes(router *mux.Router) {

	// Create a sub-router for user-related endpoints
	userRouter := router.PathPrefix("/users").Subrouter()

	// Define routes with authentication middleware

	// Get all users - Protected
	userRouter.HandleFunc("/", controllers.GetAllUsersHandler).Methods("GET").Handler(middleware.AuthMiddleware("/users", "GET"))

	// Get user by ID - Protected
	userRouter.HandleFunc("/{id}", controllers.GetUserHandler).Methods("GET").Handler(middleware.AuthMiddleware("/users/{id}", "GET"))

	// Create a new user - Protected
	userRouter.HandleFunc("/", controllers.CreateUserHandler).Methods("POST").Handler(middleware.AuthMiddleware("/users", "POST"))

	// Update user by ID - Protected
	userRouter.HandleFunc("/{id}", controllers.UpdateUserHandler).Methods("PUT").Handler(middleware.AuthMiddleware("/users/{id}", "PUT"))

	// Delete user by ID - Protected
	userRouter.HandleFunc("/{id}", controllers.DeleteUserHandler).Methods("DELETE").Handler(middleware.AuthMiddleware("/users/{id}", "DELETE"))
}
