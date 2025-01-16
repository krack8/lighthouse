package routes

import (
	"github.com/gorilla/mux"
	"github.com/krack8/lighthouse/pkg/auth/controllers"
)

// RegisterRoutes registers all application routes.
func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/users", controllers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUserHandler).Methods("DELETE")

	return router
}
