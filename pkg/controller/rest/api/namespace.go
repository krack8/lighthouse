package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/agent/tasks"
	k8s2 "github.com/krack8/lighthouse/pkg/common/k8s"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/controller/core"
	corev1 "k8s.io/api/core/v1"
)

type NamespaceControllerInterface interface {
	GetNamespaceList(ctx *gin.Context)
	GetNamespaceNameList(ctx *gin.Context)
	GetNamespaceDetails(ctx *gin.Context)
	DeployNamespace(ctx *gin.Context)
	DeleteNamespace(ctx *gin.Context)
}

type namespaceController struct {
}

var nsc namespaceController

func NamespaceController() *namespaceController {
	return &nsc
}

func (ctrl *namespaceController) GetNamespaceList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s2.GetNamespaceListInputParams)
	input.Search = ctx.Query("q")
	input.Continue = ctx.Query("continue")
	input.Limit = ctx.Query("limit")
	queryLabel := ctx.Query("labels")
	if queryLabel != "" {
		jsonLabel := []byte(queryLabel)
		queryLabelMap := map[string]string{}

		err := json.Unmarshal(jsonLabel, &queryLabelMap)
		if err != nil {
			log.Logger.Errorw("query labels unmarshal error ", "err", err.Error())
		}
		if queryLabelMap != nil {
			input.Labels = queryLabelMap
			log.Logger.Infow("Filter by param for Namespace List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	input.Search = ctx.Query("q")
	input.Continue = ctx.Query("continue")
	input.Limit = ctx.Query("limit")
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal GetNamespaceList Task input ", "err", err.Error())
	}
	log.Logger.Debugw("GetNamespaceList task started...", "Task", "GetNamespaceList")
	//group := ctx.Query("group")
	//payload := ctx.Query("payload")
	////if group == "" || payload == "" {
	////	k8s.SendErrorResponse(ctx, "Missing group or payload param")
	////	return
	////}
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s2.NamespaceService().GetNamespaceList)
	logRequestedTaskController("namespace", taskName)
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroup)
	if err != nil {
		k8s2.SendErrorResponse(ctx, err.Error())
		return
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	SendResponse(ctx, result)
}

func (ctrl *namespaceController) GetNamespaceNameList(ctx *gin.Context) {
	var result ResponseDTO
	var input = new(k8s2.GetNamespaceNamesInputParams)
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal GetNamespaceNameList Task input ", "err", err.Error())
	}
	taskName := tasks.GetTaskName(k8s2.NamespaceService().GetNamespaceNameList)
	logRequestedTaskController("namespace", taskName)
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

func (ctrl *namespaceController) DeployNamespace(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(corev1.Namespace)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy Namespace payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}
	log.Logger.Debugw("deploy namespace payload ", payload)

	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}

	input := new(k8s2.DeployNamespaceInputParams)
	input.Namespace = payload
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal deploy namespace Task input ", "err", err.Error())
	}
	taskName := tasks.GetTaskName(k8s2.NamespaceService().DeployNamespace)
	logRequestedTaskController("namespace", taskName)
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

func (ctrl *namespaceController) GetNamespaceDetails(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s2.GetNamespaceInputParams)
	taskName := tasks.GetTaskName(k8s2.NamespaceService().GetNamespaceDetails)
	logRequestedTaskController("namespace", taskName)
	input.NamespaceName = ctx.Param("name")
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
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

func (ctrl *namespaceController) DeleteNamespace(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s2.DeleteNamespaceInputParams)
	input.NamespaceName = ctx.Param("name")
	taskName := tasks.GetTaskName(k8s2.NamespaceService().DeleteNamespace)
	logRequestedTaskController("namespace", taskName)
	input.NamespaceName = ctx.Param("name")
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
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
