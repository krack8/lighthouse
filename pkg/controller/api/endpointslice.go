package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/worker"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
	"k8s.io/api/discovery/v1"
)

type EndpointSliceControllerInterface interface {
	GetEndpointSliceList(ctx *gin.Context)
	GetEndpointSliceDetails(ctx *gin.Context)
	DeployEndpointSlice(ctx *gin.Context)
	DeleteEndpointSlice(ctx *gin.Context)
}

type endpointSliceController struct {
}

var epsc endpointSliceController

func EndpointSliceController() *endpointSliceController {
	return &epsc
}

func (ctrl *endpointSliceController) GetEndpointSliceList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetEndpointSliceListInputParams)
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

		err := json.Unmarshal([]byte(jsonLabel), &queryLabelMap)
		if err != nil {
			log.Logger.Error("query labels unmarshal error ", "err", err.Error())
		}
		if queryLabelMap != nil {
			input.Labels = queryLabelMap
			log.Logger.Info("Filter by param for EndpointSlice List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	input.Namespace = ctx.Query("namespace")
	input.Search = ctx.Query("q")
	input.Continue = ctx.Query("continue")
	input.Limit = ctx.Query("limit")
	taskName := tasks.GetTaskName(k8s.EndpointSliceService().GetEndpointSliceList)
	logRequestedTaskController("endpoint-slice", taskName)
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

func (ctrl *endpointSliceController) GetEndpointSliceDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetEndpointSliceDetailsInputParams)
	input.EndpointSliceName = ctx.Param("name")
	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	input.Namespace = ctx.Query("namespace")
	taskName := tasks.GetTaskName(k8s.EndpointSliceService().GetEndpointSliceDetails)
	logRequestedTaskController("endpoint-slice", taskName)
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

func (ctrl *endpointSliceController) DeployEndpointSlice(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(v1.EndpointSlice)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy endpointSlice payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	input := new(k8s.DeployEndpointSliceInputParams)
	input.EndpointSlice = payload
	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.EndpointSliceService().DeployEndpointSlice)
	logRequestedTaskController("endpoint-slice", taskName)
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

func (ctrl *endpointSliceController) DeleteEndpointSlice(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteEndpointSliceInputParams)
	input.EndpointSliceName = ctx.Param("name")
	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.EndpointSliceService().DeleteEndpointSlice)
	logRequestedTaskController("endpoint-slice", taskName)
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
