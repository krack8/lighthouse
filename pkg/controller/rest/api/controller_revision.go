package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/agent/tasks"
	"github.com/krack8/lighthouse/pkg/common/k8s"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/controller/core"
	appsv1 "k8s.io/api/apps/v1"
)

type ControllerRevisionControllerInterface interface {
	GetControllerRevisionList(ctx *gin.Context)
	GetControllerRevisionDetails(ctx *gin.Context)
	DeployControllerRevision(ctx *gin.Context)
	DeleteControllerRevision(ctx *gin.Context)
}

type controllerRevisionController struct {
}

var ctrlrc controllerRevisionController

func ControllerRevisionController() *controllerRevisionController {
	return &ctrlrc
}

func (ctrl *controllerRevisionController) GetControllerRevisionList(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.GetControllerRevisionListInputParams)

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	input.Search = ctx.Query("q")
	input.Continue = ctx.Query("continue")
	input.Limit = ctx.Query("limit")

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
			log.Logger.Info("Filter by param for ControllerRevision List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	taskName := tasks.GetTaskName(k8s.ControllerRevisionService().GetControllerRevisionList)
	logRequestedTaskController("controller-revision", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroupName)
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

func (ctrl *controllerRevisionController) GetControllerRevisionDetails(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.GetControllerRevisionDetailsInputParams)
	input.ControllerRevisionName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}

	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}

	input.NamespaceName = queryNamespace
	taskName := tasks.GetTaskName(k8s.ControllerRevisionService().GetControllerRevisionDetails)
	logRequestedTaskController("controller-revision", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroupName)
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

func (ctrl *controllerRevisionController) DeployControllerRevision(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(appsv1.ControllerRevision)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy configmap payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	input := new(k8s.DeployControllerRevisionInputParams)
	input.ControllerRevision = payload
	if input.ControllerRevision.Namespace == "" {
		log.Logger.Errorw("namespace required in payload", "value", "config-map deploy")
		SendErrorResponse(ctx, ErrNamespaceEmpty)
		return
	}

	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.ControllerRevisionService().DeployControllerRevision)
	logRequestedTaskController("controller-revision", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroupName)
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

func (ctrl *controllerRevisionController) DeleteControllerRevision(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteControllerRevisionInputParams)
	input.ControllerRevisionName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	taskName := tasks.GetTaskName(k8s.ControllerRevisionService().DeleteControllerRevision)
	logRequestedTaskController("controller-revision", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroupName)
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
