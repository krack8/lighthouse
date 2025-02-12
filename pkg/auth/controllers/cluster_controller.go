package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"net/http"
	"os"
)

type ClusterController struct {
	ClusterService *services.ClusterService
}

func NewClusterController(clusterService *services.ClusterService) *ClusterController {
	return &ClusterController{
		ClusterService: clusterService,
	}
}

// GetClusterHandler handles fetching a Cluster by ID.
func (uc *ClusterController) GetClusterHandler(c *gin.Context) {
	id := c.Param("id")

	Cluster, err := uc.ClusterService.GetCluster(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, Cluster)
}

// GetAllClustersHandler handles fetching all Clusters.
func (uc *ClusterController) GetAllClustersHandler(c *gin.Context) {
	ClusterList, err := uc.ClusterService.GetAllClusters()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, ClusterList)
}

// CreateClusterHandler handles creating a new Cluster.
func (uc *ClusterController) CreateAgentClusterHandler(c *gin.Context) {
	var request struct {
		Name            string `json:"name"`
		MasterClusterId string `json:"masterClusterId"`
	}

	// Bind the JSON payload to the request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Create the cluster
	cluster, err := uc.ClusterService.CreateAgentCluster(request.Name, request.MasterClusterId)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Prepare the response with additional information
	response := struct {
		Id            string `json:"id"`
		Name          string `json:"name"`
		Token         string `json:"token"`
		ControllerURL string `json:"controller_url"`
		SecretName    string `json:"secret_name"`
	}{
		Id:            cluster.ID.Hex(),
		Name:          cluster.Name,
		Token:         cluster.Token.TokenHash,
		ControllerURL: os.Getenv("CONTROLLER_URL"),
		SecretName:    os.Getenv("AGENT_SECRET_NAME"),
	}

	// Respond with the newly created cluster
	utils.RespondWithJSON(c, http.StatusCreated, response)
}

func (uc *ClusterController) GetClusterHelmDetailsHandler(c *gin.Context) {
	id := c.Param("id")

	Cluster, err := uc.ClusterService.GetCluster(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Prepare the response with additional information
	response := struct {
		RepoCommand string `json:"helm_repo"`
		HelmCommand string `json:"helm_command"`
	}{
		RepoCommand: "helm repo add krack8 https://krack8.github.io/charts",
		HelmCommand: "helm install lighthouse-agent --create-namespace --namespace " + os.Getenv("RESOURCE_NAMESPACE") + " krack8/lighthouse-agent --version 1.0.0 \\\n --set auth.enabled=true \\\n --set agent.enabled=true \\\n --set controller.enabled=false \\\n --set mongo.external=true \\\n --set auth.token=" + Cluster.Token.TokenHash + " \\\n --set controller.url=" + os.Getenv("CONTROLLER_URL"),
	}

	utils.RespondWithJSON(c, http.StatusOK, response)
}
