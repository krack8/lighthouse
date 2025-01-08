package k8s

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"strconv"

	cfg "github.com/krack8/lighthouse/pkg/config"
	"github.com/krack8/lighthouse/pkg/log"
	"strings"
)

type ResourceQuotaServiceInterface interface {
	GetResourceQuotaList(c context.Context, p GetResourceQuotaListInputParams) (interface{}, error)
	GetResourceQuotaDetails(c context.Context, p GetResourceQuotaDetailsInputParams) (interface{}, error)
	DeployResourceQuota(c context.Context, p DeployResourceQuotaInputParams) (interface{}, error)
	DeleteResourceQuota(c context.Context, p DeleteResourceQuotaInputParams) (interface{}, error)
}

type resourceQuotaService struct{}

var rqs resourceQuotaService

func ResourceQuotaService() *resourceQuotaService {
	return &rqs
}

type OutputResourceQuotaList struct {
	Result    []corev1.ResourceQuota
	Resource  string
	Remaining int64
	Total     int
}

type GetResourceQuotaListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputResourceQuotaList
}

func (p *GetResourceQuotaListInputParams) Find(c context.Context, resourceQuotaClient v1.ResourceQuotaInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "resourceQuota")
	filteredResourceQuota := []corev1.ResourceQuota{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		resourceQuotaList, err := resourceQuotaClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get resourceQuota list", "err", err.Error())
			return err
		}

		for _, resourceQuota := range resourceQuotaList.Items {
			if strings.Contains(resourceQuota.Name, p.Search) {
				filteredResourceQuota = append(filteredResourceQuota, resourceQuota)
			}
		}
		length = len(filteredResourceQuota)
		nextPageToken = resourceQuotaList.Continue
		if resourceQuotaList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		resourceQuotaList, err := resourceQuotaClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get resourceQuota list", "err", err.Error())
			return err
		}
		for _, resourceQuota := range resourceQuotaList.Items {
			if strings.Contains(resourceQuota.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredResourceQuota
	p.output.Total = len(filteredResourceQuota)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetResourceQuotaListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching resourceQuota list")
	resourceQuotaClient := cfg.GetKubeClientSet().CoreV1().ResourceQuotas(p.NamespaceName)
	limit := cfg.PageLimit
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
	var resourceQuotaList *corev1.ResourceQuotaList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, resourceQuotaClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get resourceQuota list", "err", err.Error())
			return err
		}
		return nil
	} else {
		resourceQuotaList, err = resourceQuotaClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get resourceQuota list", "err", err.Error())
			return err
		}

		resourceQuotaList, err = resourceQuotaClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get resourceQuota list", "err", err.Error())
			return err
		}
		remaining := resourceQuotaList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: resourceQuotaList.Continue}
				res, err := resourceQuotaClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get resourceQuota list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = resourceQuotaList.Items
		p.output.Total = len(resourceQuotaList.Items)
		p.output.Resource = resourceQuotaList.Continue
	}
	return nil
}

func (resourceQuota *resourceQuotaService) GetResourceQuotaList(c context.Context, p GetResourceQuotaListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetResourceQuotaDetailsInputParams struct {
	NamespaceName     string
	ResourceQuotaName string
	output            corev1.ResourceQuota
}

func (p *GetResourceQuotaDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching resourceQuota details of ....", p.NamespaceName)
	resourceQuotaClient := cfg.GetKubeClientSet().CoreV1().ResourceQuotas(p.NamespaceName)
	output, err := resourceQuotaClient.Get(context.Background(), p.ResourceQuotaName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get resourceQuota ", p.ResourceQuotaName, "err", err.Error())
		return err
	}
	p.output = *output
	return nil
}

func (resourceQuota *resourceQuotaService) GetResourceQuotaDetails(c context.Context, p GetResourceQuotaDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployResourceQuotaInputParams struct {
	ResourceQuota *corev1.ResourceQuota
	output        *corev1.ResourceQuota
}

func (p *DeployResourceQuotaInputParams) Process(c context.Context) error {
	resourceQuotaClient := cfg.GetKubeClientSet().CoreV1().ResourceQuotas(p.ResourceQuota.Namespace)
	_, err := resourceQuotaClient.Get(context.Background(), p.ResourceQuota.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating resource quota in namespace "+p.ResourceQuota.Namespace, "value", p.ResourceQuota.Name)
		p.output, err = resourceQuotaClient.Create(context.Background(), p.ResourceQuota, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create resource quota in namespace "+p.ResourceQuota.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("resourceQuota created")
	} else {
		log.Logger.Infow("Resource Quota exist in namespace "+p.ResourceQuota.Namespace, "value", p.ResourceQuota.Name)
		log.Logger.Infow("Updating Resource Quota Service in namespace "+p.ResourceQuota.Namespace, "value", p.ResourceQuota.Name)
		p.output, err = resourceQuotaClient.Update(context.Background(), p.ResourceQuota, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update resourceQuota ", p.ResourceQuota.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("resourceQuota updated")
	}
	return nil
}

func (resourceQuota *resourceQuotaService) DeployResourceQuota(c context.Context, p DeployResourceQuotaInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteResourceQuotaInputParams struct {
	NamespaceName     string
	ResourceQuotaName string
}

func (p *DeleteResourceQuotaInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting resourceQuota of ....", p.NamespaceName)
	resourceQuotaClient := cfg.GetKubeClientSet().CoreV1().ResourceQuotas(p.NamespaceName)
	_, err := resourceQuotaClient.Get(context.Background(), p.ResourceQuotaName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get resourceQuota ", p.ResourceQuotaName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = resourceQuotaClient.Delete(context.Background(), p.ResourceQuotaName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete resourceQuota ", p.ResourceQuotaName, "err", err.Error())
		return err
	}
	return nil
}

func (resourceQuota *resourceQuotaService) DeleteResourceQuota(c context.Context, p DeleteResourceQuotaInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
