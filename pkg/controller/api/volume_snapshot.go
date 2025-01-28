package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/worker"
	"github.com/krack8/lighthouse/pkg/dto"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
)

type VolumeSnapshotControllerInterface interface {
	GetVolumeSnapshotList(ctx *gin.Context)
	GetVolumeSnapshotDetails(ctx *gin.Context)
	DeployVolumeSnapshot(ctx *gin.Context)
	DeleteVolumeSnapshot(ctx *gin.Context)
}

type volumeSnapshotController struct {
}

var vsc volumeSnapshotController

func VolumeSnapshotController() *volumeSnapshotController {
	return &vsc
}

func (ctrl *volumeSnapshotController) GetVolumeSnapshotList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetVolumeSnapshotListInputParams)

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	input.Search = ctx.Query("q")

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
			log.Logger.Info("Filter by param for VolumeSnapshot List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	taskName := tasks.GetTaskName(k8s.VolumeSnapshotService().GetVolumeSnapshotList)
	logRequestedTaskController("volume-snapshot", taskName)
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

func (ctrl *volumeSnapshotController) GetVolumeSnapshotDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetVolumeSnapshotDetailsInputParams)
	input.VolumeSnapshotName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	taskName := tasks.GetTaskName(k8s.VolumeSnapshotService().GetVolumeSnapshotDetails)
	logRequestedTaskController("volume-snapshot", taskName)
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

func (ctrl *volumeSnapshotController) DeployVolumeSnapshot(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(dto.VolumeSnapshotV1)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy VolumeSnapshot payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	input := new(k8s.DeployVolumeSnapshotInputParams)
	input.VolumeSnapshot = payload
	if input.VolumeSnapshot.Namespace == "" {
		log.Logger.Errorw("namespace required in payload", "value", "volumeSnapshot deploy")
		SendErrorResponse(ctx, ErrNamespaceEmpty)
		return
	}
	taskName := tasks.GetTaskName(k8s.VolumeSnapshotService().DeployVolumeSnapshot)
	logRequestedTaskController("volume-snapshot", taskName)
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

func (ctrl *volumeSnapshotController) DeleteVolumeSnapshot(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteVolumeSnapshotInputParams)
	input.VolumeSnapshotName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	taskName := tasks.GetTaskName(k8s.VolumeSnapshotService().DeleteVolumeSnapshot)
	logRequestedTaskController("volume-snapshot", taskName)
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
