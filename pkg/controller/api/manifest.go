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

type ManifestControllerInterface interface {
	DeployManifest(ctx *gin.Context)
}

type manifestController struct {
}

var manc manifestController

func ManifestController() *manifestController {
	return &manc
}

func (ctrl *manifestController) DeployManifest(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(dto.ManifestDto)
	// plural resource
	resource := ctx.Query("resource")
	if resource == "" {
		log.Logger.Errorw("resource required in query param", "value", "manifest")
		SendErrorResponse(ctx, "resource required in query")
		return
	}
	// camel case kind . ex VolumeSnapshot
	kind := ctx.Query("kind")
	if kind == "" {
		log.Logger.Errorw("kind required in query param", "value", "manifest")
		SendErrorResponse(ctx, "kind required in query")
		return
	}
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy Manifest payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	input := new(k8s.DeployManifestInputParams)
	input.Manifest = payload
	input.Manifest.Kind = kind
	input.Resource = resource
	taskName := tasks.GetTaskName(k8s.ManifestService().DeployManifest)
	logRequestedTaskController("load-balancer", taskName)
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
