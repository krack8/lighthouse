package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"net/http"
	"sync"
	"time"
)

// AgentConnection represents one agent's active Streaming connection.
type AgentConnection struct {
	Stream pb.Controller_TaskStreamServer
	//UniqueId    string
	GroupName         string
	ResultChMap       map[string]chan *pb.TaskResult
	ResultStreamChMap map[string]chan *pb.LogsResult // map of taskID -> channel that receives result
	mu                sync.Mutex
}

type AgentManager struct {
	mu             sync.RWMutex
	connectionList map[string][]*AgentConnection // GroupName -> slice of agents
}

var agentManager AgentManager

func InitAgentConnectionManager() {
	agentManager = AgentManager{
		connectionList: make(map[string][]*AgentConnection),
	}
}

func GetAgentManager() *AgentManager {
	return &agentManager
}

func (ac *AgentConnection) Lock() {
	ac.mu.Lock()
}

func (ac *AgentConnection) Unlock() {
	ac.mu.Unlock()
}

func (s *AgentManager) Lock() {
	s.mu.Lock()
}

func (s *AgentManager) Unlock() {
	s.mu.Unlock()
}

func (s *AgentManager) AddAgent(w *AgentConnection) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.connectionList[w.GroupName] = append(s.connectionList[w.GroupName], w)
	log.Logger.Infow(fmt.Sprintf("Agent added to connection list in group %q. Total agents in group: %d",
		w.GroupName, len(s.connectionList[w.GroupName])), "agent-connected", "group: "+w.GroupName)
}

func (s *AgentManager) RemoveAgent(w *AgentConnection) {
	s.mu.Lock()
	defer s.mu.Unlock()
	agents := s.connectionList[w.GroupName]
	var newList []*AgentConnection
	for _, conn := range agents {
		if conn != w {
			newList = append(newList, conn)
		}
	}
	s.connectionList[w.GroupName] = newList
	log.Logger.Warnw(fmt.Sprintf("Agent removed from connection list in group %q. Remaining agents: %d",
		w.GroupName, len(s.connectionList[w.GroupName])), "agent-removed", "group: "+w.GroupName)
}

func (s *AgentManager) RemoveAgentByGroupName(groupName string) bool {
	s.mu.Lock()
	workers, exists := s.connectionList[groupName]
	if !exists || len(workers) == 0 {
		log.Logger.Warnw(fmt.Sprintf("No agent found in group: %s", groupName), "agent-remove", "group: "+groupName)
		s.mu.Unlock()
		return false
	}

	workerCount := len(workers)
	log.Logger.Infow(fmt.Sprintf("Found %d agents in group %s", workerCount, groupName), "agent-remove", "group: "+groupName)
	s.mu.Unlock()

	for i, worker := range workers {
		log.Logger.Warnw(fmt.Sprintf("Disconnecting agent %d/%d in group %s", i+1, workerCount, groupName), "agent-remove", "group: "+groupName)
		s.disconnectWorker(worker)
	}

	return true
}

// disconnectWorker handles immediate worker disconnection
func (s *AgentManager) disconnectWorker(w *AgentConnection) {
	if w == nil || w.Stream == nil {
		log.Logger.Warnw("Invalid agent connection", "agent-disconnect", "group")
		return
	}

	// Lock before any operations
	s.mu.Lock()
	defer s.mu.Unlock()

	// Verify worker exists in the group
	workers := s.connectionList[w.GroupName]
	workerFound := false
	for _, conn := range workers {
		if conn == w {
			workerFound = true
			break
		}
	}

	if !workerFound {
		log.Logger.Warnw("Agent not found in group: "+w.GroupName, "agent-disconnect", "not found")
		return
	}

	log.Logger.Infow(fmt.Sprintf("Sending disconnect message - Group: %s, Total agents in group: %d",
		w.GroupName, len(workers)), "agent-disconnect", "sending")

	// Send disconnect message immediately
	err := w.Stream.Send(&pb.TaskStreamResponse{
		Payload: &pb.TaskStreamResponse_Ack{
			Ack: &pb.Ack{
				Message: "disconnect_requested",
			},
		},
	})

	if err != nil {
		log.Logger.Warnw("Failed to send disconnect message to group "+w.GroupName, "agent-disconnect", err)
	} else {
		log.Logger.Infow("Successfully sent disconnect message to group "+w.GroupName, "agent-disconnect", "send success")
	}

	// Cleanup channels
	for taskID, ch := range w.ResultChMap {
		close(ch)
		w.mu.Lock()
		delete(w.ResultChMap, taskID)
		w.mu.Unlock()
	}

	// Remove worker from group
	var newList []*AgentConnection
	for _, conn := range workers {
		if conn != w {
			newList = append(newList, conn)
		}
	}
	s.connectionList[w.GroupName] = newList
}

