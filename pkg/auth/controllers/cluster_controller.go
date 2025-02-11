package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"net/http"
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
