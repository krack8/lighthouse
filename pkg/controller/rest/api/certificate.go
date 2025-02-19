package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/agent/tasks"
	"github.com/krack8/lighthouse/pkg/common/dto"
	k8s2 "github.com/krack8/lighthouse/pkg/common/k8s"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/controller/core"
)

type CertificateControllerInterface interface {
	GetCertificateList(ctx *gin.Context)
	GetCertificateDetails(ctx *gin.Context)
	DeployCertificate(ctx *gin.Context)
	DeleteCertificate(ctx *gin.Context)
}

type certificateController struct {
}

var cc certificateController

func CertificateController() *certificateController {
	return &cc
}

func (ctrl *certificateController) GetCertificateList(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s2.GetCertificateListInputParams)

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
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
			log.Logger.Info("Filter by param for Certificate List param Map: ", queryLabelMap, " values: ", ctx.Query("labels"))
		}
	}
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal GetCertificateList Task input ", "err", err.Error())
	}
	taskName := tasks.GetTaskName(k8s2.CertificateService().GetCertificateList)
	logRequestedTaskController("certificate", taskName)
	core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroupName)
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroupName)
	if err != nil {
		k8s2.SendErrorResponse(ctx, err.Error())
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	SendResponse(ctx, result)
}

func (ctrl *certificateController) GetCertificateDetails(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s2.GetCertificateDetailsInputParams)
	input.CertificateName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}

	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}

	input.NamespaceName = queryNamespace
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal GetCertificateList Task input ", "err", err.Error())
	}
	taskName := tasks.GetTaskName(k8s2.CertificateService().GetCertificateDetails)
	logRequestedTaskController("certificate", taskName)
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroupName)
	if err != nil {
		k8s2.SendErrorResponse(ctx, err.Error())
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	SendResponse(ctx, result)
}

func (ctrl *certificateController) DeployCertificate(ctx *gin.Context) {
	var result ResponseDTO
	payload := new(dto.Certificate)
	if err := ctx.Bind(payload); err != nil {
		log.Logger.Errorw("Failed to bind deploy Certificate payload", "err", err.Error())
		SendErrorResponse(ctx, err.Error())
		return
	}

	input := new(k8s2.DeployCertificateInputParams)
	input.Certificate = payload
	if input.Certificate.Namespace == "" {
		log.Logger.Errorw("namespace required in payload", "value", "certificate deploy")
		SendErrorResponse(ctx, ErrNamespaceEmpty)
		return
	}
	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal DeployCertificate Task input ", "err", err.Error())
	}
	taskName := tasks.GetTaskName(k8s2.CertificateService().GetCertificateDetails)
	logRequestedTaskController("certificate", taskName)
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroupName)
	if err != nil {
		k8s2.SendErrorResponse(ctx, err.Error())
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	SendResponse(ctx, result)
}

func (ctrl *certificateController) DeleteCertificate(ctx *gin.Context) {
	var result ResponseDTO
	input := new(k8s2.DeleteCertificateInputParams)
	input.CertificateName = ctx.Param("name")

	queryNamespace := ctx.Query("namespace")
	if queryNamespace == "" {
		log.Logger.Errorw("Namespace required in query params", "value", queryNamespace)
		SendErrorResponse(ctx, "Namespace required in query params")
		return
	}
	clusterGroupName := ctx.Query("cluster_id")
	if clusterGroupName == "" {
		log.Logger.Errorw("Cluster id required in query params", "value", clusterGroupName)
		SendErrorResponse(ctx, "Cluster id required in query params")
		return
	}
	input.NamespaceName = queryNamespace
	inputTask, err := json.Marshal(input)
	if err != nil {
		log.Logger.Errorw("unable to marshal GetCertificateList Task input ", "err", err.Error())
	}
	taskName := tasks.GetTaskName(k8s2.CertificateService().DeleteCertificate)
	logRequestedTaskController("certificate", taskName)
	res, err := core.GetAgentManager().SendTaskToAgent(ctx, taskName, inputTask, clusterGroupName)
	if err != nil {
		k8s2.SendErrorResponse(ctx, err.Error())
	}
	err = json.Unmarshal([]byte(res.Output), &result)
	if err != nil {
		SendErrorResponse(ctx, err.Error())
		return
	}
	SendResponse(ctx, result)
}
