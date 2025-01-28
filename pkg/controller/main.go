package main

import (
	"github.com/joho/godotenv"
	"github.com/krack8/lighthouse/pkg/auth/config"
	"github.com/krack8/lighthouse/pkg/controller/worker"
	_log "github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/server"
	"log"
	"net/http"
)

func main() {
	_log.InitializeLogger()
	worker.StartGrpcServer()

	// Load environment variables from .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect to the database
	client, ctx, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
		return
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Error disconnecting from DB: %v", err)
		}
	}()

	// Initialize the default RBAC if needed
	config.InitRBAC()

	// Initialize the default user if needed
	config.InitializeDefaultUser()

	// Start HTTP server
	server.Start()
	log.Println("HTTP server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
