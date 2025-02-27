package v1

import (
	"fmt"
	cfg "github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"github.com/krack8/lighthouse/pkg/controller/auth/services"
	"github.com/krack8/lighthouse/pkg/controller/core"
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
			log.Logger.Warnw("Stream recv error: %v", "err", err)
			return err
		}

		switch payload := req.Payload.(type) {

		case *pb.TaskStreamRequest_WorkerInfo:
			// This is the first message from the worker identifying itself.
			groupName := payload.WorkerInfo.GroupName
			authToken := payload.WorkerInfo.AuthToken

			// Validate group name
			if groupName == "" {
				log.Logger.Warnw("Agent attempted to connect with empty group name - rejecting connection", "agent-identity", "empty group-name: "+groupName)
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
					isTokenValid, isGroupNameValid := services.IsAgentAuthTokenAndGroupValid(authToken, groupName)
					if isTokenValid == false {
						err := stream.Send(&pb.TaskStreamResponse{
							Payload: &pb.TaskStreamResponse_Ack{
								Ack: &pb.Ack{
									Message: "Unauthorized !! Invalid Agent Token",
								},
							},
						})
						if err != nil {
							return err
						}
						return nil
					}
					if isTokenValid == true && isGroupNameValid == false {
						err := stream.Send(&pb.TaskStreamResponse{
							Payload: &pb.TaskStreamResponse_Ack{
								Ack: &pb.Ack{
									Message: "Invalid cluster group. The group is not registered with this cluster",
								},
							},
						})
						if err != nil {
							return err
						}
						return nil
					}
				} else {
					log.Logger.Warnw("Agent auth token is required", "agent-identity", "token: "+authToken)
					err := stream.Send(&pb.TaskStreamResponse{
						Payload: &pb.TaskStreamResponse_Ack{
							Ack: &pb.Ack{
								Message: "Agent auth token is required",
							},
						},
					})
					if err != nil {
						return err
					}
					return nil
				}
			}
			log.Logger.Infow("New agent identified. Group: "+groupName, "agent-identity", "Group: "+groupName)

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
			log.Logger.Infow(fmt.Sprintf("Received task result from agent: task_id=%s, success=%v",
				taskRes.TaskId, taskRes.Success), "task-result", taskRes.Success)

			// Notify whoever is waiting for this task result (our HTTP handler).
			if currentAgent != nil {
				currentAgent.Lock()
				ch, ok := currentAgent.ResultChMap[taskRes.TaskId]
				currentAgent.Unlock()
				if ok {
					ch <- taskRes
				} else {
					log.Logger.Infow(fmt.Sprintf("No channel waiting for task_id=%s", taskRes.TaskId), "channel", "not waiting")
				}
			}
		case *pb.TaskStreamRequest_LogsResult:
			// The worker is streaming logs
			taskRes := payload.LogsResult
			log.Logger.Infow(fmt.Sprintf("Received task result from agent: task_id=%s",
				taskRes.TaskId), "task-result", taskRes.TaskId)

			// Notify whoever is waiting for this task result (our HTTP handler).
			if currentAgent != nil {
				currentAgent.Lock()
				ch, ok := currentAgent.ResultChMap[taskRes.TaskId]
				currentAgent.Unlock()
				if ok {
					// If it's a log streaming task, send the logs incrementally
					var logs []string
					log.Logger.Infow("printing logs: ", "print", string(taskRes.Output))
					if err != nil {
						log.Logger.Errorw("Failed to unmarshal logs", "err", err)
						ch <- &pb.TaskResult{
							Success: false,
							Output:  "Failed to unmarshal logs",
						}
						return nil
					}

					for idx, logLine := range logs {
						fmt.Println(fmt.Sprintf("%d --> %s", idx, logLine))
						ch <- &pb.TaskResult{
							Success: true,
							Output:  logLine, // Send each log line individually
						}
					}

					// Send a final message to indicate the end of the log stream
					ch <- &pb.TaskResult{
						Success: true,
						Output:  "[DONE]", // Or a specific message indicating the end of the stream
					}
				} else {
					log.Logger.Infow(fmt.Sprintf("No channel waiting for task_id=%s", taskRes.TaskId), "channel", "not waiting")
				}
			}

		default:
			log.Logger.Warnw("Unknown message type from agent stream", "message-type", "unknown")
		}
	}
}
