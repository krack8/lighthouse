package grpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pbAgent "github.com/krack8/lighthouse/proto/agent"
	"google.golang.org/grpc"
)

type ControllerClient struct {
	conn    *grpc.ClientConn
	client  pbAgent.AgentServiceClient
	agentID string
}

func NewControllerClient(controllerAddr, agentID string) (*ControllerClient, error) {
	conn, err := grpc.Dial(controllerAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &ControllerClient{
		conn:    conn,
		client:  pbAgent.NewAgentServiceClient(conn),
		agentID: agentID,
	}, nil
}

func (c *ControllerClient) Register(clusterName, agentAddr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.client.RegisterAgent(ctx, &pbAgent.RegisterAgentRequest{
		AgentId:      c.agentID,
		ClusterName:  clusterName,
		AgentAddress: agentAddr,
		Version:      "1.0.0",
		Capabilities: &pbAgent.AgentCapabilities{
			SupportsArgocd:     true,
			SupportsKubernetes: true,
			SupportsHelm:       true,
			ArgocdVersion:      "2.11.0",
			KubernetesVersion:  "1.28.0",
		},
	})

	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("registration failed: %s", resp.Message)
	}

	log.Printf("Registered with controller successfully. Heartbeat interval: %d seconds",
		resp.HeartbeatInterval)

	// Start heartbeat
	go c.startHeartbeat(resp.HeartbeatInterval)

	return nil
}

func (c *ControllerClient) startHeartbeat(interval int64) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		resp, err := c.client.Heartbeat(ctx, &pbAgent.HeartbeatRequest{
			AgentId: c.agentID,
			Status: &pbAgent.AgentStatus{
				State:  "connected",
				Health: "healthy",
			},
			Metrics: &pbAgent.AgentMetrics{
				ApplicationsCount: 10,
				ProjectsCount:     3,
				RepositoriesCount: 5,
			},
		})

		cancel()

		if err != nil {
			log.Printf("Heartbeat failed: %v", err)
			continue
		}

		// Process any commands from controller
		for _, cmd := range resp.Commands {
			log.Printf("Received command: %s", cmd.Command)
			// Process command
		}
	}
}

func (c *ControllerClient) Close() {
	c.conn.Close()
}
