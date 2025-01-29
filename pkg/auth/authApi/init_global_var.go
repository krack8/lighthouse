package authApi

import (
	"github.com/krack8/lighthouse/pkg/auth/controllers"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"sync"
)

var (
	// Global service instances
	UserService *services.UserService
	RbacService *services.RbacService

	// UserController Global controller instances
	UserController *controllers.UserController
	RbacController *controllers.RbacController

	once sync.Once
)

// Init initializes all services and controllers
func Init() {
	once.Do(func() {
		// Initialize services
		UserService = &services.UserService{}
		RbacService = &services.RbacService{}

		// Initialize controllers with services
		UserController = controllers.NewUserController(UserService)
		RbacController = controllers.NewRbacController(RbacService)
	})
}

// GetUserService returns the global UserService instance
func GetUserService() *services.UserService {
	return UserService
}

// GetRbacService returns the global RbacService instance
func GetRbacService() *services.RbacService {
	return RbacService
}

// GetUserController returns the global UserController instance
func GetUserController() *controllers.UserController {
	return UserController
}

// GetRbacController returns the global RbacController instance
func GetRbacController() *controllers.RbacController {
	return RbacController
}