// SendTaskToAgent sends a task down a particular agent’s Stream.
// Returns a channel on which the result will be delivered.
func (s *AgentManager) SendTaskToAgent(ctx context.Context, taskName string, input []byte, groupName string) (*pb.TaskResult, error) {
	w := s.PickAgent(groupName)
	if w == nil {
		return nil, errors.New("agent unreachable")
	}

	// Generate a task ID.
	taskID := uuid.NewString()

	// Prepare a channel to receive the agent’s response.
	resultCh := make(chan *pb.TaskResult, 1)

	w.mu.Lock()
	w.ResultChMap[taskID] = resultCh
	w.mu.Unlock()

	// Actually send the task to the agent.
	err := w.Stream.Send(&pb.TaskStreamResponse{
		Payload: &pb.TaskStreamResponse_NewTask{
			NewTask: &pb.Task{
				Id:      taskID,
				Payload: taskName,
				Name:    taskName,
				Input:   string(input),
			},
		},
	})
	if err != nil {
		w.mu.Lock()
		delete(w.ResultChMap, taskID)
		w.mu.Unlock()
		return nil, err
	}

	defer func() {
		w.mu.Lock()
		delete(w.ResultChMap, taskID)
		w.mu.Unlock()
	}()

	// Wait for the agent to respond with a result or time out
	select {
	case res := <-resultCh:
		// Send response to the user
		if !res.Success {
			return nil, errors.New(res.Output)
		}
		return res, nil
	case <-time.After(60 * time.Second):
		return nil, errors.New("agent response timed out")
	}
}

func (s *AgentManager) SendPodLogsStreamReqToAgent(ctx context.Context, taskName string, input []byte, groupName string, conn *websocket.Conn) (*pb.LogsResult, error) {
	w := s.PickAgent(groupName)
	if w == nil {
		return nil, errors.New("agent unreachable")
	}

	// Generate a task ID.
	taskID := uuid.NewString()

	// Prepare a channel to receive the agent’s response.
	resultCh := make(chan *pb.LogsResult)

	w.mu.Lock()
	w.ResultStreamChMap[taskID] = resultCh
	w.mu.Unlock()

	// Actually send the task to the agent.
	err := w.Stream.Send(&pb.TaskStreamResponse{
		Payload: &pb.TaskStreamResponse_NewPodLogsStream{
			NewPodLogsStream: &pb.PodLogsStream{
				Id:      taskID,
				Payload: taskName,
				Name:    taskName,
				Input:   string(input),
			},
		},
	})
	if err != nil {
		log.Logger.Warnw("error")
		w.mu.Lock()
		delete(w.ResultStreamChMap, taskID)
		conn.Close()
		w.mu.Unlock()
		return nil, err
	}

	wsCtx, wsCancel := context.WithCancel(context.Background())
	// Wait for the agent to respond with a result or time out
	go func(conn *websocket.Conn, ctx context.Context, cancel context.CancelFunc) {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case res := <-resultCh:
				err = conn.WriteMessage(websocket.TextMessage, res.Output)
				if err != nil {
					log.Logger.Errorw("unable to write to websocket", "logs-stream", err.Error())
					cancel()
				}
			case <-ticker.C:
				if conn == nil {
					log.Logger.Warnw("conn is now nil", "logs-stream", taskID)
					cancel()
				} else {
					err = conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second))
					if err != nil {
						log.Logger.Warnw("WebSocket ping failed, closing connection", "logs-stream", taskID)
						cancel()
					}
					log.Logger.Infow("conn is active", "logs-stream", taskID)
				}
				// Send a message to the gRPC stream every 3 seconds
				err = w.Stream.Send(&pb.TaskStreamResponse{
					Payload: &pb.TaskStreamResponse_NewPodLogsStream{
						NewPodLogsStream: &pb.PodLogsStream{
							Id:      taskID,
							Payload: "heartbeat",
							Name:    taskName,
							Input:   string(input),
						},
					},
				})
				if err != nil {
					log.Logger.Errorw("unable to send heartbeat to agent", "logs-heartbeat", err)
					cancel()
				}
			case <-ctx.Done():
				log.Logger.Infow("cancelling log stream task", "logs-stream", taskID)
				_ = conn.Close()
				ticker.Stop()
				err = w.Stream.Send(&pb.TaskStreamResponse{
					Payload: &pb.TaskStreamResponse_NewPodLogsStream{
						NewPodLogsStream: &pb.PodLogsStream{
							Id:      taskID,
							Payload: "cancel",
							Name:    taskName,
							Input:   string(input),
						},
					},
				})
				if err != nil {
					log.Logger.Errorw("unable to send logs cancel to agent", "logs-cancel", err)
					cancel()
				}
				w.mu.Lock()
				delete(w.ResultStreamChMap, taskID)
				w.mu.Unlock()
				return
			}
		}
	}(conn, wsCtx, wsCancel)
	return nil, nil
}

