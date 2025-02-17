package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/worker"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
)

type VolumeSnapshotContentControllerInterface interface {
	GetVolumeSnapshotContentList(ctx *gin.Context)
	GetVolumeSnapshotContentDetails(ctx *gin.Context)
}

type volumeSnapshotContentController struct {
}

var vscc volumeSnapshotContentController

func VolumeSnapshotContentController() *volumeSnapshotContentController {
	return &vscc
}

func (ctrl *volumeSnapshotContentController) GetVolumeSnapshotContentList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetVolumeSnapshotContentListInputParams)

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
			log.Logger.Info("Filter by param for VolumeSnapshotContent List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	input.Search = ctx.Query("q")
	taskName := tasks.GetTaskName(k8s.VolumeSnapshotContentService().GetVolumeSnapshotContentList)
	logRequestedTaskController("volume-snapshot-content", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := worker.TaskToAgent().SendToWorker(ctx, taskName, inputTask, clusterGroup)
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

func (ctrl *volumeSnapshotContentController) GetVolumeSnapshotContentDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetVolumeSnapshotContentDetailsInputParams)
	input.VolumeSnapshotContentName = ctx.Param("name")
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.VolumeSnapshotContentService().GetVolumeSnapshotContentDetails)
	logRequestedTaskController("volume-snapshot-content", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := worker.TaskToAgent().SendToWorker(ctx, taskName, inputTask, clusterGroup)
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
