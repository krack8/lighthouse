package grpc

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/krack8/lighthouse/controller/api/argocd"
	pbAgent "github.com/krack8/lighthouse/proto"
)

type AgentRegistrationService struct {
	pbAgent.UnimplementedAgentServiceServer
	agentManager *argocd.AgentManager
	mu           sync.RWMutex
	agentInfo    map[string]*AgentInfo
}

type AgentInfo struct {
	ID            string
	ClusterName   string
	Address       string
	Version       string
	Capabilities  *pbAgent.AgentCapabilities
	Status        *pbAgent.AgentStatus
	Metrics       *pbAgent.AgentMetrics
	Labels        map[string]string
	RegisteredAt  int64
	LastHeartbeat int64
}

func NewAgentRegistrationService(am *argocd.AgentManager) *AgentRegistrationService {
	service := &AgentRegistrationService{
		agentManager: am,
		agentInfo:    make(map[string]*AgentInfo),
	}

	// Start heartbeat monitor
	go service.monitorHeartbeats()

	return service
}

// RegisterAgent handles agent registration
func (s *AgentRegistrationService) RegisterAgent(ctx context.Context, req *pbAgent.RegisterAgentRequest) (*pbAgent.RegisterAgentResponse, error) {
	log.Printf("Agent registration request: ID=%s, Cluster=%s, Address=%s",
		req.AgentId, req.ClusterName, req.AgentAddress)

	// Validate request
	if req.AgentId == "" || req.ClusterName == "" || req.AgentAddress == "" {
		return nil, status.Errorf(codes.InvalidArgument, "agent_id, cluster_name, and agent_address are required")
	}

	// Check if agent already exists
	s.mu.RLock()
	if existing, exists := s.agentInfo[req.AgentId]; exists {
		s.mu.RUnlock()
		if existing.ClusterName != req.ClusterName {
			return nil, status.Errorf(codes.AlreadyExists,
				"agent %s already registered for cluster %s", req.AgentId, existing.ClusterName)
		}
		// Update existing registration
	} else {
		s.mu.RUnlock()
	}

	// Create gRPC connection to agent
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithTimeout(10 * time.Second),
		grpc.WithBlock(),
	}

	conn, err := grpc.Dial(req.AgentAddress, opts...)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable,
			"failed to connect to agent at %s: %v", req.AgentAddress, err)
	}

	// Register agent in manager (this manages the ArgoCD service client)
	if err := s.agentManager.RegisterAgent(req.AgentId, req.ClusterName, conn); err != nil {
		conn.Close()
		return nil, status.Errorf(codes.Internal, "failed to register agent: %v", err)
	}

	// Store agent information
	s.mu.Lock()
	s.agentInfo[req.AgentId] = &AgentInfo{
		ID:           req.AgentId,
		ClusterName:  req.ClusterName,
		Address:      req.AgentAddress,
		Version:      req.Version,
		Capabilities: req.Capabilities,
		Status: &pbAgent.AgentStatus{
			State:    "connected",
			Health:   "healthy",
			LastSeen: time.Now().Unix(),
		},
		Labels:        req.Labels,
		RegisteredAt:  time.Now().Unix(),
		LastHeartbeat: time.Now().Unix(),
	}
	s.mu.Unlock()

	log.Printf("Agent %s registered successfully for cluster %s", req.AgentId, req.ClusterName)

	return &pbAgent.RegisterAgentResponse{
		Success:           true,
		Message:           "Agent registered successfully",
		ControllerVersion: "1.0.0",
		HeartbeatInterval: 30, // 30 seconds
	}, nil
}

// UnregisterAgent handles agent unregistration
func (s *AgentRegistrationService) UnregisterAgent(ctx context.Context, req *pbAgent.UnregisterAgentRequest) (*pbAgent.UnregisterAgentResponse, error) {
	log.Printf("Agent unregistration request: ID=%s, Reason=%s", req.AgentId, req.Reason)

	// Remove from agent manager
	s.agentManager.UnregisterAgent(req.AgentId)

	// Remove from local storage
	s.mu.Lock()
	delete(s.agentInfo, req.AgentId)
	s.mu.Unlock()

	return &pbAgent.UnregisterAgentResponse{
		Success: true,
		Message: fmt.Sprintf("Agent %s unregistered successfully", req.AgentId),
	}, nil
}

