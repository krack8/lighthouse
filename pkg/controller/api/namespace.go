package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/worker"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
)

type NamespaceControllerInterface interface {
	GetNamespaceList(ctx *gin.Context)
	//GetNamespaceNameList(ctx *gin.Context)
	//GetNamespaceDetails(ctx *gin.Context)
	//DeployNamespace(ctx *gin.Context)
	//DeleteNamespace(ctx *gin.Context)
}

type namespaceController struct {
}

var nsc namespaceController

func NamespaceController() *namespaceController {
	return &nsc
}

func (ctrl *namespaceController) GetNamespaceList(ctx *gin.Context) {
	var result k8s.ResponseDTO

	input := new(k8s.GetNamespaceListInputParams)
	input.Search = ctx.Query("q")
	input.Continue = ctx.Query("continue")
	input.Limit = ctx.Query("limit")
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
	input.Search = ctx.Query("q")
	input.Continue = ctx.Query("continue")
	input.Limit = ctx.Query("limit")
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal GetNamespaceList Task input ", "err", err.Error())
	}
	log.Logger.Debugw("GetNamespaceList task started...", "Task", "GetNamespaceList")
	group := ctx.Query("group")
	payload := ctx.Query("payload")
	if group == "" || payload == "" {
		k8s.SendErrorResponse(ctx, "Missing group or payload param")
		return
	}
	taskName := tasks.GetTaskName(k8s.NamespaceService().GetNamespaceList)
	log.Logger.Info("Printing task name from namespace controller: " + taskName)
	res, err := worker.TaskToAgent().SendToWorker(ctx, group, payload, taskName, inputTask)
	if err != nil {
		k8s.SendErrorResponse(ctx, err.Error())
	}
	output := k8s.OutputNamespaceList{}
	err = json.Unmarshal([]byte(res.Output), &output)
	if err != nil {
		k8s.SendErrorResponse(ctx, err.Error())
	}
	result.Data = output
	k8s.SendSuccessResponse(ctx, result)
}
