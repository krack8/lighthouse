package routes

import (
	"github.com/gorilla/mux"
	"github.com/krack8/lighthouse/pkg/auth/controllers"
	middleware "github.com/krack8/lighthouse/pkg/auth/middlewares"
	"net/http"
)

// InitUserRoutes initializes user-related routes
func InitUserRoutes(userController *controllers.UserController, router *mux.Router) {

	// Create a sub-router for user-related endpoints
	userRouter := router.PathPrefix("/users").Subrouter()

	// Define routes with authentication middleware

	// Get all users - Protected
	userRouter.Handle("", middleware.AuthMiddleware("/users", "GET", http.HandlerFunc(userController.GetAllUsersHandler))).Methods("GET")

	// Get user by ID - Protected
	userRouter.Handle("/{id}", middleware.AuthMiddleware("/users/{id}", "GET", http.HandlerFunc(userController.GetUserHandler))).Methods("GET")

	// Create a new user - Protected
	userRouter.Handle("", middleware.AuthMiddleware("/users", "POST", http.HandlerFunc(userController.CreateUserHandler))).Methods("POST")

	// Update user by ID - Protected
	userRouter.Handle("/{id}", middleware.AuthMiddleware("/users/{id}", "PUT", http.HandlerFunc(userController.UpdateUserHandler))).Methods("PUT")

	// Delete user by ID - Protected
	userRouter.Handle("/{id}", middleware.AuthMiddleware("/users/{id}", "DELETE", http.HandlerFunc(userController.DeleteUserHandler))).Methods("DELETE")
}
