package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/agent/tasks"
	"github.com/krack8/lighthouse/pkg/common/k8s"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/controller/core"
)

type VolumeSnapshotClassControllerInterface interface {
	GetVolumeSnapshotClassList(ctx *gin.Context)
	GetVolumeSnapshotClassDetails(ctx *gin.Context)
}

type volumeSnapshotClassController struct {
}

var vsclc volumeSnapshotClassController

func VolumeSnapshotClassController() *volumeSnapshotClassController {
	return &vsclc
}

func (ctrl *volumeSnapshotClassController) GetVolumeSnapshotClassList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetVolumeSnapshotClassListInputParams)

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
			log.Logger.Info("Filter by param for VolumeSnapshotClass List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	input.Search = ctx.Query("q")
	taskName := tasks.GetTaskName(k8s.VolumeSnapshotClassService().GetVolumeSnapshotClassList)
	logRequestedTaskController("volume-snapshot-class", taskName)
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

func (ctrl *volumeSnapshotClassController) GetVolumeSnapshotClassDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetVolumeSnapshotClassDetailsInputParams)
	input.VolumeSnapshotClassName = ctx.Param("name")
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.VolumeSnapshotClassService().GetVolumeSnapshotClassDetails)
	logRequestedTaskController("volume-snapshot-class", taskName)
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
