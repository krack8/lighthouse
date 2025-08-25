package argocd

import "time"

// Application represents an ArgoCD application
type Application struct {
	Metadata ApplicationMetadata `json:"metadata"`
	Spec     ApplicationSpec     `json:"spec"`
	Status   ApplicationStatus   `json:"status,omitempty"`
}

type ApplicationMetadata struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	UID               string            `json:"uid,omitempty"`
	ResourceVersion   string            `json:"resourceVersion,omitempty"`
	Generation        int64             `json:"generation,omitempty"`
	CreationTimestamp time.Time         `json:"creationTimestamp,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
}

type ApplicationSpec struct {
	Name        string                 `json:"-"` // Not sent in JSON
	Namespace   string                 `json:"-"` // Not sent in JSON
	Source      ApplicationSource      `json:"source"`
	Destination ApplicationDestination `json:"destination"`
	Project     string                 `json:"project"`
	SyncPolicy  *SyncPolicy            `json:"syncPolicy,omitempty"`
}

type ApplicationSource struct {
	RepoURL        string     `json:"repoURL"`
	Path           string     `json:"path,omitempty"`
	TargetRevision string     `json:"targetRevision,omitempty"`
	Helm           *Helm      `json:"helm,omitempty"`
	Kustomize      *Kustomize `json:"kustomize,omitempty"`
}

type Helm struct {
	ValueFiles []string        `json:"valueFiles,omitempty"`
	Parameters []HelmParameter `json:"parameters,omitempty"`
	Values     string          `json:"values,omitempty"`
}

type HelmParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Kustomize struct {
	NamePrefix string   `json:"namePrefix,omitempty"`
	NameSuffix string   `json:"nameSuffix,omitempty"`
	Images     []string `json:"images,omitempty"`
}

type ApplicationDestination struct {
	Server    string `json:"server,omitempty"`
	Namespace string `json:"namespace"`
	Name      string `json:"name,omitempty"`
}

type SyncPolicy struct {
	Automated   *SyncPolicyAutomated `json:"automated,omitempty"`
	SyncOptions []string             `json:"syncOptions,omitempty"`
	Retry       *RetryStrategy       `json:"retry,omitempty"`
}

type SyncPolicyAutomated struct {
	Prune      bool `json:"prune,omitempty"`
	SelfHeal   bool `json:"selfHeal,omitempty"`
	AllowEmpty bool `json:"allowEmpty,omitempty"`
}

type RetryStrategy struct {
	Limit   int64    `json:"limit,omitempty"`
	Backoff *Backoff `json:"backoff,omitempty"`
}

type Backoff struct {
	Duration    string `json:"duration,omitempty"`
	Factor      *int64 `json:"factor,omitempty"`
	MaxDuration string `json:"maxDuration,omitempty"`
}

type ApplicationStatus struct {
	Resources      []ResourceStatus       `json:"resources,omitempty"`
	Sync           SyncStatus             `json:"sync,omitempty"`
	Health         HealthStatus           `json:"health,omitempty"`
	History        []RevisionHistory      `json:"history,omitempty"`
	Conditions     []ApplicationCondition `json:"conditions,omitempty"`
	ReconciledAt   *time.Time             `json:"reconciledAt,omitempty"`
	OperationState *OperationState        `json:"operationState,omitempty"`
	SourceType     string                 `json:"sourceType,omitempty"`
	Summary        ApplicationSummary     `json:"summary,omitempty"`
}

type ResourceStatus struct {
	Group     string        `json:"group,omitempty"`
	Version   string        `json:"version,omitempty"`
	Kind      string        `json:"kind"`
	Namespace string        `json:"namespace,omitempty"`
	Name      string        `json:"name"`
	Status    string        `json:"status,omitempty"`
	Health    *HealthStatus `json:"health,omitempty"`
}

type SyncStatus struct {
	Status     string     `json:"status"`
	ComparedTo ComparedTo `json:"comparedTo,omitempty"`
	Revision   string     `json:"revision,omitempty"`
}

type ComparedTo struct {
	Source      ApplicationSource      `json:"source"`
	Destination ApplicationDestination `json:"destination"`
}

type HealthStatus struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type RevisionHistory struct {
	Revision   string    `json:"revision"`
	DeployedAt time.Time `json:"deployedAt"`
	ID         int64     `json:"id"`
}

type ApplicationCondition struct {
	Type               string     `json:"type"`
	Message            string     `json:"message"`
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
}

type OperationState struct {
	Operation  Operation            `json:"operation,omitempty"`
	Phase      string               `json:"phase,omitempty"`
	Message    string               `json:"message,omitempty"`
	SyncResult *SyncOperationResult `json:"syncResult,omitempty"`
	StartedAt  time.Time            `json:"startedAt,omitempty"`
	FinishedAt *time.Time           `json:"finishedAt,omitempty"`
}

type Operation struct {
	Sync *SyncOperation `json:"sync,omitempty"`
}

type SyncOperation struct {
	Revision     string                  `json:"revision,omitempty"`
	Prune        bool                    `json:"prune,omitempty"`
	DryRun       bool                    `json:"dryRun,omitempty"`
	SyncStrategy *SyncStrategy           `json:"syncStrategy,omitempty"`
	Resources    []SyncOperationResource `json:"resources,omitempty"`
}

type SyncStrategy struct {
	Apply *SyncStrategyApply `json:"apply,omitempty"`
	Hook  *SyncStrategyHook  `json:"hook,omitempty"`
}

type SyncStrategyApply struct {
	Force bool `json:"force,omitempty"`
}

type SyncStrategyHook struct {
	Force bool `json:"force,omitempty"`
}

type SyncOperationResource struct {
	Group     string `json:"group,omitempty"`
	Version   string `json:"version,omitempty"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name"`
}

