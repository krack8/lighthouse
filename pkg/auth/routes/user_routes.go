package routes

import (
	"github.com/gorilla/mux"
	"github.com/krack8/lighthouse/pkg/auth/controllers"
)

// InitUserRoutes initializes user-related routes
func InitUserRoutes(router *mux.Router) {

	// Create a sub-router for user-related endpoints
	userRouter := router.PathPrefix("/users").Subrouter()

	// Define routes
	userRouter.HandleFunc("/", controllers.GetAllUsersHandler).Methods("GET")       // Get all users
	userRouter.HandleFunc("/{id}", controllers.GetUserHandler).Methods("GET")       // Get user by ID
	userRouter.HandleFunc("/", controllers.CreateUserHandler).Methods("POST")       // Create a new user
	userRouter.HandleFunc("/{id}", controllers.UpdateUserHandler).Methods("PUT")    // Update user by ID
	userRouter.HandleFunc("/{id}", controllers.DeleteUserHandler).Methods("DELETE") // Delete user by ID
}
