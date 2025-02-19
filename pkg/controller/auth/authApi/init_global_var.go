package authApi

import (
	controllers2 "github.com/krack8/lighthouse/pkg/controller/auth/controllers"
	services2 "github.com/krack8/lighthouse/pkg/controller/auth/services"
	"sync"
)

var (
	// Global service instances
	UserService    *services2.UserService
	RbacService    *services2.RbacService
	ClusterService *services2.ClusterService

	// UserController Global controller instances
	UserController    *controllers2.UserController
	RbacController    *controllers2.RbacController
	ClusterController *controllers2.ClusterController

	once sync.Once
)

// Init initializes all services and controllers
func Init() {
	once.Do(func() {
		// Initialize services
		UserService = &services2.UserService{}
		RbacService = &services2.RbacService{}
		ClusterService = &services2.ClusterService{}

		// Initialize controllers with services
		UserController = controllers2.NewUserController(UserService)
		RbacController = controllers2.NewRbacController(RbacService)
		ClusterController = controllers2.NewClusterController(ClusterService)
	})
}

// GetUserService returns the global UserService instance
func GetUserService() *services2.UserService {
	return UserService
}

// GetRbacService returns the global RbacService instance
func GetRbacService() *services2.RbacService {
	return RbacService
}

// GetUserController returns the global UserController instance
func GetUserController() *controllers2.UserController {
	return UserController
}

// GetRbacController returns the global RbacController instance
func GetRbacController() *controllers2.RbacController {
	return RbacController
}
