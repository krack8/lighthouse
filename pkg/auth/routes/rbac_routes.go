package routes

import (
	"github.com/gorilla/mux"
	"github.com/krack8/lighthouse/pkg/auth/controllers"
	middleware "github.com/krack8/lighthouse/pkg/auth/middlewares"
)

// InitPermissionRoutes initializes permission-related routes
func InitPermissionRoutes(router *mux.Router) {

	// Create a sub-router for permission-related endpoints
	permissionRouter := router.PathPrefix("/permissions").Subrouter()

	// Create a new Permission - Protected
	permissionRouter.HandleFunc("/", controllers.CreatePermissionHandler).Methods("POST").Handler(middleware.AuthMiddleware("/permissions", "POST")) // Create new permission
}

// InitRoleRoutes initializes role-related routes.
func InitRoleRoutes(router *mux.Router) {

	// Create a sub-router for role-related endpoints
	roleRouter := router.PathPrefix("/roles").Subrouter()

	// Create a new Role - Protected
	roleRouter.HandleFunc("/", controllers.CreateRoleHandler).Methods("POST").Handler(middleware.AuthMiddleware("/roles", "POST")) // Create new role

	// Assign multiple roles to a user - Protected
	roleRouter.HandleFunc("/assign/multiple", controllers.AssignRolesHandler).Methods("POST").Handler(middleware.AuthMiddleware("/roles/assign/multiple", "POST")) // Assign role to user
}
