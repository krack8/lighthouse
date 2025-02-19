package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/server"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
	rbacv1 "k8s.io/api/rbac/v1"
)

type ClusterRoleBindingControllerInterface interface {
	GetClusterRoleBindingList(ctx *gin.Context)
	GetClusterRoleBindingDetails(ctx *gin.Context)
	DeployClusterRoleBinding(ctx *gin.Context)
	DeleteClusterRoleBinding(ctx *gin.Context)
}

type clusterRoleBindingController struct {
}

var crbc clusterRoleBindingController

func ClusterRoleBindingController() *clusterRoleBindingController {
	return &crbc
}

func (ctrl *clusterRoleBindingController) GetClusterRoleBindingList(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.GetClusterRoleBindingListInputParams)

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
			log.Logger.Info("Filter by param for ClusterRoleBinding List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
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
		log.Logger.Errorw("unable to marshal GetClusterRoleBindingList Task input ", "err", err.Error())
	}
	taskName := tasks.GetTaskName(k8s.ClusterRoleBindingService().GetClusterRoleBindingList)
	logRequestedTaskController("cluster-role-binding", taskName)
	res, err := server.TaskToAgent().SendToWorker(ctx, taskName, inputTask, clusterGroupName)
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

func (ctrl *clusterRoleBindingController) GetClusterRoleBindingDetails(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.GetClusterRoleBindingDetailsInputParams)
	input.ClusterRoleBindingName = ctx.Param("name")
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal GetClusterRoleBindingDetails Task input ", "err", err.Error())
	}
	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.ClusterRoleBindingService().GetClusterRoleBindingDetails)
	logRequestedTaskController("cluster-role-binding", taskName)
	res, err := server.TaskToAgent().SendToWorker(ctx, taskName, inputTask, clusterGroupName)
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

func (ctrl *clusterRoleBindingController) DeployClusterRoleBinding(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(rbacv1.ClusterRoleBinding)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy clusterRoleBinding payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}

	input := new(k8s.DeployClusterRoleBindingInputParams)
	input.ClusterRoleBinding = payload
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal DeployClusterRoleBinding Task input ", "err", err.Error())
	}
	taskName := tasks.GetTaskName(k8s.ClusterRoleBindingService().DeployClusterRoleBinding)
	logRequestedTaskController("cluster-role-binding", taskName)
	res, err := server.TaskToAgent().SendToWorker(ctx, taskName, inputTask, clusterGroupName)
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

func (ctrl *clusterRoleBindingController) DeleteClusterRoleBinding(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteClusterRoleBindingInputParams)
	input.ClusterRoleBindingName = ctx.Param("name")
	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.ClusterRoleBindingService().DeleteClusterRoleBinding)
	logRequestedTaskController("cluster-role-binding", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := server.TaskToAgent().SendToWorker(ctx, taskName, inputTask, clusterGroupName)
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
