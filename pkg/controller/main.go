package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/krack8/lighthouse/pkg/auth/config"
	"github.com/krack8/lighthouse/pkg/auth/routes"
	"github.com/krack8/lighthouse/pkg/common/pb" // Import the generated proto package
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

// workerConnection represents one worker's active streaming connection.
type workerConnection struct {
	stream      pb.Controller_TaskStreamServer
	groupName   string
	resultChMap map[string]chan *pb.TaskResult // map of taskID -> channel that receives result
}

// serverImpl implements the pb.ControllerServer interface.
type serverImpl struct {
	pb.UnimplementedControllerServer

	mu     sync.Mutex
	groups map[string][]*workerConnection // groupName -> slice of workers
}

// TaskStream is a bidirectional stream method. The worker connects and sends messages here.
func (s *serverImpl) TaskStream(stream pb.Controller_TaskStreamServer) error {
	// We’ll store the worker once we receive a WorkerIdentification message.
	var currentWorker *workerConnection
	defer func() {
		if currentWorker != nil {
			s.removeWorker(currentWorker)
		}
	}()

	// Listen for incoming messages from the worker.
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Printf("Stream recv error: %v", err)
			return err
		}

		switch payload := req.Payload.(type) {

		case *pb.TaskStreamRequest_WorkerInfo:
			// This is the first message from the worker identifying itself.
			groupName := payload.WorkerInfo.GroupName
			authToken := payload.WorkerInfo.AuthToken
			// Here you could verify the auth token.
			log.Printf("New worker identified. group=%s, token=%s", groupName, authToken)

			// Create the worker connection instance.
			currentWorker = &workerConnection{
				stream:      stream,
				groupName:   groupName,
				resultChMap: make(map[string]chan *pb.TaskResult),
			}

			// Add to the server’s group map.
			s.addWorker(currentWorker)

			// Send back a simple Ack
			stream.Send(&pb.TaskStreamResponse{
				Payload: &pb.TaskStreamResponse_Ack{
					Ack: &pb.Ack{Message: "Registered successfully"},
				},
			})

		case *pb.TaskStreamRequest_TaskResult:
			// The worker has completed a task and is sending the result.
			taskRes := payload.TaskResult
			log.Printf("Received task result from worker: task_id=%s, success=%v",
				taskRes.TaskId, taskRes.Success)

			// Notify whoever is waiting for this task result (our HTTP handler).
			if currentWorker != nil {
				s.mu.Lock()
				ch, ok := currentWorker.resultChMap[taskRes.TaskId]
				s.mu.Unlock()
				if ok {
					ch <- taskRes
				} else {
					log.Printf("No channel waiting for task_id=%s", taskRes.TaskId)
				}
			}

		default:
			log.Printf("Unknown message type from worker stream")
		}
	}
}

func (s *serverImpl) addWorker(w *workerConnection) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.groups[w.groupName] = append(s.groups[w.groupName], w)
	log.Printf("Worker added to group %q. Total workers in group: %d",
		w.groupName, len(s.groups[w.groupName]))
}

func (s *serverImpl) removeWorker(w *workerConnection) {
	s.mu.Lock()
	defer s.mu.Unlock()
	workers := s.groups[w.groupName]
	var newList []*workerConnection
	for _, conn := range workers {
		if conn != w {
			newList = append(newList, conn)
		}
	}
	s.groups[w.groupName] = newList
	log.Printf("Worker removed from group %q. Remaining workers: %d",
		w.groupName, len(s.groups[w.groupName]))
}

// sendTaskToWorker sends a task down a particular worker’s stream.
// Returns a channel on which the result will be delivered.
func (s *serverImpl) sendTaskToWorker(w *workerConnection, payload string) (<-chan *pb.TaskResult, error) {
	// Generate a task ID.
	taskID := uuid.NewString()

	// Prepare a channel to receive the worker’s response.
	resultCh := make(chan *pb.TaskResult, 1)

	s.mu.Lock()
	w.resultChMap[taskID] = resultCh
	s.mu.Unlock()

	// Actually send the task to the worker.
	err := w.stream.Send(&pb.TaskStreamResponse{
		Payload: &pb.TaskStreamResponse_NewTask{
			NewTask: &pb.Task{
				Id:      taskID,
				Name:    "TASK_NAME",
				Payload: payload,
			},
		},
	})
	if err != nil {
		s.mu.Lock()
		delete(w.resultChMap, taskID)
		s.mu.Unlock()
		return nil, err
	}

	return resultCh, nil
}

// pickWorker returns any worker from the specified group (round-robin or random).
// For simplicity, let's just pick the first.
func (s *serverImpl) pickWorker(groupName string) *workerConnection {
	s.mu.Lock()
	defer s.mu.Unlock()
	workers := s.groups[groupName]
	if len(workers) == 0 {
		return nil
	}
	// naive pick: the first worker
	return workers[0]
}

// HTTP handler: /execute?group=GroupA&payload=SomeData
func (s *serverImpl) httpExecuteHandler(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	payload := r.URL.Query().Get("payload")
	if group == "" || payload == "" {
		http.Error(w, "Missing group or payload param", http.StatusBadRequest)
		return
	}

	worker := s.pickWorker(group)
	if worker == nil {
		http.Error(w, "No worker in group "+group, http.StatusServiceUnavailable)
		return
	}

	resultCh, err := s.sendTaskToWorker(worker, payload)
	if err != nil {
		http.Error(w, "Failed to send task to worker: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Wait for the worker to respond with a result or time out
	select {
	case res := <-resultCh:
		// Send response to the user
		fmt.Fprintf(w, "Task ID: %s\nSuccess: %v\nOutput: %s\n",
			res.TaskId, res.Success, res.Output)
	case <-time.After(10 * time.Second):
		http.Error(w, "Timed out waiting for worker result", http.StatusGatewayTimeout)
	}
}

func main() {
	srv := &serverImpl{
		groups: make(map[string][]*workerConnection),
	}

	// Start gRPC server
	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterControllerServer(grpcServer, srv)

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		log.Println("Starting Controller gRPC server on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

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

	// Initialize the default user if needed
	config.InitializeDefaultUser()

	// Initialize router
	router := mux.NewRouter()

	// Initialize routes from various route files
	routes.InitAuthRoutes(router) // Auth-related routes
	routes.InitUserRoutes(router) // user-related routes                               // User-related routes

	// Get the application port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	// Start HTTP server
	http.HandleFunc("/execute", srv.httpExecuteHandler)
	log.Println("HTTP server listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
