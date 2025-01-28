package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cfg "github.com/krack8/lighthouse/pkg/config"
	_log "github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/server/router"
	"net/http"
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
	router.AddApiRoutes(httpRouter)
	//r.GET("/health", api.Home().Health)
	//r.GET("/", api.Home().Index)
	err := r.Run(":" + cfg.ServerPort) // listen and serve on 0.0.0.0:8080
	if err != nil {
		_log.Logger.Errorw("Failed to start server", "err", err.Error())
	}
}
