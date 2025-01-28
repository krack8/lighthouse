package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/worker"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
	corev1 "k8s.io/api/core/v1"
)

type ReplicationControllerControllerInterface interface {
	GetReplicationControllerList(ctx *gin.Context)
	GetReplicationControllerDetails(ctx *gin.Context)
	DeployReplicationController(ctx *gin.Context)
	DeleteReplicationController(ctx *gin.Context)
}

type replicationControllerController struct {
}

var rcc replicationControllerController

func ReplicationControllerController() *replicationControllerController {
	return &rcc
}

func (ctrl *replicationControllerController) GetReplicationControllerList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetReplicationControllerListInputParams)

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
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

		err := json.Unmarshal([]byte(jsonLabel), &queryLabelMap)
		if err != nil {
			log.Logger.Error("query labels unmarshal error ", "err", err.Error())
		}
		if queryLabelMap != nil {
			input.Labels = queryLabelMap
			log.Logger.Info("Filter by param for ReplicationController List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	taskName := tasks.GetTaskName(k8s.ReplicationControllerService().GetReplicationControllerList)
	logRequestedTaskController("replication-controller", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := worker.TaskToAgent().SendToWorker(ctx, taskName, inputTask)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
	}
	SendResponse(ctx, result)
}

func (ctrl *replicationControllerController) GetReplicationControllerDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetReplicationControllerDetailsInputParams)
	input.ReplicationControllerName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	taskName := tasks.GetTaskName(k8s.ReplicationControllerService().GetReplicationControllerDetails)
	logRequestedTaskController("replication-controller", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := worker.TaskToAgent().SendToWorker(ctx, taskName, inputTask)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
	}
	SendResponse(ctx, result)
}

func (ctrl *replicationControllerController) DeployReplicationController(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(corev1.ReplicationController)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy configmap payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	input := new(k8s.DeployReplicationControllerInputParams)
	input.ReplicationController = payload
	if input.ReplicationController.Namespace == "" {
		log.Logger.Errorw("namespace required in payload", "value", "config-map deploy")
		SendErrorResponse(ctx, ErrNamespaceEmpty)
		return
	}
	taskName := tasks.GetTaskName(k8s.ReplicationControllerService().DeployReplicationController)
	logRequestedTaskController("replication-controller", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := worker.TaskToAgent().SendToWorker(ctx, taskName, inputTask)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
	}
	SendResponse(ctx, result)
}

func (ctrl *replicationControllerController) DeleteReplicationController(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteReplicationControllerInputParams)
	input.ReplicationControllerName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	taskName := tasks.GetTaskName(k8s.ReplicationControllerService().DeleteReplicationController)
	logRequestedTaskController("replication-controller", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := worker.TaskToAgent().SendToWorker(ctx, taskName, inputTask)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
	}
	SendResponse(ctx, result)
}
