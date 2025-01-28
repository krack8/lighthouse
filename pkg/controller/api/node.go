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

type NodeControllerInterface interface {
	GetNodeList(ctx *gin.Context)
	GetNodeDetails(ctx *gin.Context)
	NodeCordon(ctx *gin.Context)
	NodeTaint(ctx *gin.Context)
	NodeUnTaint(ctx *gin.Context)
}

type nodeController struct {
}

var nc nodeController

func NodeController() *nodeController {
	return &nc
}

func (ctrl *nodeController) GetNodeList(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.GetNodeListInputParams)
	input.Search = ctx.Query("q")

	queryLabel := ctx.Query("labels")
	if queryLabel != "" {
		jsonLabel := []byte(queryLabel)
		queryLabelMap := map[string]string{}

		err := json.Unmarshal([]byte(jsonLabel), &queryLabelMap)
		if err != nil {
			log.Logger.Errorw("query labels unmarshal error ", "err", err.Error())
		}
		if queryLabelMap != nil {
			input.Labels = queryLabelMap
			log.Logger.Infow("Filter by param for Namespace List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	taskName := tasks.GetTaskName(k8s.NodeService().GetNodeList)
	logRequestedTaskController("node", taskName)
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

func (ctrl *nodeController) GetNodeDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetNodeInputParams)
	input.NodeName = ctx.Param("name")
	taskName := tasks.GetTaskName(k8s.NodeService().GetNodeDetails)
	logRequestedTaskController("node", taskName)
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

func (ctrl *nodeController) NodeCordon(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.NodeCordonInputParams)
	input.NodeName = ctx.Param("name")
	taskName := tasks.GetTaskName(k8s.NodeService().NodeCordon)
	logRequestedTaskController("node", taskName)
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

func (ctrl *nodeController) NodeTaint(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.NodeTaintInputParams)
	input.NodeName = ctx.Param("name")

	payload := new(dto.TaintList)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind node taint payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}
	log.Logger.Debugw("node taint payload ", payload)

	input.TaintList = payload
	taskName := tasks.GetTaskName(k8s.NodeService().NodeTaint)
	logRequestedTaskController("node", taskName)
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

func (ctrl *nodeController) NodeUnTaint(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.NodeUnTaintInputParams)
	input.NodeName = ctx.Param("name")
	payload := new(dto.UnTaintKeys)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind node untaint payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}
	log.Logger.Debugw("node untaint payload ", payload)
	input.Keys = payload.Keys
	taskName := tasks.GetTaskName(k8s.NodeService().NodeUnTaint)
	logRequestedTaskController("node", taskName)
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
