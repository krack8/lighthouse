package main

import (
	"context"
	"fmt"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/tasks"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/krack8/lighthouse/pkg/common/pb" // Import the generated proto package
)

func main() {
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

				// Here is where you do the actual business logic.
				// We'll just pretend to do some work and return a result.
				go func(taskID, taskPayload string, task *pb.Task) {
					// Simulate some processing time.
					var res interface{}
					newTask := tasks.TaskRegistry[task.Name]
					switch v := newTask.TaskInput.(type) {
					case k8s.GetNamespaceInputParams:
						fmt.Println("Bhua:", v)
					case k8s.GetNamespaceListInputParams:
						fmt.Println("Get Namespace List Input Params:", v)
						execute, _ := newTask.TaskFunc.(func(context.Context, k8s.GetNamespaceListInputParams) (interface{}, error))
						res, err = execute(context.Background(), k8s.GetNamespaceListInputParams{})
						if err != nil {
							log.Printf(err.Error())
						}
					}
					time.Sleep(2 * time.Second)

					// Build the result.
					resultMsg := &pb.TaskStreamRequest{
						Payload: &pb.TaskStreamRequest_TaskResult{
							TaskResult: &pb.TaskResult{
								TaskId:  taskID,
								Success: true,
								Output:  fmt.Sprintf("Processed payload: %s", res),
							},
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
