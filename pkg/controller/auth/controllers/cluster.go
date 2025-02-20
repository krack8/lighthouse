package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/controller/auth/services"
	"github.com/krack8/lighthouse/pkg/controller/auth/utils"
	"github.com/krack8/lighthouse/pkg/controller/core"
	"net/http"
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
		utils.RespondWithError(c, http.StatusOK, "Cluster not found")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, ClusterList)
}

// CreateClusterHandler handles creating a new Cluster.
func (uc *ClusterController) CreateAgentClusterHandler(c *gin.Context) {
	var request struct {
		Name          string `json:"name"`
		ControllerURl string `json:"grpc_url"`
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
		ControllerURL string `json:"grpc_url"`
		SecretName    string `json:"secret_name"`
	}{
		Id:            cluster.ID.Hex(),
		Name:          cluster.Name,
		Token:         cluster.Token.RawTokenHash,
		ControllerURL: config.GrpcServer,
		SecretName:    config.AgentSecretName,
	}

	// Respond with the newly created cluster
	utils.RespondWithJSON(c, http.StatusCreated, response)
}

// DeleteRoleHandler handles the deletion of a role by its ID
func (uc *ClusterController) DeleteClusterHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cluster ID is required"})
		return
	}

	cluster, err := uc.ClusterService.GetCluster(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if removed := core.GetAgentManager().RemoveAgentByGroupName(cluster.AgentGroup); !removed {
		_ = fmt.Errorf("failed to remove agent group: %s", cluster.AgentGroup)
	}
	// Call the service to delete the role
	err = uc.ClusterService.DeleteClusterByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "no cluster found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting cluster"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Cluster %s deleted successfully", id)})
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
		HelmCommand: "helm install lighthouse --create-namespace --namespace " + config.ResourceNamespace + " krack8/lighthouse \\\n --set agent.enabled=true \\\n --set config.controller.grpc.tls.enabled=true \\\n --set config.controller.grpc.tls.skipVerification=false  \\\n --set agent.group=" + Cluster.AgentGroup + " \\\n --set auth.token=" + Cluster.Token.CombinedToken + " \\\n --set config.controller.grpc.host=" + config.GrpcServer,
	}

	utils.RespondWithJSON(c, http.StatusOK, response)
}
