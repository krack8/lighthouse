package k8s

import (
	"context"
	"errors"
	cfg "github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

type LoadBalancerServiceInterface interface {
	GetLoadBalancerList(c context.Context, p GetLoadBalancerListInputParams) (interface{}, error)
	GetLoadBalancerDetails(c context.Context, p GetLoadBalancerDetailsInputParams) (interface{}, error)
}

type loadBalancerService struct{}

var lbs loadBalancerService

func LoadBalancerService() *loadBalancerService {
	return &lbs
}

type GetLoadBalancerListInputParams struct {
	NamespaceName string
	Search        string
	Labels        map[string]string
	output        []corev1.Service
}

func (p *GetLoadBalancerListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching loadBalancer list")
	loadBalancerClient := cfg.GetKubeClientSet().CoreV1().Services(p.NamespaceName)
	listOptions := metav1.ListOptions{}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions = metav1.ListOptions{
			LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		}
	}
	if p.Search != "" {
		listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
	}
	loadBalancerList, err := loadBalancerClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to get loadBalancer list", "err", err.Error())
		return err
	}
	p.output = loadBalancerList.Items
	return nil
}

func (p *GetLoadBalancerListInputParams) PostProcess(c context.Context) error {
	var temp []corev1.Service
	for _, each := range p.output {
		if each.Spec.Type == "LoadBalancer" {
			temp = append(temp, each)
		}
	}
	p.output = temp
	return nil
}

func (svc *loadBalancerService) GetLoadBalancerList(c context.Context, p GetLoadBalancerListInputParams) (interface{}, error) {
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

type GetLoadBalancerDetailsInputParams struct {
	NamespaceName    string
	LoadBalancerName string
	output           corev1.Service
}

func (p *GetLoadBalancerDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching loadBalancer details of ....", p.NamespaceName)
	loadBalancersClient := cfg.GetKubeClientSet().CoreV1().Services(p.NamespaceName)
	output, err := loadBalancersClient.Get(context.Background(), p.LoadBalancerName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get loadBalancer ", p.LoadBalancerName, "err", err.Error())
		return err
	}
	if output.Spec.Type != "LoadBalancer" {
		log.Logger.Errorw("No loadBalancer type svc", "value", p.LoadBalancerName)
		return errors.New("No load balancer type svc named " + p.LoadBalancerName)
	}
	p.output = *output
	return nil
}

func (svc *loadBalancerService) GetLoadBalancerDetails(c context.Context, p GetLoadBalancerDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}
