package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_log "github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/server/router"
	"net/http"
	"os"
)

func Start() {
	//gin.DisableConsoleColor()
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*", "http://localhost:4200"}
	corsConfig.AllowMethods = []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Origin", "*"}
	corsConfig.AddAllowMethods("OPTIONS")
	r.Use(cors.New(corsConfig), gin.LoggerWithWriter(gin.DefaultWriter, "/health", "/swagger/*any"))
	// Setting API Base Path for HTTP APIs
	httpRouter := r.Group("/")
	//r.GET("/ws/v1/pod/logs/:name", v1.PodController().GetPodStreamLogs)
	// Setting up all Http Routes

	/*// Initialize services and controllers
	userService := &services.UserService{} // Ensure it is properly initialized
	userController := &controllers.UserController{UserService: userService}

	rbacService := &services.RbacService{} // Ensure it is properly initialized
	rbacController := &controllers.RbacController{RbacService: rbacService}

	// Initialize routes from various route files
	routes.InitPermissionRoutes(rbacController, router) // permission-related routes
	routes.InitRoleRoutes(rbacController, router)       // role-related routes

	routes.InitAuthRoutes(router)                 // Auth-related routes
	routes.InitUserRoutes(userController, router) // user-related routes*/
	// Get the application port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	router.AddApiRoutes(httpRouter)
	//r.GET("/health", api.Home().Health)
	//r.GET("/", api.Home().Index)
	err := r.Run(":" + port) // listen and serve
	if err != nil {
		_log.Logger.Errorw("Failed to start server", "err", err.Error())
	}
}
