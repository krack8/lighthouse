package routes

import (
	"github.com/gorilla/mux"
	"github.com/krack8/lighthouse/pkg/auth/controllers"
	middleware "github.com/krack8/lighthouse/pkg/auth/middlewares"
	"net/http"
)

// InitPermissionRoutes initializes permission-related routes
func InitPermissionRoutes(rbacController *controllers.RbacController, router *mux.Router) {

	// Create a sub-router for permission-related endpoints
	permissionRouter := router.PathPrefix("/permissions").Subrouter()

	// Create a new Permission - Protected
	permissionRouter.Handle("", middleware.AuthMiddleware("/permissions", "POST", http.HandlerFunc(rbacController.CreatePermissionHandler))).Methods("POST") // Create new permission
}

// InitRoleRoutes initializes role-related routes.
func InitRoleRoutes(rbacController *controllers.RbacController, router *mux.Router) {

	// Create a sub-router for role-related endpoints
	roleRouter := router.PathPrefix("/roles").Subrouter()

	// Create a new Role - Protected
	roleRouter.Handle("", middleware.AuthMiddleware("/roles", "POST", http.HandlerFunc(rbacController.CreatePermissionHandler))).Methods("POST") // Create new role

	// Assign multiple roles to a user - Protected
	roleRouter.Handle("/assign/multiple", middleware.AuthMiddleware("/roles/assign/multiple", "POST", http.HandlerFunc(rbacController.AssignRolesHandler))).Methods("POST") // Create new role  // Assign role to user
}
