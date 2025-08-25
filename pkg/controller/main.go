package main

import (
	"github.com/gin-gonic/gin"
	cfg "github.com/krack8/lighthouse/pkg/common/config"
	_log "github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/controller/api/argocd"
	"github.com/krack8/lighthouse/pkg/controller/auth/authApi"
	"github.com/krack8/lighthouse/pkg/controller/auth/config"
	"github.com/krack8/lighthouse/pkg/controller/core"
	srvr "github.com/krack8/lighthouse/pkg/controller/server"
	"log"
	"net/http"
	"sync"
)

func main() {
	_log.InitializeLogger()

	cfg.InitEnvironmentVariables()
	// Init Agent Connection Manager
	core.InitAgentConnectionManager()

	go srvr.StartGrpcServer()

	// Connect to the database ..
	client, ctx, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
		return
	}
	// Add this if you want to reset collections on startup
	//if err := config.ResetCollections(client, ctx); err != nil {
	//	_log.Logger.Errorw("Failed to reset collections", "err", err)
	//	return
	//}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Error disconnecting from DB: %v", err)
		}
	}()

	// Initialize the default RBAC if needed
	config.InitRBAC()

	// Initialize the default user if needed
	config.InitializeDefaultUser()

	// Initialize the default cluster if needed
	config.InitializeClusters()

	// Initialize auth controllers with services
	authApi.Init()

	agentManager := argocd.NewAgentManager()

	// Use WaitGroup to run both servers
	var wg sync.WaitGroup
	wg.Add(2)

	// Start ArgoCD HTTP server on different port (8081)
	go func() {
		defer wg.Done()
		startArgoCDServer(agentManager)
	}()

	// Start existing HTTP server on original port
	go func() {
		defer wg.Done()
		srvr.StartHttServer()
		_log.Logger.Infow("Main HTTP server listening on :"+cfg.ServerPort, "[Server-Port]", cfg.ServerPort)
		log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, nil))
	}()

	wg.Wait()
}

func startArgoCDServer(agentManager *argocd.AgentManager) {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		agents := agentManager.ListAgents()
		c.JSON(200, gin.H{
			"status": "healthy",
			"agents": len(agents),
		})
	})

	// ArgoCD API routes
	v1 := router.Group("/api/v1")
	{
		argoCDHandler := argocd.NewHandler(agentManager)
		argocd.RegisterRoutes(v1, argoCDHandler)
	}

	// Run on different port for ArgoCD
	argoCDPort := "8081"
	_log.Logger.Infow("ArgoCD API server listening on :"+argoCDPort, "[ArgoCD-Port]", argoCDPort)
	if err := router.Run(":" + argoCDPort); err != nil {
		log.Fatal("Failed to start ArgoCD server:", err)
	}
}
