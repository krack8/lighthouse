package shared

import "time"

// ============== ArgoCD Application Types ==============

// Application represents an ArgoCD application
type Application struct {
	Name         string                 `json:"name"`
	Namespace    string                 `json:"namespace"`
	Project      string                 `json:"project"`
	Source       ApplicationSource      `json:"source"`
	Destination  ApplicationDestination `json:"destination"`
	SyncStatus   SyncStatus             `json:"sync_status"`
	HealthStatus HealthStatus           `json:"health_status"`
	CreatedAt    *time.Time             `json:"created_at,omitempty"`
	UpdatedAt    *time.Time             `json:"updated_at,omitempty"`
}

// ApplicationSource defines the source of an application
type ApplicationSource struct {
	RepoURL        string      `json:"repo_url"`
	Path           string      `json:"path"`
	TargetRevision string      `json:"target_revision"`
	Helm           *HelmSource `json:"helm,omitempty"`
	Kustomize      *Kustomize  `json:"kustomize,omitempty"`
}

// HelmSource defines helm specific options
type HelmSource struct {
	ValueFiles []string        `json:"value_files,omitempty"`
	Parameters []HelmParameter `json:"parameters,omitempty"`
	Values     string          `json:"values,omitempty"`
}

// HelmParameter is a helm parameter
type HelmParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Kustomize defines kustomize specific options
type Kustomize struct {
	NamePrefix string   `json:"name_prefix,omitempty"`
	NameSuffix string   `json:"name_suffix,omitempty"`
	Images     []string `json:"images,omitempty"`
}

// ApplicationDestination defines the destination of an application
type ApplicationDestination struct {
	Server    string `json:"server"`
	Namespace string `json:"namespace"`
}

// SyncStatus contains sync status information
type SyncStatus struct {
	Status     string     `json:"status"`
	ComparedTo string     `json:"compared_to,omitempty"`
	Revision   string     `json:"revision,omitempty"`
	SyncedAt   *time.Time `json:"synced_at,omitempty"`
}

// HealthStatus contains health status information
type HealthStatus struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// ============== ArgoCD Project Types ==============

// Project represents an ArgoCD project
type Project struct {
	Name                       string                   `json:"name"`
	Description                string                   `json:"description"`
	SourceRepos                []string                 `json:"source_repos"`
	Destinations               []ApplicationDestination `json:"destinations,omitempty"`
	ClusterResourceWhitelist   []GroupKind              `json:"cluster_resource_whitelist,omitempty"`
	NamespaceResourceWhitelist []GroupKind              `json:"namespace_resource_whitelist,omitempty"`
	Roles                      []ProjectRole            `json:"roles,omitempty"`
	CreatedAt                  *time.Time               `json:"created_at,omitempty"`
}

// GroupKind defines a Kubernetes group/kind
type GroupKind struct {
	Group string `json:"group"`
	Kind  string `json:"kind"`
}

// ProjectRole defines a project role
type ProjectRole struct {
	Name     string   `json:"name"`
	Policies []string `json:"policies"`
	Groups   []string `json:"groups,omitempty"`
}

// ============== ArgoCD Repository Types ==============

// Repository represents an ArgoCD repository
type Repository struct {
	URL                   string `json:"url"`
	Name                  string `json:"name,omitempty"`
	Type                  string `json:"type"`
	Username              string `json:"username,omitempty"`
	Password              string `json:"password,omitempty"`
	SSHPrivateKey         string `json:"ssh_private_key,omitempty"`
	ConnectionState       string `json:"connection_state,omitempty"`
	Status                string `json:"status,omitempty"`
	Message               string `json:"message,omitempty"`
	Project               string `json:"project,omitempty"`
	EnableLFS             bool   `json:"enable_lfs,omitempty"`
	InsecureIgnoreHostKey bool   `json:"insecure_ignore_host_key,omitempty"`
}

// ============== Request/Response Types ==============

// ListApplicationsRequest for listing applications
type ListApplicationsRequest struct {
	Project   string            `json:"project,omitempty"`
	Namespace string            `json:"namespace,omitempty"`
	Labels    map[string]string `json:"labels,omitempty"`
}

// ListApplicationsResponse contains list of applications
type ListApplicationsResponse struct {
	Applications []Application `json:"applications"`
	Total        int           `json:"total"`
}

// CreateApplicationRequest for creating application
type CreateApplicationRequest struct {
	Name        string                 `json:"name"`
	Namespace   string                 `json:"namespace"`
	Project     string                 `json:"project"`
	Source      ApplicationSource      `json:"source"`
	Destination ApplicationDestination `json:"destination"`
	SyncPolicy  *SyncPolicy            `json:"sync_policy,omitempty"`
}

// UpdateApplicationRequest for updating application
type UpdateApplicationRequest struct {
	Namespace   string                 `json:"namespace,omitempty"`
	Project     string                 `json:"project,omitempty"`
	Source      ApplicationSource      `json:"source,omitempty"`
	Destination ApplicationDestination `json:"destination,omitempty"`
	SyncPolicy  *SyncPolicy            `json:"sync_policy,omitempty"`
}

// SyncPolicy defines sync policy
type SyncPolicy struct {
	Automated   *SyncPolicyAutomated `json:"automated,omitempty"`
	SyncOptions []string             `json:"sync_options,omitempty"`
	Retry       *RetryStrategy       `json:"retry,omitempty"`
}

