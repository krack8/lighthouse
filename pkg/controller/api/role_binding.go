package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/worker"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
	rbacv1 "k8s.io/api/rbac/v1"
)

type RoleBindingControllerInterface interface {
	GetRoleBindingList(ctx *gin.Context)
	GetRoleBindingDetails(ctx *gin.Context)
	DeployRoleBinding(ctx *gin.Context)
	DeleteRoleBinding(ctx *gin.Context)
}

type roleBindingController struct {
}

var rbc roleBindingController

func RoleBindingController() *roleBindingController {
	return &rbc
}

func (ctrl *roleBindingController) GetRoleBindingList(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetRoleBindingListInputParams)

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
			log.Logger.Info("Filter by param for RoleBinding List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	taskName := tasks.GetTaskName(k8s.RoleBindingService().GetRoleBindingList)
	logRequestedTaskController("role-binding", taskName)
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

func (ctrl *roleBindingController) GetRoleBindingDetails(ctx *gin.Context) {
	var result ResponseDTO

	input := new(k8s.GetRoleBindingDetailsInputParams)
	input.RoleBindingName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	taskName := tasks.GetTaskName(k8s.RoleBindingService().GetRoleBindingDetails)
	logRequestedTaskController("role-binding", taskName)
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

func (ctrl *roleBindingController) DeployRoleBinding(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(rbacv1.RoleBinding)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy RoleBinding payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	input := new(k8s.DeployRoleBindingInputParams)
	input.RoleBinding = payload
	if input.RoleBinding.Namespace == "" {
		log.Logger.Errorw("namespace required in payload", "value", "roleBinding deploy")
		SendErrorResponse(ctx, ErrNamespaceEmpty)
		return
	}
	taskName := tasks.GetTaskName(k8s.RoleBindingService().DeployRoleBinding)
	logRequestedTaskController("role-binding", taskName)
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

func (ctrl *roleBindingController) DeleteRoleBinding(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s.DeleteRoleBindingInputParams)
	input.RoleBindingName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	taskName := tasks.GetTaskName(k8s.RoleBindingService().DeleteRoleBinding)
	logRequestedTaskController("role-binding", taskName)
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
