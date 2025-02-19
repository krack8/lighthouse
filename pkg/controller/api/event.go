package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/core"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
)

type EventControllerInterface interface {
	GetEventList(ctx *gin.Context)
	GetEventDetails(ctx *gin.Context)
}

type eventController struct {
}

var ec eventController

func EventController() *eventController {
	return &ec
}

func (ctrl *eventController) GetEventList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetEventListInputParams)

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	input.Search = ctx.Query("q")

	queryInvolvedObjectName := ctx.Query("involved_object_name")
	input.InvolvedObjectName = queryInvolvedObjectName

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
			log.Logger.Info("Filter by param for Event List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	taskName := tasks.GetTaskName(k8s.EventService().GetEventList)
	logRequestedTaskController("event", taskName)
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

func (ctrl *eventController) GetEventDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetEventDetailsInputParams)
	input.EventName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	taskName := tasks.GetTaskName(k8s.EventService().GetEventDetails)
	logRequestedTaskController("event", taskName)
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
