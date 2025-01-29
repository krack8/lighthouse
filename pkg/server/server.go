package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/controllers"
	middleware "github.com/krack8/lighthouse/pkg/auth/middlewares"
	cfg "github.com/krack8/lighthouse/pkg/config"
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
	//r.Use(cors.New(corsConfig))
	r.Use(cors.New(corsConfig), gin.LoggerWithWriter(gin.DefaultWriter, "/health", "/swagger/*any"))

	// Setting API Base Path for HTTP APIs
	httpRouter := r.Group("api/v1")
	fmt.Println(cfg.IsNoAuth())
	if !cfg.IsNoAuth() {
		// Apply the AuthMiddleware to the / routes
		httpRouter = r.Group("api/v1", middleware.AuthMiddleware())
	}
	// Get the application port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	router.AddApiRoutes(httpRouter)
	r.GET("/health", Home().Health)
	r.GET("/", Home().Index)
	// Define the login route separately without middleware
	// Login route
	r.POST("/auth/login", controllers.LoginHandler)
	// Refresh token route
	r.POST("/auth/refresh-token", controllers.RefreshTokenHandler)

	err := r.Run(":" + port) // listen and serve
	if err != nil {
		_log.Logger.Errorw("Failed to start server", "err", err.Error())
	}
}

type HomeInf interface {
	Index(c *gin.Context)
	Health(c *gin.Context)
}

type home struct{}

func Home() HomeInf {
	return &home{}
}

func (h *home) Index(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/json; charset=utf-8", []byte("This is KloverCloud Lighthouse"))
	return
}

func (h *home) Health(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/json; charset=utf-8", []byte("I am live!"))
	return
}
