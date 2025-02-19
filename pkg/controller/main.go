package main

import (
	cfg "github.com/krack8/lighthouse/pkg/config"
	"github.com/krack8/lighthouse/pkg/controller/auth/authApi"
	"github.com/krack8/lighthouse/pkg/controller/auth/config"
	srvr "github.com/krack8/lighthouse/pkg/controller/server"
	_log "github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/server"
	"log"
	"net/http"
)

func main() {
	_log.InitializeLogger()
	srvr.StartGrpcServer()
	cfg.InitEnvironmentVariables()

	// Connect to the database ..
	client, ctx, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
		return
	}
	/*// Add this if you want to reset collections on startup
	if err := config.ResetCollections(client, ctx); err != nil {
		log.Printf("Failed to reset collections: %v", err)
		return
	}*/
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
	server.Start()
	log.Println("HTTP server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
