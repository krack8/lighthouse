package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/agent/tasks"
	"github.com/krack8/lighthouse/pkg/common/dto"
	"github.com/krack8/lighthouse/pkg/common/k8s"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/controller/core"
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
	} else if kind == "Node" {
		log.Logger.Errorw("kind node is not allowed", "value", "manifest")
		SendErrorResponse(ctx, "kind node is not allowed")
		return
	}
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
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
