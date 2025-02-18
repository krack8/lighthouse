package worker

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/common/pb"
	cfg "github.com/krack8/lighthouse/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
	"time"
)

type TaskToAgentInterface interface {
	SendToWorker(c context.Context, taskName string, input []byte, groupName string) (*pb.TaskResult, error)
}

type taskToAgent struct{}

var tta taskToAgent

func TaskToAgent() *taskToAgent {
	return &tta
}

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

			// Validate group name
			if groupName == "" {
				log.Printf("Worker attempted to connect with empty group name - rejecting connection")
				err := stream.Send(&pb.TaskStreamResponse{
					Payload: &pb.TaskStreamResponse_Ack{
						Ack: &pb.Ack{
							Message: "group_name_required",
						},
					},
				})
				if err != nil {
					return err
				}
				return fmt.Errorf("worker connection rejected: empty group name")
			}

			//Verify the auth token if AUTH is ENABLED
			if cfg.IsAuth() {
				if authToken != "" {
					if services.IsAgentAuthTokenValid(authToken) == false {
						err := stream.Send(&pb.TaskStreamResponse{
							Payload: &pb.TaskStreamResponse_Ack{
								Ack: &pb.Ack{
									Message: "Invalid Agent Token",
								},
							},
						})
						if err != nil {
							return err
						}
						return nil
					}
				} else {
					log.Printf("Worker auth token is required")
					err := stream.Send(&pb.TaskStreamResponse{
						Payload: &pb.TaskStreamResponse_Ack{
							Ack: &pb.Ack{
								Message: "Worker auth token is required",
							},
						},
					})
					if err != nil {
						return err
					}
					return nil
				}
			}
			log.Printf("New worker identified. group=%s", groupName)

			// Create the worker connection instance.
			currentWorker = &workerConnection{
				stream:      stream,
				groupName:   groupName,
				resultChMap: make(map[string]chan *pb.TaskResult),
			}

			// Add to the server’s group map.
			s.addWorker(currentWorker)

			// Send back a simple Ack
			_ = stream.Send(&pb.TaskStreamResponse{
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

func (s *serverImpl) Process(groupName string, payload string, taskName string, input []byte) (<-chan *pb.TaskResult, error) {
	w := s.pickWorker(groupName)
	if w == nil {
		return nil, errors.New("worker unreachable")
	}
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
				Name:    taskName,
				Payload: payload,
				Input:   string(input),
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

func (tta *taskToAgent) SendToWorker(c context.Context, taskName string, input []byte, groupName string) (*pb.TaskResult, error) {
	payload := taskName
	// Get the server instance
	server := GetServerInstance()
	resultCh, err := server.Process(groupName, payload, taskName, input)
	if err != nil {
		return nil, err
	}
	select {
	case res := <-resultCh:
		// Send response to the user
		if !res.Success {
			return nil, errors.New(res.Output)
		}
		return res, nil
	case <-time.After(10 * time.Second):
		return nil, errors.New("timed out waiting for worker result")
	}
}

// disconnectWorker handles immediate worker disconnection
func (s *serverImpl) disconnectWorker(w *workerConnection) {
	if w == nil || w.stream == nil {
		log.Printf("Invalid worker connection")
		return
	}

	// Lock before any operations
	s.mu.Lock()

	// Verify worker exists in the group
	workers := s.groups[w.groupName]
	workerFound := false
	for _, conn := range workers {
		if conn == w {
			workerFound = true
			break
		}
	}

	if !workerFound {
		log.Printf("Worker not found in group %s", w.groupName)
		s.mu.Unlock()
		return
	}

	log.Printf("Sending disconnect message - Group: %s, Total workers in group: %d",
		w.groupName, len(workers))

	// Send disconnect message immediately
	err := w.stream.Send(&pb.TaskStreamResponse{
		Payload: &pb.TaskStreamResponse_Ack{
			Ack: &pb.Ack{
				Message: "disconnect_requested",
			},
		},
	})

	if err != nil {
		log.Printf("Failed to send disconnect message to group %s: %v", w.groupName, err)
	} else {
		log.Printf("Successfully sent disconnect message to group %s", w.groupName)
	}

	// Cleanup channels
	for taskID, ch := range w.resultChMap {
		close(ch)
		delete(w.resultChMap, taskID)
	}

	// Remove worker from group
	var newList []*workerConnection
	for _, conn := range workers {
		if conn != w {
			newList = append(newList, conn)
		}
	}
	s.groups[w.groupName] = newList

	s.mu.Unlock()
}

// RemoveWorkerByGroupName removes all workers in a group
func (s *serverImpl) RemoveWorkerByGroupName(groupName string) bool {
	s.mu.Lock()
	workers, exists := s.groups[groupName]
	if !exists || len(workers) == 0 {
		log.Printf("No workers found in group: %s", groupName)
		s.mu.Unlock()
		return false
	}

	workerCount := len(workers)
	log.Printf("Found %d workers in group %s", workerCount, groupName)
	s.mu.Unlock()

	for i, worker := range workers {
		log.Printf("Disconnecting worker %d/%d in group %s", i+1, workerCount, groupName)
		s.disconnectWorker(worker)
	}

	return true
}

// Use a proper singleton pattern:
var (
	srv     *serverImpl
	srvOnce sync.Once
	srvMu   sync.RWMutex
)

// GetServerInstance returns the singleton server instance
func GetServerInstance() *serverImpl {
	srvOnce.Do(func() {
		srvMu.Lock()
		defer srvMu.Unlock()
		if srv == nil {
			srv = &serverImpl{
				groups: make(map[string][]*workerConnection),
			}
		}
	})
	return srv
}

func StartGrpcServer() {
	// Start gRPC server
	go func() {
		grpcServer := grpc.NewServer()
		server := GetServerInstance()
		pb.RegisterControllerServer(grpcServer, server)
		reflection.Register(grpcServer)

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		log.Println("Starting Controller gRPC server on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()
}