type SyncOperationResult struct {
	Resources []ResourceResult  `json:"resources,omitempty"`
	Revision  string            `json:"revision"`
	Source    ApplicationSource `json:"source,omitempty"`
}

type ResourceResult struct {
	Group     string `json:"group,omitempty"`
	Version   string `json:"version,omitempty"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name"`
	Status    string `json:"status,omitempty"`
	Message   string `json:"message,omitempty"`
	HookPhase string `json:"hookPhase,omitempty"`
	HookType  string `json:"hookType,omitempty"`
	SyncPhase string `json:"syncPhase,omitempty"`
}

type ApplicationSummary struct {
	ExternalURLs []string `json:"externalURLs,omitempty"`
	Images       []string `json:"images,omitempty"`
}

// List types
type ApplicationList struct {
	Items []Application `json:"items"`
}

type Repository struct {
	Repo            string          `json:"repo"`
	Username        string          `json:"username,omitempty"`
	Password        string          `json:"password,omitempty"`
	SSHPrivateKey   string          `json:"sshPrivateKey,omitempty"`
	ConnectionState ConnectionState `json:"connectionState,omitempty"`
	Type            string          `json:"type,omitempty"`
	Name            string          `json:"name,omitempty"`
	Project         string          `json:"project,omitempty"`
}

type ConnectionState struct {
	Status      string    `json:"status"`
	Message     string    `json:"message"`
	AttemptedAt time.Time `json:"attemptedAt"`
}

type RepositoryList struct {
	Items []Repository `json:"items"`
}

type Project struct {
	Metadata ProjectMetadata `json:"metadata"`
	Spec     ProjectSpec     `json:"spec"`
}

type ProjectMetadata struct {
	Name string `json:"name"`
}

type ProjectSpec struct {
	Name         string                   `json:"-"` // Not sent in JSON
	Description  string                   `json:"description,omitempty"`
	SourceRepos  []string                 `json:"sourceRepos,omitempty"`
	Destinations []ApplicationDestination `json:"destinations,omitempty"`
}

type GroupKind struct {
	Group string `json:"group,omitempty"`
	Kind  string `json:"kind"`
}

type ProjectRole struct {
	Name     string   `json:"name"`
	Policies []string `json:"policies,omitempty"`
	Groups   []string `json:"groups,omitempty"`
}

type ProjectList struct {
	Items []Project `json:"items"`
}
