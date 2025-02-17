package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/worker"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
	rbacv1 "k8s.io/api/rbac/v1"
)

type ClusterRoleControllerInterface interface {
	GetClusterRoleList(ctx *gin.Context)
	GetClusterRoleDetails(ctx *gin.Context)
	DeployClusterRole(ctx *gin.Context)
	DeleteClusterRole(ctx *gin.Context)
}

type clusterRoleController struct {
}

var crc clusterRoleController

func ClusterRoleController() *clusterRoleController {
	return &crc
}

func (ctrl *clusterRoleController) GetClusterRoleList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetClusterRoleListInputParams)

	queryLabel := ctx.Query("labels")
	if queryLabel != "" {
		jsonLabel := []byte(queryLabel)
		queryLabelMap := map[string]string{}

		err := json.Unmarshal(jsonLabel, &queryLabelMap)
		if err != nil {
			log.Logger.Error("query labels unmarshal error ", "err", err.Error())
		}
		if queryLabelMap != nil {
			input.Labels = queryLabelMap
			log.Logger.Info("Filter by param for ClusterRole List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	input.Search = ctx.Query("q")
	input.Continue = ctx.Query("continue")
	input.Limit = ctx.Query("limit")
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal GetClusterRoleList Task input ", "err", err.Error())
	}
	taskName := tasks.GetTaskName(k8s.ClusterRoleService().GetClusterRoleList)
	logRequestedTaskController("cluster-role", taskName)
	res, err := worker.TaskToAgent().SendToWorker(ctx, taskName, inputTask, clusterGroupName)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	SendResponse(ctx, result)
}

func (ctrl *clusterRoleController) GetClusterRoleDetails(ctx *gin.Context) {
	var result ResponseDTO
	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}

	input := new(k8s.GetClusterRoleDetailsInputParams)
	input.ClusterRoleName = ctx.Param("name")
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal GetClusterRoleDetails Task input ", "err", err.Error())
	}
	taskName := tasks.GetTaskName(k8s.ClusterRoleService().GetClusterRoleDetails)
	logRequestedTaskController("cluster-role", taskName)
	res, err := worker.TaskToAgent().SendToWorker(ctx, taskName, inputTask, clusterGroupName)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	SendResponse(ctx, result)
}

func (ctrl *clusterRoleController) DeployClusterRole(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(rbacv1.ClusterRole)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy clusterRole payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}

	input := new(k8s.DeployClusterRoleInputParams)
	input.ClusterRole = payload
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal DeployClusterRole Task input ", "err", err.Error())
	}
	taskName := tasks.GetTaskName(k8s.ClusterRoleService().DeployClusterRole)
	logRequestedTaskController("cluster-role", taskName)
	res, err := worker.TaskToAgent().SendToWorker(ctx, taskName, inputTask, clusterGroupName)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	SendResponse(ctx, result)
}

func (ctrl *clusterRoleController) DeleteClusterRole(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteClusterRoleInputParams)
	input.ClusterRoleName = ctx.Param("name")
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal DeleteClusterRole Task input ", "err", err.Error())
	}
	taskName := tasks.GetTaskName(k8s.ClusterRoleService().DeleteClusterRole)
	logRequestedTaskController("cluster-role", taskName)

	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}

	res, err := worker.TaskToAgent().SendToWorker(ctx, taskName, inputTask, clusterGroupName)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	SendResponse(ctx, result)
}
