package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/log"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

type HpaServiceInterface interface {
	GetHpaList(c context.Context, p GetHpaListInputParams) (interface{}, error)
	GetHpaDetails(c context.Context, p GetHpaDetailsInputParams) (interface{}, error)
}

type hpaService struct{}

var hpas hpaService

func HpaService() *hpaService {
	return &hpas
}

const (
	HpaApiVersion = "autoscaling/v2"
	HpaKind       = "HorizontalPodAutoscaler"
)

type GetHpaListInputParams struct {
	NamespaceName string
	Search        string
	Labels        map[string]string
	output        []autoscalingv1.HorizontalPodAutoscaler
}

func (p *GetHpaListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching hpa list")
	hpaClient := GetKubeClientSet().AutoscalingV1().HorizontalPodAutoscalers(p.NamespaceName)
	listOptions := metav1.ListOptions{}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions.LabelSelector = labels.Set(labelSelector.MatchLabels).String()
	}
	if p.Search != "" {
		listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
	}
	hpaList, err := hpaClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to get hpa list", "err", err.Error())
		return err
	}
	p.output = hpaList.Items
	return nil
}

func (p *GetHpaListInputParams) PostProcess(ctx context.Context) error {
	for i := 0; i < len(p.output); i++ {
		p.output[i].ManagedFields = nil
		p.output[i].APIVersion = HpaApiVersion
		p.output[i].Kind = HpaKind
	}
	return nil
}

func (svc *hpaService) GetHpaList(c context.Context, p GetHpaListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}
	_ = p.PostProcess(c)
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetHpaDetailsInputParams struct {
	NamespaceName string
	HpaName       string
	output        autoscalingv1.HorizontalPodAutoscaler
}

func (p *GetHpaDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching hpa details of ....", p.NamespaceName)
	hpasClient := GetKubeClientSet().AutoscalingV1().HorizontalPodAutoscalers(p.NamespaceName)
	output, err := hpasClient.Get(context.Background(), p.HpaName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get hpa ", p.HpaName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.ManagedFields = nil
	p.output.APIVersion = HpaApiVersion
	p.output.Kind = HpaKind
	return nil
}

func (svc *hpaService) GetHpaDetails(c context.Context, p GetHpaDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}
