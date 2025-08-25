package argocd

import (
	"github.com/krack8/lighthouse/shared"
)

// Request types specific to controller with validation
// These already use shared types for nested structures

// CreateApplicationRequest for creating application
type CreateApplicationRequest struct {
	ClusterID   string                        `json:"cluster_id" binding:"required"`
	Name        string                        `json:"name" binding:"required"`
	Namespace   string                        `json:"namespace" binding:"required"`
	Project     string                        `json:"project" binding:"required"`
	Source      shared.ApplicationSource      `json:"source" binding:"required"`      // Using shared type
	Destination shared.ApplicationDestination `json:"destination" binding:"required"` // Using shared type
}

// UpdateApplicationRequest for updating application
type UpdateApplicationRequest struct {
	ClusterID   string                        `json:"cluster_id" binding:"required"`
	Namespace   string                        `json:"namespace"`
	Project     string                        `json:"project"`
	Source      shared.ApplicationSource      `json:"source"`      // Using shared type
	Destination shared.ApplicationDestination `json:"destination"` // Using shared type
}

// SyncApplicationRequest for syncing application
type SyncApplicationRequest struct {
	ClusterID string `json:"cluster_id"`
	Prune     bool   `json:"prune"`
	DryRun    bool   `json:"dry_run"`
}

// CreateProjectRequest for creating project
type CreateProjectRequest struct {
	ClusterID   string   `json:"cluster_id" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	SourceRepos []string `json:"source_repos"`
}

// CreateRepositoryRequest for creating repository
type CreateRepositoryRequest struct {
	ClusterID string `json:"cluster_id" binding:"required"`
	URL       string `json:"url" binding:"required"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Type      string `json:"type" binding:"required,oneof=git helm"`
}

// ErrorResponse for API errors
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

// SuccessResponse for API success
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
