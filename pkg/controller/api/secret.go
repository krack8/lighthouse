package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/worker"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
	corev1 "k8s.io/api/core/v1"
)

type SecretControllerInterface interface {
	GetSecretList(ctx *gin.Context)
	GetSecretDetails(ctx *gin.Context)
	DeploySecret(ctx *gin.Context)
	DeleteSecret(ctx *gin.Context)
}

type secretController struct {
}

var sc secretController

func SecretController() *secretController {
	return &sc
}

func (ctrl *secretController) GetSecretList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetSecretListInputParams)

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
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

		err := json.Unmarshal([]byte(jsonLabel), &queryLabelMap)
		if err != nil {
			log.Logger.Error("query labels unmarshal error ", "err", err.Error())
		}
		if queryLabelMap != nil {
			input.Labels = queryLabelMap
			log.Logger.Info("Filter by param for Secret List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	taskName := tasks.GetTaskName(k8s.SecretService().GetSecretList)
	logRequestedTaskController("secret", taskName)
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

func (ctrl *secretController) GetSecretDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetSecretDetailsInputParams)
	input.SecretName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	taskName := tasks.GetTaskName(k8s.SecretService().GetSecretDetails)
	logRequestedTaskController("secret", taskName)
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

func (ctrl *secretController) DeploySecret(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(corev1.Secret)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy secret payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	input := new(k8s.DeploySecretInputParams)
	input.Secret = payload
	taskName := tasks.GetTaskName(k8s.SecretService().DeploySecret)
	logRequestedTaskController("secret", taskName)
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

func (ctrl *secretController) DeleteSecret(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteSecretInputParams)
	input.SecretName = ctx.Param("name")
	taskName := tasks.GetTaskName(k8s.SecretService().DeleteSecret)
	logRequestedTaskController("secret", taskName)
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
