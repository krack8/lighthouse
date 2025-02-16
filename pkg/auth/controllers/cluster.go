package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"net/http"
	"os"
	"strings"
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
		Name          string `json:"name"`
		ControllerURl string `json:"controller_url"`
	}

	// Bind the JSON payload to the request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Create the cluster
	cluster, err := uc.ClusterService.CreateAgentCluster(request.Name, request.ControllerURl)
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
		Token:         cluster.Token.RawTokenHash,
		ControllerURL: os.Getenv("CONTROLLER_URL"),
		SecretName:    os.Getenv("AGENT_SECRET_NAME"),
	}

	// Respond with the newly created cluster
	utils.RespondWithJSON(c, http.StatusCreated, response)
}

// DeleteRoleHandler handles the deletion of a role by its ID
func (uc *ClusterController) DeleteClusterHandler(c *gin.Context) {
	roleID := c.Param("id")
	if roleID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cluster ID is required"})
		return
	}

	// Call the service to delete the role
	err := uc.ClusterService.DeleteClusterByID(roleID)
	if err != nil {
		if strings.Contains(err.Error(), "no cluster found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting cluster"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Cluster %s deleted successfully", roleID)})
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
		RepoCommand: "helm repo add krack8 https://krack8.github.io/helm-charts",
		HelmCommand: "helm install lighthouse --create-namespace --namespace " + os.Getenv("RESOURCE_NAMESPACE") + " krack8/lighthouse \\\n --set auth.enabled=true \\\n --set agent.enabled=true \\\n --set auth.token=" + Cluster.Token.CombinedToken + " \\\n --set controller.url=" + os.Getenv("CONTROLLER_URL"),
	}

	utils.RespondWithJSON(c, http.StatusOK, response)
}
