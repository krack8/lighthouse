package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/worker"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
	"k8s.io/api/core/v1"
)

type EndpointsControllerInterface interface {
	GetEndpointsList(ctx *gin.Context)
	GetEndpointsDetails(ctx *gin.Context)
	DeployEndpoints(ctx *gin.Context)
	DeleteEndpoints(ctx *gin.Context)
}

type endpointsController struct {
}

var epc endpointsController

func EndpointsController() *endpointsController {
	return &epc
}

func (ctrl *endpointsController) GetEndpointsList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetEndpointsListInputParams)
	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
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
			log.Logger.Info("Filter by param for Endpoints List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	input.Namespace = ctx.Query("namespace")
	input.Continue = ctx.Query("continue")
	input.Search = ctx.Query("q")
	input.Limit = ctx.Query("limit")
	taskName := tasks.GetTaskName(k8s.EndpointsService().GetEndpointsList)
	logRequestedTaskController("endpoint", taskName)
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

func (ctrl *endpointsController) GetEndpointsDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetEndpointsDetailsInputParams)
	input.EndpointsName = ctx.Param("name")
	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	input.Namespace = ctx.Query("namespace")
	taskName := tasks.GetTaskName(k8s.EndpointsService().GetEndpointsDetails)
	logRequestedTaskController("endpoint", taskName)
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

func (ctrl *endpointsController) DeployEndpoints(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(v1.Endpoints)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy Endpoints payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	input := new(k8s.DeployEndpointsInputParams)
	input.Endpoints = payload
	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.EndpointsService().DeployEndpoints)
	logRequestedTaskController("endpoint", taskName)
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

func (ctrl *endpointsController) DeleteEndpoints(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteEndpointsInputParams)
	input.EndpointsName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.EndpointsService().DeleteEndpoints)
	logRequestedTaskController("endpoint", taskName)
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
