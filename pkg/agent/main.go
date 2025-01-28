package main

import (
	"context"
	"encoding/json"
	"github.com/krack8/lighthouse/pkg/config"
	_log "github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"

	"github.com/krack8/lighthouse/pkg/common/pb" // Import the generated proto package
)

var taskMutex sync.Mutex

func main() {
	_log.InitializeLogger()
	config.InitiateKubeClientSet()
	// For demonstration, we'll just run a single worker that belongs to "GroupA".
	groupName := "GroupA"
	authToken := "my-secret"
	tasks.InitTaskRegistry()

	// Dial the controller's gRPC server.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial controller: %v", err)
	}
	defer conn.Close()

	client := pb.NewControllerClient(conn)
	// Open the bi-directional stream.
	stream, err := client.TaskStream(context.Background())
	if err != nil {
		log.Fatalf("Failed to create TaskStream: %v", err)
	}

	// 1) Send WorkerIdentification
	err = stream.Send(&pb.TaskStreamRequest{
		Payload: &pb.TaskStreamRequest_WorkerInfo{
			WorkerInfo: &pb.WorkerIdentification{
				GroupName: groupName,
				AuthToken: authToken,
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to send worker info: %v", err)
	}

	// We'll handle incoming messages in a separate goroutine.
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				log.Printf("Stream Recv error (worker): %v", err)
				return
			}

			switch payload := in.Payload.(type) {

			case *pb.TaskStreamResponse_NewTask:
				task := payload.NewTask
				log.Printf("Worker received a new task: ID=%s, payload=%s",
					task.Id, task.Payload)
				go func(taskID, taskPayload string, task *pb.Task) {
					taskMutex.Lock()
					defer taskMutex.Unlock()
					TaskResult := &pb.TaskResult{}
					res, err := tasks.TaskSelector(task)
					if err != nil {
						TaskResult.Success = false
						TaskResult.Output = err.Error()
					} else {
						output, err := json.Marshal(res)
						if err != nil {
							TaskResult.Success = false
							TaskResult.Output = err.Error()
						}
						TaskResult.Success = true
						TaskResult.Output = string(output)
					}
					TaskResult.TaskId = taskID
					resultMsg := &pb.TaskStreamRequest{
						Payload: &pb.TaskStreamRequest_TaskResult{
							TaskResult: TaskResult,
						},
					}

					// Send the result back to the controller.
					if err := stream.Send(resultMsg); err != nil {
						log.Printf("Failed to send task result: %v", err)
					}
				}(task.Id, task.Payload, task)

			case *pb.TaskStreamResponse_Ack:
				log.Printf("Worker received an ACK from server: %s", payload.Ack.Message)

			default:
				log.Printf("Unknown payload from server.")
			}
		}
	}()

	// Keep the worker alive.
	select {}
}