// PickAgent returns any agent from the specified group (round-robin or random).
// For simplicity, let's just pick the first.
func (s *AgentManager) PickAgent(id string) *AgentConnection {
	s.mu.RLock()
	defer s.mu.RUnlock()
	agents := s.connectionList[id]
	if len(agents) == 0 {
		return nil
	}
	// naive pick: the first agent
	return agents[0]
}

func (s *AgentManager) SendPodLogsStreamReqToAgentForHttpStream(ctx *gin.Context, taskName string, input []byte, groupName string) (*pb.LogsResult, error) {
	w := s.PickAgent(groupName)
	if w == nil {
		return nil, errors.New("agent unreachable")
	}
	// Generate a task ID.
	taskID := uuid.NewString()

	// Prepare a channel to receive the agent’s response.
	resultCh := make(chan *pb.LogsResult)

	w.mu.Lock()
	w.ResultStreamChMap[taskID] = resultCh
	w.mu.Unlock()

	// Actually send the task to the agent.
	err := w.Stream.Send(&pb.TaskStreamResponse{
		Payload: &pb.TaskStreamResponse_NewPodLogsStream{
			NewPodLogsStream: &pb.PodLogsStream{
				Id:      taskID,
				Payload: taskName,
				Name:    taskName,
				Input:   string(input),
			},
		},
	})
	if err != nil {
		log.Logger.Warnw("error")
		w.mu.Lock()
		delete(w.ResultStreamChMap, taskID)
		w.mu.Unlock()
		return nil, err
	}

	wsCtx, wsCancel := context.WithCancel(context.Background())
	// Wait for the agent to respond with a result or time out
	func(gctx *gin.Context, ctx context.Context, cancel context.CancelFunc) {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case res := <-resultCh:
				_, err := gctx.Writer.Write(res.Output)
				if err != nil {
					log.Logger.Errorw("unable to write to HTTP stream", "logs-stream", err.Error())
					cancel()
				}
				gctx.Writer.Flush()
			case <-ticker.C:
				if gctx.Writer.Status() != http.StatusOK {
					// If the response status is not OK, the client may have disconnected
					fmt.Println("Client disconnected, closing stream.")
				}
				_, err = gctx.Writer.Write([]byte(""))
				if err != nil {
					log.Logger.Errorw("unable to write to HTTP stream -- stream may be closed by client", "logs-stream-health", err.Error())
					cancel()
				} else {
					log.Logger.Infow("conn is active", "logs-stream", taskID)
				}
				// Send a message to the gRPC stream every 3 seconds
				err = w.Stream.Send(&pb.TaskStreamResponse{
					Payload: &pb.TaskStreamResponse_NewPodLogsStream{
						NewPodLogsStream: &pb.PodLogsStream{
							Id:      taskID,
							Payload: "heartbeat",
							Name:    taskName,
							Input:   string(input),
						},
					},
				})
				if err != nil {
					log.Logger.Errorw("unable to send heartbeat to agent", "logs-heartbeat", err)
					cancel()
				}
			case <-gctx.Done():
				cancel()
			case <-ctx.Done():
				log.Logger.Infow("cancelling log stream task", "logs-stream", taskID)
				ticker.Stop()
				err = w.Stream.Send(&pb.TaskStreamResponse{
					Payload: &pb.TaskStreamResponse_NewPodLogsStream{
						NewPodLogsStream: &pb.PodLogsStream{
							Id:      taskID,
							Payload: "cancel",
							Name:    taskName,
							Input:   string(input),
						},
					},
				})
				if err != nil {
					log.Logger.Errorw("unable to send logs cancel to agent", "logs-cancel", err)
				}
				w.mu.Lock()
				delete(w.ResultStreamChMap, taskID)
				w.mu.Unlock()
				return
			}
		}
	}(ctx, wsCtx, wsCancel)
	return nil, nil
}
