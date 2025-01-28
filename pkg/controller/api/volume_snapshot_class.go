package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/worker"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
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

		err := json.Unmarshal([]byte(jsonLabel), &queryLabelMap)
		if err != nil {
			log.Logger.Error("query labels unmarshal error ", "err", err.Error())
		}
		if queryLabelMap != nil {
			input.Labels = queryLabelMap
			log.Logger.Info("Filter by param for VolumeSnapshotClass List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	input.Search = ctx.Query("q")
	taskName := tasks.GetTaskName(k8s.VolumeSnapshotClassService().GetVolumeSnapshotClassList)
	logRequestedTaskController("volume-snapshot-class", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := worker.TaskToAgent().SendToWorker(ctx, taskName, inputTask)
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
	taskName := tasks.GetTaskName(k8s.VolumeSnapshotClassService().GetVolumeSnapshotClassDetails)
	logRequestedTaskController("volume-snapshot-class", taskName)
	inputTask, err := json.Marshal(input)
	if err != nil {
		logErrMarshalTaskController(taskName, err)
	}
	res, err := worker.TaskToAgent().SendToWorker(ctx, taskName, inputTask)
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
