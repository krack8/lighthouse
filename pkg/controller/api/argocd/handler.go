package argocd

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/core"
)

type Handler struct {
	// Uses core agent manager through SendTaskToAgent
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ListApplications(c *gin.Context) {
	agentGroup := c.Query("agent_group")
	project := c.Query("project")

	if agentGroup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "agent_group is required"})
		return
	}

	// Prepare input data
	input := map[string]interface{}{
		"project": project,
	}
	inputJSON, _ := json.Marshal(input)

	// Send task using the existing SendTaskToAgent function
	result, err := core.GetAgentManager().SendTaskToAgent(
		c.Request.Context(),
		"argocd:list_applications", // task name
		inputJSON,                  // input data
		agentGroup,                 // agent group
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Parse result
	var apps interface{}
	if err := json.Unmarshal([]byte(result.Output), &apps); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, apps)
}

func (h *Handler) GetApplication(c *gin.Context) {
	name := c.Param("name")
	agentGroup := c.Query("agent_group")

	if agentGroup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "agent_group is required"})
		return
	}

	input := map[string]interface{}{
		"name": name,
	}
	inputJSON, _ := json.Marshal(input)

	result, err := core.GetAgentManager().SendTaskToAgent(
		c.Request.Context(),
		"argocd:get_application",
		inputJSON,
		agentGroup,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var app interface{}
	if err := json.Unmarshal([]byte(result.Output), &app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, app)
}

func (h *Handler) CreateApplication(c *gin.Context) {
	agentGroup := c.Query("agent_group")
	if agentGroup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "agent_group is required"})
		return
	}

	var req CreateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert request to input data
	inputJSON, _ := json.Marshal(req)

	result, err := core.GetAgentManager().SendTaskToAgent(
		c.Request.Context(),
		"argocd:create_application",
		inputJSON,
		agentGroup,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var app interface{}
	if err := json.Unmarshal([]byte(result.Output), &app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	c.JSON(http.StatusCreated, app)
}

func (h *Handler) UpdateApplication(c *gin.Context) {
	name := c.Param("name")
	agentGroup := c.Query("agent_group")

	if agentGroup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "agent_group is required"})
		return
	}

	var req UpdateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := map[string]interface{}{
		"name":        name,
		"application": req,
	}
	inputJSON, _ := json.Marshal(input)

	result, err := core.GetAgentManager().SendTaskToAgent(
		c.Request.Context(),
		"argocd:update_application",
		inputJSON,
		agentGroup,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var app interface{}
	if err := json.Unmarshal([]byte(result.Output), &app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, app)
}

func (h *Handler) RollbackApplication(c *gin.Context) {
	name := c.Param("name")
	agentGroup := c.Query("agent_group")

	if agentGroup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "agent_group is required"})
		return
	}

	var req struct {
		Revision string `json:"revision" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := map[string]interface{}{
		"name":     name,
		"revision": req.Revision,
	}
	inputJSON, _ := json.Marshal(input)

	result, err := core.GetAgentManager().SendTaskToAgent(
		c.Request.Context(),
		"argocd:rollback_application",
		inputJSON,
		agentGroup,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var app interface{}
	if err := json.Unmarshal([]byte(result.Output), &app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, app)
}

func (h *Handler) ListProjects(c *gin.Context) {
	agentGroup := c.Query("agent_group")

	if agentGroup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "agent_group is required"})
		return
	}

	result, err := core.GetAgentManager().SendTaskToAgent(
		c.Request.Context(),
		"argocd:list_projects",
		nil, // no input needed
		agentGroup,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var projects interface{}
	if err := json.Unmarshal([]byte(result.Output), &projects); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

func (h *Handler) CreateProject(c *gin.Context) {
	agentGroup := c.Query("agent_group")
	if agentGroup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "agent_group is required"})
		return
	}

	var req CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inputJSON, _ := json.Marshal(req)

	result, err := core.GetAgentManager().SendTaskToAgent(
		c.Request.Context(),
		"argocd:create_project",
		inputJSON,
		agentGroup,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var project interface{}
	if err := json.Unmarshal([]byte(result.Output), &project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func (h *Handler) ListRepositories(c *gin.Context) {
	agentGroup := c.Query("agent_group")

	if agentGroup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "agent_group is required"})
		return
	}

	result, err := core.GetAgentManager().SendTaskToAgent(
		c.Request.Context(),
		"argocd:list_repositories",
		nil,
		agentGroup,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var repos interface{}
	if err := json.Unmarshal([]byte(result.Output), &repos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"repositories": repos})
}

func (h *Handler) CreateRepository(c *gin.Context) {
	agentGroup := c.Query("agent_group")
	if agentGroup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "agent_group is required"})
		return
	}

	var req CreateRepositoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inputJSON, _ := json.Marshal(req)

	result, err := core.GetAgentManager().SendTaskToAgent(
		c.Request.Context(),
		"argocd:create_repository",
		inputJSON,
		agentGroup,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var repo interface{}
	if err := json.Unmarshal([]byte(result.Output), &repo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	c.JSON(http.StatusCreated, repo)
}

// Stub handlers for remaining routes
func (h *Handler) RegisterAgent(c *gin.Context) {
	// Not needed for gRPC-based agents
	c.JSON(http.StatusOK, gin.H{"message": "Agent registration through gRPC"})
}

func (h *Handler) Heartbeat(c *gin.Context) {
	// Not needed for gRPC-based agents
	c.JSON(http.StatusOK, gin.H{"message": "Heartbeat through gRPC"})
}

func (h *Handler) ListAgents(c *gin.Context) {
	// Return empty array properly
	agents := make([]interface{}, 0)
	c.JSON(http.StatusOK, gin.H{"agents": agents})
}

// Example of updated handler using task sender functions (modify the above functions)
func (h *Handler) DeleteApplication(c *gin.Context) {
	name := c.Param("name")
	agentGroup := c.Query("agent_group")
	cascade := c.Query("cascade") == "true"

	if agentGroup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "agent_group is required"})
		return
	}

	// Use the task sender function
	result, err := SendDeleteApplicationTask(c.Request.Context(), agentGroup, name, cascade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": result.Success,
		"message": "Application deleted",
	})
}

func (h *Handler) SyncApplication(c *gin.Context) {
	name := c.Param("name")
	agentGroup := c.Query("agent_group")

	if agentGroup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "agent_group is required"})
		return
	}

	var req struct {
		Revision string `json:"revision,omitempty"`
		Prune    bool   `json:"prune"`
		DryRun   bool   `json:"dry_run"`
	}
	c.ShouldBindJSON(&req)

	if req.Revision == "" {
		req.Revision = "HEAD"
	}

	// Use the task sender function
	result, err := SendSyncApplicationTask(
		c.Request.Context(),
		agentGroup,
		name,
		req.Revision,
		req.Prune,
		req.DryRun,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var app interface{}
	if err := json.Unmarshal([]byte(result.Output), &app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, app)
}
