package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/core"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
	networkingv1 "k8s.io/api/networking/v1"
)

type IngressControllerInterface interface {
	GetIngressList(ctx *gin.Context)
	GetIngressDetails(ctx *gin.Context)
	DeployIngress(ctx *gin.Context)
	DeleteIngress(ctx *gin.Context)
}

type ingressController struct {
}

var ic ingressController

func IngressController() *ingressController {
	return &ic
}

func (ctrl *ingressController) GetIngressList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetIngressListInputParams)

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
	input.Continue = ctx.Query("continue")
	input.Limit = ctx.Query("limit")

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
			log.Logger.Info("Filter by param for Ingress List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	taskName := tasks.GetTaskName(k8s.IngressService().GetIngressList)
	logRequestedTaskController("ingress", taskName)
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

func (ctrl *ingressController) GetIngressDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetIngressDetailsInputParams)
	input.IngressName = ctx.Param("name")

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
	taskName := tasks.GetTaskName(k8s.IngressService().GetIngressDetails)
	logRequestedTaskController("ingress", taskName)
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

func (ctrl *ingressController) DeployIngress(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(networkingv1.Ingress)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy Ingress payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	input := new(k8s.DeployIngressInputParams)
	input.Ingress = payload
	if input.Ingress.Namespace == "" {
		log.Logger.Errorw("namespace required in payload", "value", "ingress deploy")
		SendErrorResponse(ctx, ErrNamespaceEmpty)
		return
	}
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.IngressService().DeployIngress)
	logRequestedTaskController("ingress", taskName)
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

func (ctrl *ingressController) DeleteIngress(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteIngressInputParams)
	input.IngressName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	taskName := tasks.GetTaskName(k8s.IngressService().DeleteIngress)
	logRequestedTaskController("ingress", taskName)
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
