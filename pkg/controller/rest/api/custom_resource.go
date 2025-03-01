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

type CustomResourceControllerInterface interface {
	GetCustomResourceList(ctx *gin.Context)
	GetCustomResourceDetails(ctx *gin.Context)
	DeployCustomResource(ctx *gin.Context)
	DeleteCustomResource(ctx *gin.Context)
}

type customResourceController struct{}

var crec customResourceController

func CustomResourceController() *customResourceController {
	return &crec
}

func (ctrl *customResourceController) GetCustomResourceList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetCustomResourceListInputParams)

	input.NamespaceName = ctx.Query("namespace")
	input.CustomResourceSGVR.Resource = ctx.Query("resource")
	input.CustomResourceSGVR.Group = ctx.Query("group")
	input.CustomResourceSGVR.Version = ctx.Query("version")
	input.Search = ctx.Query("q")
	input.Continue = ctx.Query("continue")
	input.Limit = ctx.Query("limit")

	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}

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
			log.Logger.Info("Filter by param for CustomResource List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	input.Search = ctx.Query("q")
	taskName := tasks.GetTaskName(k8s.CustomResourceService().GetCustomResourceList)
	logRequestedTaskController("custom-resource", taskName)
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

func (ctrl *customResourceController) GetCustomResourceDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetCustomResourceDetailsInputParams)

	input.Name = ctx.Param("name")
	input.NamespaceName = ctx.Query("namespace")
	input.CustomResourceSGVR.Resource = ctx.Query("resource")
	input.CustomResourceSGVR.Group = ctx.Query("group")
	input.CustomResourceSGVR.Version = ctx.Query("version")

	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}

	taskName := tasks.GetTaskName(k8s.CustomResourceService().GetCustomResourceDetails)
	logRequestedTaskController("custom-resource", taskName)
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

func (ctrl *customResourceController) DeployCustomResource(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(dto.CustomResource)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy CustomResource payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	input := new(k8s.DeployCustomResourceInputParams)
	input.CustomResource = payload
	input.Kind = ctx.Query("kind")
	input.NamespaceName = ctx.Query("namespace")
	input.CustomResourceSGVR.Resource = ctx.Query("resource")
	input.CustomResourceSGVR.Version = ctx.Query("version")
	input.CustomResourceSGVR.Group = ctx.Query("group")

	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}

	taskName := tasks.GetTaskName(k8s.CustomResourceService().DeployCustomResource)
	logRequestedTaskController("custom-resource", taskName)
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

func (ctrl *customResourceController) DeleteCustomResource(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteCustomResourceInputParams)

	input.CustomResourceName = ctx.Param("name")
	input.NamespaceName = ctx.Query("namespace")
	input.CustomResourceSGVR.Resource = ctx.Query("resource")
	input.CustomResourceSGVR.Version = ctx.Query("version")
	input.CustomResourceSGVR.Group = ctx.Query("group")

	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.CustomResourceService().DeleteCustomResource)
	logRequestedTaskController("custom-resource", taskName)
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
