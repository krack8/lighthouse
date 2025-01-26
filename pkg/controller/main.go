package main

import (
	"github.com/krack8/lighthouse/pkg/controller/worker"
	_log "github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/server"
	"log"
	"net/http"
)

func main() {
	_log.InitializeLogger()
	worker.StartGrpcServer()

	// Start HTTP server
	server.Start()
	log.Println("HTTP server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
