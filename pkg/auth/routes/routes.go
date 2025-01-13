package routes

import (
	"github.com/krack8/lighthouse/pkg/auth/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes defines the routes for user operations.
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.CreateUser).Methods(http.MethodPost)
}