// SyncPolicyAutomated defines automated sync policy
type SyncPolicyAutomated struct {
	Prune    bool `json:"prune"`
	SelfHeal bool `json:"self_heal"`
}

// RetryStrategy defines retry strategy
type RetryStrategy struct {
	Limit   int64    `json:"limit"`
	Backoff *Backoff `json:"backoff,omitempty"`
}

// Backoff defines backoff strategy
type Backoff struct {
	Duration    string `json:"duration"`
	Factor      int64  `json:"factor,omitempty"`
	MaxDuration string `json:"max_duration,omitempty"`
}

// SyncApplicationRequest for syncing application
type SyncApplicationRequest struct {
	Name      string            `json:"name"`
	Revision  string            `json:"revision,omitempty"`
	Prune     bool              `json:"prune"`
	DryRun    bool              `json:"dry_run"`
	Strategy  *SyncStrategy     `json:"strategy,omitempty"`
	Resources []ResourceDetails `json:"resources,omitempty"`
}

// SyncStrategy defines sync strategy
type SyncStrategy struct {
	Apply *ApplyStrategy `json:"apply,omitempty"`
	Hook  *HookStrategy  `json:"hook,omitempty"`
}

// ApplyStrategy defines apply strategy
type ApplyStrategy struct {
	Force bool `json:"force"`
}

// HookStrategy defines hook strategy
type HookStrategy struct {
	Force bool `json:"force"`
}

// ResourceDetails defines resource details for sync
type ResourceDetails struct {
	Group     string `json:"group,omitempty"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}

// SyncApplicationResponse for sync response
type SyncApplicationResponse struct {
	Status    string           `json:"status"`
	Message   string           `json:"message"`
	Revision  string           `json:"revision,omitempty"`
	Resources []ResourceStatus `json:"resources,omitempty"`
}

// ResourceStatus defines resource sync status
type ResourceStatus struct {
	Group     string `json:"group,omitempty"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
	Status    string `json:"status"`
	Message   string `json:"message,omitempty"`
	Hook      bool   `json:"hook,omitempty"`
}

// DeleteApplicationRequest for deleting application
type DeleteApplicationRequest struct {
	Name              string `json:"name"`
	Cascade           bool   `json:"cascade"`
	PropagationPolicy string `json:"propagation_policy,omitempty"`
}

// DeleteApplicationResponse for delete response
type DeleteApplicationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// RollbackApplicationRequest for rollback
type RollbackApplicationRequest struct {
	Name     string `json:"name"`
	Revision string `json:"revision"`
	DryRun   bool   `json:"dry_run"`
}

// ============== Agent Registration Types ==============

// AgentRegistrationRequest for agent registration
type AgentRegistrationRequest struct {
	AgentID      string            `json:"agent_id"`
	ClusterName  string            `json:"cluster_name"`
	AgentURL     string            `json:"agent_url"`
	Version      string            `json:"version"`
	Labels       map[string]string `json:"labels,omitempty"`
	Capabilities AgentCapabilities `json:"capabilities,omitempty"`
}

// AgentRegistrationResponse for registration response
type AgentRegistrationResponse struct {
	Success           bool   `json:"success"`
	Message           string `json:"message"`
	HeartbeatInterval int    `json:"heartbeat_interval"`
	ControllerVersion string `json:"controller_version,omitempty"`
}

// AgentCapabilities defines agent capabilities
type AgentCapabilities struct {
	SupportsArgoCD     bool   `json:"supports_argocd"`
	SupportsKubernetes bool   `json:"supports_kubernetes"`
	SupportsHelm       bool   `json:"supports_helm"`
	ArgoCDVersion      string `json:"argocd_version,omitempty"`
	KubernetesVersion  string `json:"kubernetes_version,omitempty"`
}

// HeartbeatRequest for agent heartbeat
type HeartbeatRequest struct {
	AgentID string       `json:"agent_id"`
	Status  string       `json:"status"`
	Health  string       `json:"health"`
	Metrics AgentMetrics `json:"metrics,omitempty"`
}

// HeartbeatResponse for heartbeat response
type HeartbeatResponse struct {
	Success  bool           `json:"success"`
	Commands []AgentCommand `json:"commands,omitempty"`
}

// AgentMetrics contains agent metrics
type AgentMetrics struct {
	ApplicationsCount int     `json:"applications_count"`
	ProjectsCount     int     `json:"projects_count"`
	RepositoriesCount int     `json:"repositories_count"`
	CPUUsage          float64 `json:"cpu_usage"`
	MemoryUsage       float64 `json:"memory_usage"`
	DiskUsage         float64 `json:"disk_usage"`
}

// AgentCommand defines a command for agent
type AgentCommand struct {
	ID         string            `json:"id"`
	Command    string            `json:"command"`
	Parameters map[string]string `json:"parameters"`
	Timestamp  time.Time         `json:"timestamp"`
}

// ============== Common Response Types ==============

// GenericResponse for generic API responses
type GenericResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse for API errors
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

// PaginationRequest for paginated requests
type PaginationRequest struct {
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
	SortBy   string `json:"sort_by,omitempty"`
	Order    string `json:"order,omitempty"`
}

// PaginationResponse for paginated responses
type PaginationResponse struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}
