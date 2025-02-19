package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	"strconv"
	"strings"
)

type IngressServiceInterface interface {
	GetIngressList(c context.Context, p GetIngressListInputParams) (interface{}, error)
	GetIngressDetails(c context.Context, p GetIngressDetailsInputParams) (interface{}, error)
	DeployIngress(c context.Context, p DeployIngressInputParams) (interface{}, error)
	DeleteIngress(c context.Context, p DeleteIngressInputParams) (interface{}, error)
}

type ingressService struct{}

var ings ingressService

func IngressService() *ingressService {
	return &ings
}

type OutputIngressList struct {
	Result    []networkingv1.Ingress
	Resource  string
	Remaining int64
	Total     int
}

type GetIngressListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputIngressList
}

func (p *GetIngressListInputParams) Find(c context.Context, ingressClient v1.IngressInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "ingress")
	filteredIngresses := []networkingv1.Ingress{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		ingressList, err := ingressClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get ingress list", "err", err.Error())
			return err
		}

		for _, ingress := range ingressList.Items {
			if strings.Contains(ingress.Name, p.Search) {
				filteredIngresses = append(filteredIngresses, ingress)
			}
		}
		length = len(filteredIngresses)
		nextPageToken = ingressList.Continue
		if ingressList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		ingressList, err := ingressClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get ingress list", "err", err.Error())
			return err
		}
		for _, ingress := range ingressList.Items {
			if strings.Contains(ingress.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredIngresses
	p.output.Total = len(filteredIngresses)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetIngressListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching ingress list")
	ingressClient := GetKubeClientSet().NetworkingV1().Ingresses(p.NamespaceName)
	limit := config.PageLimit
	if p.Limit != "" {
		limit, _ = strconv.ParseInt(p.Limit, 10, 64)
	}
	listOptions := metav1.ListOptions{Limit: limit, Continue: p.Continue}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions = metav1.ListOptions{
			LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		}
	}
	var err error
	var ingressList *networkingv1.IngressList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, ingressClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get ingress list", "err", err.Error())
			return err
		}
		return nil
	} else {
		ingressList, err = ingressClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get ingress list", "err", err.Error())
			return err
		}

		ingressList, err = ingressClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get ingress list", "err", err.Error())
			return err
		}
		remaining := ingressList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: ingressList.Continue}
				res, err := ingressClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get ingress list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = ingressList.Items
		p.output.Total = len(ingressList.Items)
		p.output.Resource = ingressList.Continue
	}
	return nil
}

func (svc *ingressService) GetIngressList(c context.Context, p GetIngressListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetIngressDetailsInputParams struct {
	NamespaceName string
	IngressName   string
	output        networkingv1.Ingress
}

func (p *GetIngressDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching ingress details of ....", p.NamespaceName)
	ingresssClient := GetKubeClientSet().NetworkingV1().Ingresses(p.NamespaceName)
	output, err := ingresssClient.Get(context.Background(), p.IngressName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get ingress ", p.IngressName, "err", err.Error())
		return err
	}
	p.output = *output
	return nil
}

func (svc *ingressService) GetIngressDetails(c context.Context, p GetIngressDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployIngressInputParams struct {
	Ingress *networkingv1.Ingress
	output  *networkingv1.Ingress
}

func (p *DeployIngressInputParams) Process(c context.Context) error {
	IngressClient := GetKubeClientSet().NetworkingV1().Ingresses(p.Ingress.Namespace)
	_, err := IngressClient.Get(context.Background(), p.Ingress.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating ingress in namespace "+p.Ingress.Namespace, "value", p.Ingress.Name)
		p.output, err = IngressClient.Create(context.Background(), p.Ingress, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create ingress in namespace "+p.Ingress.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("ingress created")
	} else {
		log.Logger.Infow("ingress exist in namespace "+p.Ingress.Namespace, "value", p.Ingress.Name)
		log.Logger.Infow("Updating ingress in namespace "+p.Ingress.Namespace, "value", p.Ingress.Name)
		p.output, err = IngressClient.Update(context.Background(), p.Ingress, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update ingress ", p.Ingress.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("ingress updated")
	}
	return nil
}

func (svc *ingressService) DeployIngress(c context.Context, p DeployIngressInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteIngressInputParams struct {
	NamespaceName string
	IngressName   string
}

func (p *DeleteIngressInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting Ingress of ....", p.NamespaceName)
	IngressClient := GetKubeClientSet().NetworkingV1().Ingresses(p.NamespaceName)
	_, err := IngressClient.Get(context.Background(), p.IngressName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get Ingress ", p.IngressName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = IngressClient.Delete(context.Background(), p.IngressName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete Ingress ", p.IngressName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *ingressService) DeleteIngress(c context.Context, p DeleteIngressInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
