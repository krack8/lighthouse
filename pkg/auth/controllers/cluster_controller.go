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
		Namespace       string `json:"namespace"`
		MasterClusterId string `json:"masterClusterId"`
	}

	// Bind the JSON payload to the request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Create the cluster
	cluster, err := uc.ClusterService.CreateAgentCluster(request.Name, request.Namespace, request.MasterClusterId)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Prepare the response with additional information
	response := struct {
		Id            string `json:"id"`
		Name          string `json:"name"`
		Namespace     string `json:"namespace"`
		Token         string `json:"token"`
		ControllerURL string `json:"controller_url"`
		SecretName    string `json:"secret_name"`
	}{
		Id:            cluster.ID.Hex(),
		Name:          cluster.Name,
		Namespace:     request.Namespace,
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
		Id            string `json:"id"`
		Name          string `json:"name"`
		Namespace     string `json:"namespace"`
		Token         string `json:"token"`
		ControllerURL string `json:"controller_url"`
		SecretName    string `json:"secret_name"`
	}{
		Id:            Cluster.ID.Hex(),
		Name:          Cluster.Name,
		Namespace:     Cluster.SecretNamespace,
		Token:         Cluster.Token.TokenHash,
		ControllerURL: os.Getenv("CONTROLLER_URL"),
		SecretName:    os.Getenv("AGENT_SECRET_NAME"),
	}

	utils.RespondWithJSON(c, http.StatusOK, response)
}
