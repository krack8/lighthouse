package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/krack8/lighthouse/pkg/agent/tasks"
	"github.com/krack8/lighthouse/pkg/common/k8s"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/controller/core"
	corev1 "k8s.io/api/core/v1"
	"net/http"
	"strconv"
)

type PodControllerInterface interface {
	GetPodList(ctx *gin.Context)
	GetPodDetails(ctx *gin.Context)
	GetPodStats(ctx *gin.Context)
	GetPodLogs(ctx *gin.Context)
	DeployPod(ctx *gin.Context)
	DeletePod(ctx *gin.Context)
}

type podController struct {
}

var pc podController

func PodController() *podController {
	return &pc
}

func (ctrl *podController) GetPodList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetPodListInputParams)

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
	input.Continue = ctx.Query("continue")
	input.Search = ctx.Query("q")
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
			log.Logger.Info("Filter by param for Pod List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	taskName := tasks.GetTaskName(k8s.PodService().GetPodList)
	logRequestedTaskController("pod", taskName)
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

func (ctrl *podController) GetPodDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetPodDetailsInputParams)
	input.PodName = ctx.Param("name")

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
	taskName := tasks.GetTaskName(k8s.PodService().GetPodDetails)
	logRequestedTaskController("pod", taskName)
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

func (ctrl *podController) DeployPod(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(corev1.Pod)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy pod payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}
	clusterGroup := ctx.Query("cluster_id")
	if clusterGroup == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroup)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	input := new(k8s.DeployPodInputParams)
	input.Pod = payload
	if input.Pod.Namespace == "" {
		log.Logger.Errorw("namespace required in payload", "value", "pod deploy")
		SendErrorResponse(ctx, ErrNamespaceEmpty)
		return
	}
	taskName := tasks.GetTaskName(k8s.PodService().DeployPod)
	logRequestedTaskController("pod", taskName)
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

func (ctrl *podController) DeletePod(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeletePodInputParams)
	input.PodName = ctx.Param("name")

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
	taskName := tasks.GetTaskName(k8s.PodService().DeletePod)
	logRequestedTaskController("pod", taskName)
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

func (ctrl *podController) ExecPod(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.PodExecInputParams)
	input.PodName = ctx.Param("name")

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
	containerName := ctx.Query("container")
	if containerName == "" {
		log.Logger.Errorw("Container required in query params", "value", containerName)
		SendErrorResponse(ctx, "Container required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	input.ContainerName = containerName

	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal PodExec Task input", "err", err.Error())
	}

	var wsocket = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := wsocket.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Logger.Errorw("unable to initiate websocket connection", "err", err.Error())
		SendErrorResponse(ctx, "Unable to initiate websocket connection")
		return
	}

	_, err = core.GetAgentManager().SendTerminalExecRequestToAgent(ctx, string(inputTask), clusterGroup, conn)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}

	SendResponse(ctx, result)
}

func (ctrl *podController) GetPodStats(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetPodStatsInputParams)

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
			log.Logger.Info("Filter by param for Pod List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	taskName := tasks.GetTaskName(k8s.PodService().GetPodStats)
	logRequestedTaskController("pod", taskName)
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

func (ctrl *podController) GetPodLogs(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.GetPodLogsInputParams)
	input.Pod = ctx.Param("name")
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
	input.Container = ctx.Query("container")
	if ctx.Query("lines") != "" {
		tailLines, err := strconv.ParseInt(ctx.Query("lines"), 10, 64)
		if err == nil {
			input.TailLines = &tailLines
		}
	}
	if ctx.Query("since") != "" {
		sinceSeconds, err := strconv.ParseInt(ctx.Query("since"), 10, 64)
		if err == nil {
			input.SinceSeconds = &sinceSeconds
		}
	}
	input.Timestamps = ctx.Query("timestamps")
	input.Previous = ctx.Query("previous")
	taskName := tasks.GetTaskName(k8s.PodService().GetPodLogs)
	logRequestedTaskController("pod", taskName)
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
