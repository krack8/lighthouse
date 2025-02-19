package main

import (
	cfg "github.com/krack8/lighthouse/pkg/common/config"
	_log "github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/controller/auth/authApi"
	"github.com/krack8/lighthouse/pkg/controller/auth/config"
	"github.com/krack8/lighthouse/pkg/controller/core"
	srvr "github.com/krack8/lighthouse/pkg/controller/server"
	"log"
	"net/http"
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

	// Start HTTP server
	srvr.StartHttServer()
	_log.Logger.Infow("HTTP server listening on :"+cfg.ServerPort, "[Server-Port]", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, nil))
}
