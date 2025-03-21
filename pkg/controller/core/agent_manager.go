package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/krack8/lighthouse/pkg/common/consts"
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
	GroupName             string
	ResultChMap           map[string]chan *pb.TaskResult           // map of taskID -> channel that receives result
	TerminalExecRespChMap map[string]chan *pb.TerminalExecResponse // map of taskID -> channel that receives result
	mu                    sync.Mutex
}

var webSocketClientMap = make(map[string]*WebSocketClient) // Key: TaskID -> WebsocketClient
var ws_mu sync.Mutex                                       // Mutex to synchronize access to connections map

type WebSocketClient struct {
	Conn            *websocket.Conn
	IsConnected     bool
	CloseSignal     chan string
	ReconnectWait   time.Duration
	HeartbeatTicker *time.Ticker
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
		log.Logger.Warnw("Invalid agent connection", "agent-disconnect", "group: "+w.GroupName)
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

// SendTerminalExecRequestToAgent sends a terminal exec request to a particular agent’s Stream.
// Returns a channel on which the result will be delivered.
func (s *AgentManager) SendTerminalExecRequestToAgent(ctx *gin.Context, taskID string, input string, groupName string, isReconnect bool) (*pb.TerminalExecResponse, error) {
	if isReconnect == true && webSocketClientMap[taskID] == nil {
		log.Logger.Errorw(fmt.Sprintf("Unable to reconnect websocket: no previous connection exists"), "TaskType", "PodExec", "AgentGroup", groupName, "TaskID", taskID)
		return nil, errors.New("unable to reconnect websocket")
	}

	w := s.PickAgent(groupName)
	if w == nil {
		log.Logger.Errorw(fmt.Sprintf("Unable to get agent: agent unreachable"), "TaskType", "PodExec", "AgentGroup", groupName)
		return nil, errors.New("agent unreachable")
	}

	var wsocket = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// Checking if websocket connection already exists for taskID
	// If new connection then initiate agent connection
	if webSocketClientMap[taskID] == nil {
		// Websocket initialization request
		if w.TerminalExecRespChMap[taskID] != nil {
			log.Logger.Errorw(fmt.Sprintf("Mismatched websocket and agent pod exec connection"), "TaskType", "PodExec", "AgentGroup", groupName, "TaskID", taskID)
			return nil, errors.New("unable to initiate websocket connection")
		}

		conn, err := wsocket.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Logger.Errorw(fmt.Sprintf("Unable to initiate websocket connection"), "TaskType", "PodExec", "AgentGroup", groupName, "TaskID", taskID)
			return nil, errors.New("unable to initiate websocket connection")
		}

		// Initiating new websocket Connection
		// Prepare a channel to receive the agent’s response.
		resultCh := make(chan *pb.TerminalExecResponse, 1)

		w.mu.Lock()
		w.TerminalExecRespChMap[taskID] = resultCh
		w.mu.Unlock()

		// Sending an init connection message
		err = w.Stream.Send(&pb.TaskStreamResponse{
			Payload: &pb.TaskStreamResponse_ExecReq{
				ExecReq: &pb.TerminalExecRequest{
					TaskId:  taskID,
					Input:   input,
					Command: []byte{},
					Payload: consts.TaskPodExecInitConn,
				},
			},
		})

		if err != nil {
			w.mu.Lock()
			log.Logger.Errorw(fmt.Sprintf("Closing Connection! Unable to initiate connection: %s", err.Error()), "TaskID", taskID, "TaskType", "PodExec")
			_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			delete(w.TerminalExecRespChMap, taskID)
			conn.Close()
			w.mu.Unlock()
			return nil, err
		}

		log.Logger.Infow(fmt.Sprintf("Sending initial TaskID to websocket..."), "TaskID", taskID, "TaskType", "PodExec")

		message := taskID
		err = conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			w.mu.Lock()
			log.Logger.Errorw(fmt.Sprintf("Closing Connection! Unable to send initial TaskID to websocket: %s", err.Error()), "TaskID", taskID, "TaskType", "PodExec")
			_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			delete(w.TerminalExecRespChMap, taskID)
			conn.Close()
			w.mu.Unlock()
			return nil, err
		}

		ws_mu.Lock()
		webSocketClientMap[taskID] = &WebSocketClient{
			Conn:          conn,
			IsConnected:   true,
			CloseSignal:   make(chan string),
			ReconnectWait: 10 * time.Second,
		}
		ws_mu.Unlock()

	} else {
		// Websocket reconnection request
		// Check if agent connection is active or not, if not then close the websocket connection
		if w.TerminalExecRespChMap[taskID] == nil {
			log.Logger.Errorw(fmt.Sprintf("Unable to initiate websocket connection, closing connection!"), "TaskType", "PodExec", "AgentGroup", groupName, "TaskID", taskID)
			return nil, errors.New("unable to initiate websocket connection")
		}

		// Replacing old connection with new one
		conn, err := wsocket.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Logger.Errorw(fmt.Sprintf("Unable to initiate websocket connection"), "TaskType", "PodExec", "AgentGroup", groupName, "TaskID", taskID)
			return nil, errors.New("unable to initiate websocket connection")
		}

		ws_mu.Lock()
		webSocketClientMap[taskID].Conn = conn
		webSocketClientMap[taskID].IsConnected = true
		webSocketClientMap[taskID].CloseSignal <- "wait_reconnect"
		ws_mu.Unlock()
	}

	// Send the task to the agent.
	// Goroutine to listen for websocket input and send to Agent
	go func(client *WebSocketClient) {
		defer func() {
			client.CloseSignal <- "wait_reconnect"
		}()
		for {
			_, command, err := client.Conn.ReadMessage()
			if err != nil {
				if w.TerminalExecRespChMap[taskID] == nil {
					client.CloseSignal <- "force_close"
					return
				}
				log.Logger.Errorw(fmt.Sprintf("WebSocket read error: %s", err.Error()), "TaskID", taskID, "TaskType", "PodExec")
				return

			} else {
				err := w.Stream.Send(&pb.TaskStreamResponse{
					Payload: &pb.TaskStreamResponse_ExecReq{
						ExecReq: &pb.TerminalExecRequest{
							TaskId:  taskID,
							Input:   input,
							Command: command,
							Payload: consts.TaskPodExecCommand,
						},
					},
				})
				if err != nil {
					log.Logger.Errorw(fmt.Sprintf("Unable to send command to agent: %s", err.Error()), "TaskID", taskID, "TaskType", "PodExec")
					client.CloseSignal <- "force_close"
					return
				}
			}
		}
	}(webSocketClientMap[taskID])

	// Wait for the agent to respond with a result or time out
	go func(client *WebSocketClient, resultCh chan *pb.TerminalExecResponse) {
		// Create a ticker for sending messages every 3 seconds
		client.HeartbeatTicker = time.NewTicker(3 * time.Second)

		for {
			select {
			case res := <-resultCh:
				if res.Success == false {
					log.Logger.Errorw(fmt.Sprintf("Error response from agent: %s", string(res.Output)), "TaskID", taskID, "TaskType", "PodExec", "Response", res)
					client.CloseSignal <- "force_close"
					return

				} else {
					ws_mu.Lock()
					if client.IsConnected {
						err := client.Conn.WriteMessage(websocket.TextMessage, res.Output)
						if err != nil {
							log.Logger.Errorw(fmt.Sprintf("Unable to write message to websocket: %s", err.Error()), "TaskID", taskID, "TaskType", "PodExec", "Response", res)
							ws_mu.Unlock()
							client.CloseSignal <- "wait_reconnect"
							return
						}
					}
					ws_mu.Unlock()
				}
			case <-client.HeartbeatTicker.C:
				// Send a message to the gRPC stream every 3 seconds
				w.mu.Lock()
				err := w.Stream.Send(&pb.TaskStreamResponse{
					Payload: &pb.TaskStreamResponse_ExecReq{
						ExecReq: &pb.TerminalExecRequest{
							TaskId:  taskID,
							Input:   input,
							Command: []byte{},
							Payload: consts.TaskPodExecHeartbeat,
						},
					},
				})

				if err != nil {
					log.Logger.Errorw(fmt.Sprintf("Unable to send heartbeat to agent: %s", err.Error()), "TaskID", taskID, "TaskType", "PodExec")
					w.mu.Unlock()
					client.CloseSignal <- "force_close"
					return
				}
				w.mu.Unlock()
			}
		}
	}(webSocketClientMap[taskID], w.TerminalExecRespChMap[taskID])

	// Goroutine to disconnect websocket connection
	go func(client *WebSocketClient) {
		// Wait for a signal to disconnect (either force_close or wait_reconnect)
		action := <-client.CloseSignal

		log.Logger.Infow("Received closed connection signal", "TaskID", taskID, "TaskType", "PodExec", "Action", action)
		ws_mu.Lock()
		if !client.IsConnected {
			ws_mu.Unlock()
			log.Logger.Infow("Websocket Client is not connected, Exiting...", "TaskID", taskID, "TaskType", "PodExec")
			return
		}
		client.IsConnected = false
		ws_mu.Unlock()

		if action == "force_close" {
			ws_mu.Lock()
			_ = client.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			log.Logger.Infow("Disconnecting Websocket client forcefully..", "TaskID", taskID, "TaskType", "PodExec")
			client.Conn.Close()
			client.HeartbeatTicker.Stop()
			delete(webSocketClientMap, taskID)
			ws_mu.Unlock()
			w.mu.Lock()
			delete(w.TerminalExecRespChMap, taskID)
			w.mu.Unlock()
			return

		} else if action == "wait_reconnect" {
			log.Logger.Infow("Connection lost. Waiting for Websocket client reconnection...", "TaskID", taskID, "TaskType", "PodExec")

			// Wait for reconnection for 10 seconds
			timer := time.NewTimer(client.ReconnectWait)
			defer timer.Stop()

			select {
			case <-timer.C:
				w.mu.Lock()
				ws_mu.Lock()
				log.Logger.Infow("No reconnection within 10 seconds, closing connection.", "TaskID", taskID, "TaskType", "PodExec")
				_ = w.Stream.Send(&pb.TaskStreamResponse{
					Payload: &pb.TaskStreamResponse_ExecReq{
						ExecReq: &pb.TerminalExecRequest{
							TaskId:  taskID,
							Input:   input,
							Command: []byte{},
							Payload: consts.TaskPodExecCloseConn,
						},
					},
				})
				_ = client.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				client.Conn.Close()
				client.HeartbeatTicker.Stop()
				delete(w.TerminalExecRespChMap, taskID)
				delete(webSocketClientMap, taskID)
				ws_mu.Unlock()
				w.mu.Unlock()

			case <-client.CloseSignal:
				log.Logger.Infow("Reconnection attempt detected, keeping connection open.", "TaskID", taskID, "TaskType", "PodExec")
			}
		}
	}(webSocketClientMap[taskID])

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

func (s *AgentManager) CloseWebsocketConnectionByTask(taskId string) {
	ws_mu.Lock()
	defer ws_mu.Unlock()
	if webSocketClientMap[taskId] != nil {
		webSocketClientMap[taskId].CloseSignal <- "force_close"
	}
	return
}
