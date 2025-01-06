package k8s

import (
	"context"
	"encoding/json"
	cfg "github.com/krack8/lighthouse/pkg/config"
	"github.com/krack8/lighthouse/pkg/dto"
	"github.com/krack8/lighthouse/pkg/log"
)

type PodMetricsServiceInterface interface {
	GetPodMetricsList(c context.Context, p GetPodMetricsListInputParams) (interface{}, error)
	GetPodMetricsDetails(c context.Context, p GetPodMetricsDetailsInputParams) (interface{}, error)
}

type podMetricsService struct{}

var pms podMetricsService

func PodMetricsService() *podMetricsService {
	return &pms
}

type GetPodMetricsListInputParams struct {
	NamespaceName string
	Labels        map[string]string
	output        dto.PodMetricsList
}

func (p *GetPodMetricsListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching config map list")
	var podMetricsList dto.PodMetricsList
	podMetricsClient := cfg.GetKubeClientSet().RESTClient()
	data, err := podMetricsClient.Get().AbsPath(cfg.MetricsAbsPath + p.NamespaceName + "/pods").DoRaw(context.Background())
	if err != nil {
		log.Logger.Errorw("Failed to get podMetrics list", "err", err.Error())
		return err
	}
	err = json.Unmarshal(data, &podMetricsList)
	if err != nil {
		log.Logger.Errorw("Failed to unmarshal podMetrics list", "err", err.Error())
		return err
	}
	p.output = podMetricsList
	return nil
}

func (svc *podMetricsService) GetPodMetricsList(c context.Context, p GetPodMetricsListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetPodMetricsDetailsInputParams struct {
	NamespaceName string
	PodName       string
	output        dto.PodMetrics
}

func (p *GetPodMetricsDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching podMetrics details of ....", p.NamespaceName)
	var podMetrics dto.PodMetrics

	podMetricsClient := cfg.GetKubeClientSet().RESTClient()
	data, err := podMetricsClient.Get().AbsPath(cfg.MetricsAbsPath + p.NamespaceName + "/pods/" + p.PodName).DoRaw(context.Background())
	if err != nil {
		log.Logger.Errorw("Failed to get podMetrics details of "+p.PodName, "err", err.Error())
		return err
	}
	err = json.Unmarshal(data, &podMetrics)
	if err != nil {
		log.Logger.Errorw("Failed to get podMetrics details of "+p.PodName, "err", err.Error())
		return err
	}
	p.output = podMetrics
	return nil
}

func (svc *podMetricsService) GetPodMetricsDetails(c context.Context, p GetPodMetricsDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}
