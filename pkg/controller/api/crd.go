package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/worker"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type CrdControllerInterface interface {
	GetCrdList(ctx *gin.Context)
	GetCrdDetails(ctx *gin.Context)
	DeployCrd(ctx *gin.Context)
	DeleteCrd(ctx *gin.Context)
}

type crdController struct {
}

var crdc crdController

func CrdController() *crdController {
	return &crdc
}

func (ctrl *crdController) GetCrdList(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.GetCrdListInputParams)
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
			log.Logger.Info("Filter by param for Crd List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	input.Continue = ctx.Query("continue")
	input.Search = ctx.Query("q")
	input.Limit = ctx.Query("limit")
	taskName := tasks.GetTaskName(k8s.CrdService().GetCrdList)
	logRequestedTaskController("crd", taskName)
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

func (ctrl *crdController) GetCrdDetails(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.GetCrdDetailsInputParams)
	input.CrdName = ctx.Param("name")
	taskName := tasks.GetTaskName(k8s.CrdService().GetCrdList)
	logRequestedTaskController("crd", taskName)
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

func (ctrl *crdController) DeployCrd(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(v1.CustomResourceDefinition)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy Crd payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}
	input := new(k8s.DeployCrdInputParams)
	input.Crd = payload
	taskName := tasks.GetTaskName(k8s.CrdService().DeployCrd)
	logRequestedTaskController("crd", taskName)
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

func (ctrl *crdController) DeleteCrd(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteCrdInputParams)
	input.CrdName = ctx.Param("name")
	taskName := tasks.GetTaskName(k8s.CrdService().DeleteCrd)
	logRequestedTaskController("crd", taskName)
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
