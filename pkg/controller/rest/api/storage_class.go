package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/agent/tasks"
	"github.com/krack8/lighthouse/pkg/common/k8s"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/controller/core"
	storagev1 "k8s.io/api/storage/v1"
)

type StorageClassControllerInterface interface {
	GetStorageClassList(ctx *gin.Context)
	GetStorageClassDetails(ctx *gin.Context)
	DeployStorageClass(ctx *gin.Context)
	DeleteStorageClass(ctx *gin.Context)
}

type storageClassController struct {
}

var scc storageClassController

func StorageClassController() *storageClassController {
	return &scc
}

func (ctrl *storageClassController) GetStorageClassList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetStorageClassListInputParams)

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
			log.Logger.Info("Filter by param for StorageClass List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	input.Search = ctx.Query("q")
	input.Continue = ctx.Query("continue")
	input.Limit = ctx.Query("limit")
	taskName := tasks.GetTaskName(k8s.StorageClassService().GetStorageClassList)
	logRequestedTaskController("storage-class", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroup)
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

func (ctrl *storageClassController) GetStorageClassDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetStorageClassDetailsInputParams)
	input.StorageClassName = ctx.Param("name")
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.StorageClassService().GetStorageClassDetails)
	logRequestedTaskController("storage-class", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroup)
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

func (ctrl *storageClassController) DeployStorageClass(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(storagev1.StorageClass)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy StorageClass payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	input := new(k8s.DeployStorageClassInputParams)
	input.StorageClass = payload
	taskName := tasks.GetTaskName(k8s.StorageClassService().DeployStorageClass)
	logRequestedTaskController("storage-class", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroup)
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

func (ctrl *storageClassController) DeleteStorageClass(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteStorageClassInputParams)
	input.StorageClassName = ctx.Param("name")
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.StorageClassService().DeleteStorageClass)
	logRequestedTaskController("storage-class", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroup)
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
