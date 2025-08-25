package argocd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/krack8/lighthouse/shared"
)

// Agent represents a connected agent
type Agent struct {
	ID          string            `json:"id"`
	ClusterName string            `json:"cluster_name"`
	URL         string            `json:"url"`
	Connected   bool              `json:"connected"`
	LastSeen    time.Time         `json:"last_seen"`
	Version     string            `json:"version"`
	Labels      map[string]string `json:"labels"`
	HTTPClient  *http.Client      `json:"-"`
}

// AgentManager manages all connected agents
type AgentManager struct {
	agents map[string]*Agent
	mu     sync.RWMutex
}

// NewAgentManager creates a new agent manager
func NewAgentManager() *AgentManager {
	am := &AgentManager{
		agents: make(map[string]*Agent),
	}

	// Start health check routine
	go am.healthCheck()

	return am
}

// RegisterAgent registers a new agent
func (am *AgentManager) RegisterAgent(req shared.AgentRegistrationRequest) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	// Create HTTP client for this agent
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Test connection to agent
	healthURL := fmt.Sprintf("%s/health", req.AgentURL)
	resp, err := httpClient.Get(healthURL)
	if err != nil {
		return fmt.Errorf("failed to connect to agent: %w", err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("agent health check failed with status: %d", resp.StatusCode)
	}

	// Register the agent
	agent := &Agent{
		ID:          req.AgentID,
		ClusterName: req.ClusterName,
		URL:         req.AgentURL,
		Connected:   true,
		LastSeen:    time.Now(),
		Version:     req.Version,
		Labels:      req.Labels,
		HTTPClient:  httpClient,
	}

	am.agents[req.AgentID] = agent

	return nil
}

// UnregisterAgent removes an agent
func (am *AgentManager) UnregisterAgent(agentID string) {
	am.mu.Lock()
	defer am.mu.Unlock()

	delete(am.agents, agentID)
}

// GetAgent retrieves an agent by ID
func (am *AgentManager) GetAgent(agentID string) (*Agent, error) {
	am.mu.RLock()
	defer am.mu.RUnlock()

	agent, exists := am.agents[agentID]
	if !exists {
		return nil, fmt.Errorf("agent %s not found", agentID)
	}

	if !agent.Connected {
		return nil, fmt.Errorf("agent %s is disconnected", agentID)
	}

	return agent, nil
}

// GetAgentByCluster retrieves an agent by cluster name
func (am *AgentManager) GetAgentByCluster(clusterName string) (*Agent, error) {
	am.mu.RLock()
	defer am.mu.RUnlock()

	for _, agent := range am.agents {
		if agent.ClusterName == clusterName && agent.Connected {
			return agent, nil
		}
	}

	return nil, fmt.Errorf("no agent found for cluster %s", clusterName)
}

// ListAgents returns all connected agents
func (am *AgentManager) ListAgents() []*Agent {
	am.mu.RLock()
	defer am.mu.RUnlock()

	agents := make([]*Agent, 0, len(am.agents))
	for _, agent := range am.agents {
		if agent.Connected {
			agents = append(agents, agent)
		}
	}

	return agents
}

// UpdateHeartbeat updates the last seen time for an agent
func (am *AgentManager) UpdateHeartbeat(agentID string) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	agent, exists := am.agents[agentID]
	if !exists {
		return fmt.Errorf("agent %s not found", agentID)
	}

	agent.LastSeen = time.Now()
	agent.Connected = true

	return nil
}

// healthCheck periodically checks agent health
func (am *AgentManager) healthCheck() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		am.mu.Lock()
		for _, agent := range am.agents {
			// Mark as disconnected if no heartbeat for 60 seconds
			if time.Since(agent.LastSeen) > 60*time.Second {
				agent.Connected = false
			}
		}
		am.mu.Unlock()
	}
}

// CallAgent makes an HTTP call to an agent
func (agent *Agent) CallAgent(method, endpoint string, body interface{}, response interface{}) error {
	url := fmt.Sprintf("%s%s", agent.URL, endpoint)

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := agent.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call agent: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("agent returned error: %s - %s", resp.Status, string(respBody))
	}

	if response != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, response); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}
