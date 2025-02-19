package core

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"log"
	"sync"
	"time"
)

// AgentConnection represents one agent's active Streaming connection.
type AgentConnection struct {
	Stream pb.Controller_TaskStreamServer
	//UniqueId    string
	GroupName   string
	ResultChMap map[string]chan *pb.TaskResult // map of taskID -> channel that receives result
	mu          sync.Mutex
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
	log.Printf("Agent added to connection list %q. Total agents in group: %d",
		w.GroupName, len(s.connectionList[w.GroupName]))
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
	log.Printf("Agent removed from connection list %q. Remaining agents: %d",
		w.GroupName, len(s.connectionList[w.GroupName]))
}

func (s *AgentManager) RemoveAgentByGroupName(groupName string) bool {
	s.mu.Lock()
	workers, exists := s.connectionList[groupName]
	if !exists || len(workers) == 0 {
		log.Printf("No agent found in group: %s", groupName)
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

// disconnectWorker handles immediate worker disconnection
func (s *AgentManager) disconnectWorker(w *AgentConnection) {
	if w == nil || w.Stream == nil {
		log.Printf("Invalid worker connection")
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
		log.Printf("Worker not found in group %s", w.GroupName)
		return
	}

	log.Printf("Sending disconnect message - Group: %s, Total workers in group: %d",
		w.GroupName, len(workers))

	// Send disconnect message immediately
	err := w.Stream.Send(&pb.TaskStreamResponse{
		Payload: &pb.TaskStreamResponse_Ack{
			Ack: &pb.Ack{
				Message: "disconnect_requested",
			},
		},
	})

	if err != nil {
		log.Printf("Failed to send disconnect message to group %s: %v", w.GroupName, err)
	} else {
		log.Printf("Successfully sent disconnect message to group %s", w.GroupName)
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
