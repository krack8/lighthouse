package argocd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/shared"
)

// Handler handles ArgoCD API requests
type Handler struct {
	agentManager *AgentManager
}

// NewHandler creates a new handler
func NewHandler(am *AgentManager) *Handler {
	return &Handler{
		agentManager: am,
	}
}

// RegisterAgent handles agent registration
func (h *Handler) RegisterAgent(c *gin.Context) {
	var req shared.AgentRegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.agentManager.RegisterAgent(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shared.AgentRegistrationResponse{
		Success:           true,
		Message:           "Agent registered successfully",
		HeartbeatInterval: 30,
	})
}

// Heartbeat handles agent heartbeat
func (h *Handler) Heartbeat(c *gin.Context) {
	var req shared.HeartbeatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.agentManager.UpdateHeartbeat(req.AgentID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shared.HeartbeatResponse{
		Success: true,
	})
}

// ListAgents returns all registered agents
func (h *Handler) ListAgents(c *gin.Context) {
	agents := h.agentManager.ListAgents()
	c.JSON(http.StatusOK, gin.H{"agents": agents})
}

// ListApplications lists ArgoCD applications
func (h *Handler) ListApplications(c *gin.Context) {
	clusterID := c.Query("cluster_id")
	project := c.Query("project")

	if clusterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cluster_id is required"})
		return
	}

	agent, err := h.agentManager.GetAgent(clusterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp shared.ListApplicationsResponse
	req := shared.ListApplicationsRequest{Project: project}

	if err := agent.CallAgent("POST", "/api/argocd/applications/list", req, &resp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetApplication gets a specific application
func (h *Handler) GetApplication(c *gin.Context) {
	name := c.Param("name")
	clusterID := c.Query("cluster_id")

	if clusterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cluster_id is required"})
		return
	}

	agent, err := h.agentManager.GetAgent(clusterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var app shared.Application
	endpoint := "/api/argocd/applications/" + name

	if err := agent.CallAgent("GET", endpoint, nil, &app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

// CreateApplication creates a new application
func (h *Handler) CreateApplication(c *gin.Context) {
	var req CreateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.ClusterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cluster_id is required"})
		return
	}

	agent, err := h.agentManager.GetAgent(req.ClusterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Create shared request - Source and Destination are already shared types
	sharedReq := shared.CreateApplicationRequest{
		Name:        req.Name,
		Namespace:   req.Namespace,
		Project:     req.Project,
		Source:      req.Source,      // Already shared.ApplicationSource
		Destination: req.Destination, // Already shared.ApplicationDestination
	}

	var app shared.Application
	if err := agent.CallAgent("POST", "/api/argocd/applications", sharedReq, &app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, app)
}

// UpdateApplication updates an existing application
func (h *Handler) UpdateApplication(c *gin.Context) {
	name := c.Param("name")

	var req UpdateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.ClusterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cluster_id is required"})
		return
	}

	agent, err := h.agentManager.GetAgent(req.ClusterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Create shared request
	sharedReq := shared.UpdateApplicationRequest{
		Namespace:   req.Namespace,
		Project:     req.Project,
		Source:      req.Source,      // Already shared.ApplicationSource
		Destination: req.Destination, // Already shared.ApplicationDestination
	}

	endpoint := "/api/argocd/applications/" + name

	var app shared.Application
	if err := agent.CallAgent("PUT", endpoint, sharedReq, &app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

// DeleteApplication deletes an application
func (h *Handler) DeleteApplication(c *gin.Context) {
	name := c.Param("name")
	clusterID := c.Query("cluster_id")
	cascade := c.Query("cascade") == "true"

	if clusterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cluster_id is required"})
		return
	}

	agent, err := h.agentManager.GetAgent(clusterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	req := shared.DeleteApplicationRequest{
		Name:    name,
		Cascade: cascade,
	}

	endpoint := "/api/argocd/applications/" + name

	var resp shared.DeleteApplicationResponse
	if err := agent.CallAgent("DELETE", endpoint, req, &resp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SyncApplication syncs an application
func (h *Handler) SyncApplication(c *gin.Context) {
	name := c.Param("name")

	var req SyncApplicationRequest
	c.ShouldBindJSON(&req)

	if req.ClusterID == "" {
		req.ClusterID = c.Query("cluster_id")
	}

	if req.ClusterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cluster_id is required"})
		return
	}

	agent, err := h.agentManager.GetAgent(req.ClusterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	sharedReq := shared.SyncApplicationRequest{
		Name:   name,
		Prune:  req.Prune,
		DryRun: req.DryRun,
	}

	var resp shared.SyncApplicationResponse
	endpoint := "/api/argocd/applications/" + name + "/sync"

	if err := agent.CallAgent("POST", endpoint, sharedReq, &resp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ListProjects lists ArgoCD projects
func (h *Handler) ListProjects(c *gin.Context) {
	clusterID := c.Query("cluster_id")

	if clusterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cluster_id is required"})
		return
	}

	agent, err := h.agentManager.GetAgent(clusterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var projects []shared.Project
	if err := agent.CallAgent("GET", "/api/argocd/projects", nil, &projects); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// CreateProject creates a new project
func (h *Handler) CreateProject(c *gin.Context) {
	var req CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.ClusterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cluster_id is required"})
		return
	}

	agent, err := h.agentManager.GetAgent(req.ClusterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Create project directly - already using shared types
	project := shared.Project{
		Name:        req.Name,
		Description: req.Description,
		SourceRepos: req.SourceRepos,
	}

	var created shared.Project
	if err := agent.CallAgent("POST", "/api/argocd/projects", project, &created); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

// ListRepositories lists ArgoCD repositories
func (h *Handler) ListRepositories(c *gin.Context) {
	clusterID := c.Query("cluster_id")

	if clusterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cluster_id is required"})
		return
	}

	agent, err := h.agentManager.GetAgent(clusterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var repos []shared.Repository
	if err := agent.CallAgent("GET", "/api/argocd/repositories", nil, &repos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"repositories": repos})
}

// CreateRepository creates a new repository
func (h *Handler) CreateRepository(c *gin.Context) {
	var req CreateRepositoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.ClusterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cluster_id is required"})
		return
	}

	agent, err := h.agentManager.GetAgent(req.ClusterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Create repository - note Password field handling
	repo := shared.Repository{
		URL:      req.URL,
		Type:     req.Type,
		Username: req.Username,
		Password: req.Password, // Password is in shared.Repository
	}

	var created shared.Repository
	if err := agent.CallAgent("POST", "/api/argocd/repositories", repo, &created); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}