// Heartbeat handles agent heartbeat
func (s *AgentRegistrationService) Heartbeat(ctx context.Context, req *pbAgent.HeartbeatRequest) (*pbAgent.HeartbeatResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	info, exists := s.agentInfo[req.AgentId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "agent %s not found", req.AgentId)
	}

	// Update agent status and metrics
	info.LastHeartbeat = time.Now().Unix()
	if req.Status != nil {
		info.Status = req.Status
		info.Status.LastSeen = time.Now().Unix()
	}
	if req.Metrics != nil {
		info.Metrics = req.Metrics
	}

	// Update agent manager status
	if agent, err := s.agentManager.GetAgent(req.AgentId); err == nil {
		agent.Connected = info.Status.State == "connected"
	}

	// Return any pending commands
	commands := s.getPendingCommands(req.AgentId)

	return &pbAgent.HeartbeatResponse{
		Success:  true,
		Commands: commands,
	}, nil
}

// GetAgentInfo returns information about a specific agent
func (s *AgentRegistrationService) GetAgentInfo(ctx context.Context, req *pbAgent.GetAgentInfoRequest) (*pbAgent.GetAgentInfoResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	info, exists := s.agentInfo[req.AgentId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "agent %s not found", req.AgentId)
	}

	return &pbAgent.GetAgentInfoResponse{
		Agent: &pbAgent.AgentInfo{
			AgentId:      info.ID,
			ClusterName:  info.ClusterName,
			AgentAddress: info.Address,
			Version:      info.Version,
			Capabilities: info.Capabilities,
			Status:       info.Status,
			Metrics:      info.Metrics,
			Labels:       info.Labels,
			RegisteredAt: info.RegisteredAt,
		},
	}, nil
}

// UpdateAgentStatus updates the status of an agent
func (s *AgentRegistrationService) UpdateAgentStatus(ctx context.Context, req *pbAgent.UpdateAgentStatusRequest) (*pbAgent.UpdateAgentStatusResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	info, exists := s.agentInfo[req.AgentId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "agent %s not found", req.AgentId)
	}

	info.Status = req.Status
	info.Status.LastSeen = time.Now().Unix()

	return &pbAgent.UpdateAgentStatusResponse{
		Success: true,
	}, nil
}

// monitorHeartbeats checks for agents that haven't sent heartbeats
func (s *AgentRegistrationService) monitorHeartbeats() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		s.mu.Lock()
		now := time.Now().Unix()

		for agentID, info := range s.agentInfo {
			// If no heartbeat for 2 minutes, mark as disconnected
			if now-info.LastHeartbeat > 120 {
				log.Printf("Agent %s missed heartbeat, marking as disconnected", agentID)

				info.Status.State = "disconnected"
				info.Status.Health = "unhealthy"
				info.Status.Message = "Heartbeat timeout"

				// Update agent manager
				if agent, err := s.agentManager.GetAgent(agentID); err == nil {
					agent.Connected = false
				}
			}
		}
		s.mu.Unlock()
	}
}

// getPendingCommands returns any pending commands for an agent
func (s *AgentRegistrationService) getPendingCommands(agentID string) []*pbAgent.AgentCommand {
	// This could be implemented to fetch commands from a queue or database
	// For now, return empty
	return []*pbAgent.AgentCommand{}
}

// GetAllAgents returns information about all registered agents
func (s *AgentRegistrationService) GetAllAgents() []*AgentInfo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	agents := make([]*AgentInfo, 0, len(s.agentInfo))
	for _, info := range s.agentInfo {
		agents = append(agents, info)
	}

	return agents
}

// GetAgentsByCluster returns agents for a specific cluster
func (s *AgentRegistrationService) GetAgentsByCluster(clusterName string) []*AgentInfo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var agents []*AgentInfo
	for _, info := range s.agentInfo {
		if info.ClusterName == clusterName {
			agents = append(agents, info)
		}
	}

	return agents
}
