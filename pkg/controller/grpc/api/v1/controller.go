package v1

import (
	"fmt"
	cfg "github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"github.com/krack8/lighthouse/pkg/controller/auth/services"
	"github.com/krack8/lighthouse/pkg/controller/core"
	"log"
)

type ControllerServer struct {
	pb.UnimplementedControllerServer
}

// TaskStream is a bidirectional stream method. The agent connects and sends messages here.
func (s *ControllerServer) TaskStream(stream pb.Controller_TaskStreamServer) error {
	// We’ll store the agent once we receive a AgentIdentification message.
	var currentAgent *core.AgentConnection
	defer func() {
		if currentAgent != nil {
			core.GetAgentManager().RemoveAgent(currentAgent)
		}
	}()

	// Listen for incoming messages from the agent.
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
			// Create the agent connection instance.
			currentAgent = &core.AgentConnection{
				Stream:      stream,
				GroupName:   groupName,
				ResultChMap: make(map[string]chan *pb.TaskResult),
			}

			// Add to the server’s group map.
			core.GetAgentManager().AddAgent(currentAgent)

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
			if currentAgent != nil {
				core.GetAgentManager().Lock()
				ch, ok := currentAgent.ResultChMap[taskRes.TaskId]
				core.GetAgentManager().Unlock()
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
