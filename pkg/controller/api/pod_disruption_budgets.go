package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/core"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
	v1 "k8s.io/api/policy/v1"
)

type PodDisruptionBudgetsControllerInterface interface {
	GetPodDisruptionBudgetsList(ctx *gin.Context)
	GetPodDisruptionBudgetsDetails(ctx *gin.Context)
	DeployPodDisruptionBudgets(ctx *gin.Context)
	DeletePodDisruptionBudgets(ctx *gin.Context)
}

type podDisruptionBudgetsController struct {
}

var pdbc podDisruptionBudgetsController

func PodDisruptionBudgetsController() *podDisruptionBudgetsController {
	return &pdbc
}

func (ctrl *podDisruptionBudgetsController) GetPodDisruptionBudgetsList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetPodDisruptionBudgetsListInputParams)

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
			log.Logger.Info("Filter by param for PodDisruptionBudgets List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	taskName := tasks.GetTaskName(k8s.PodDisruptionBudgetsService().GetPodDisruptionBudgetsList)
	logRequestedTaskController("pod-disruption-budget", taskName)
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

func (ctrl *podDisruptionBudgetsController) GetPodDisruptionBudgetsDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetPodDisruptionBudgetsDetailsInputParams)
	input.PodDisruptionBudgetsName = ctx.Param("name")

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
	taskName := tasks.GetTaskName(k8s.PodDisruptionBudgetsService().GetPodDisruptionBudgetsDetails)
	logRequestedTaskController("pod-disruption-budget", taskName)
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

func (ctrl *podDisruptionBudgetsController) DeployPodDisruptionBudgets(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(v1.PodDisruptionBudget)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy PodDisruptionBudgets payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	input := new(k8s.DeployPodDisruptionBudgetsInputParams)
	input.PodDisruptionBudgets = payload
	if input.PodDisruptionBudgets.Namespace == "" {
		log.Logger.Errorw("namespace required in payload", "value", "podDisruptionBudgets deploy")
		SendErrorResponse(ctx, ErrNamespaceEmpty)
		return
	}
	taskName := tasks.GetTaskName(k8s.PodDisruptionBudgetsService().DeployPodDisruptionBudgets)
	logRequestedTaskController("pod-disruption-budget", taskName)
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

func (ctrl *podDisruptionBudgetsController) DeletePodDisruptionBudgets(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeletePodDisruptionBudgetsInputParams)
	input.PodDisruptionBudgetsName = ctx.Param("name")

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
	taskName := tasks.GetTaskName(k8s.PodDisruptionBudgetsService().DeletePodDisruptionBudgets)
	logRequestedTaskController("pod-disruption-budget", taskName)
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
