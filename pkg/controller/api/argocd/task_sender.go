package argocd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/krack8/lighthouse/pkg/common/pb"
	"github.com/krack8/lighthouse/pkg/controller/core"
)

// SendArgoCDTask sends an ArgoCD task to an agent via gRPC
func SendArgoCDTask(ctx context.Context, agentGroup, taskName string, payload interface{}) (*pb.TaskResult, error) {
	// Convert payload to JSON bytes
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Send task using the existing SendTaskToAgent function
	// taskName goes as the task name (e.g., "argocd:list_applications")
	// payloadJSON goes as the input bytes
	return core.GetAgentManager().SendTaskToAgent(ctx, taskName, payloadJSON, agentGroup)
}

// Helper functions for specific ArgoCD operations
func SendListApplicationsTask(ctx context.Context, agentGroup, project string) (*pb.TaskResult, error) {
	payload := map[string]string{
		"project": project,
	}
	return SendArgoCDTask(ctx, agentGroup, "argocd:list_applications", payload)
}

func SendGetApplicationTask(ctx context.Context, agentGroup, name string) (*pb.TaskResult, error) {
	payload := map[string]string{
		"name": name,
	}
	return SendArgoCDTask(ctx, agentGroup, "argocd:get_application", payload)
}

func SendCreateApplicationTask(ctx context.Context, agentGroup string, app interface{}) (*pb.TaskResult, error) {
	return SendArgoCDTask(ctx, agentGroup, "argocd:create_application", app)
}

func SendSyncApplicationTask(ctx context.Context, agentGroup, name, revision string, prune, dryRun bool) (*pb.TaskResult, error) {
	payload := map[string]interface{}{
		"name":     name,
		"revision": revision,
		"prune":    prune,
		"dryRun":   dryRun,
	}
	return SendArgoCDTask(ctx, agentGroup, "argocd:sync_application", payload)
}

func SendDeleteApplicationTask(ctx context.Context, agentGroup, name string, cascade bool) (*pb.TaskResult, error) {
	payload := map[string]interface{}{
		"name":    name,
		"cascade": cascade,
	}
	return SendArgoCDTask(ctx, agentGroup, "argocd:delete_application", payload)
}
