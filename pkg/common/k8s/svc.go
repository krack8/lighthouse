package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"strconv"
	"strings"
)

type SvcServiceInterface interface {
	GetSvcList(c context.Context, p GetSvcListInputParams) (interface{}, error)
	GetSvcDetails(c context.Context, p GetSvcDetailsInputParams) (interface{}, error)
	DeploySvc(c context.Context, p DeploySvcInputParams) (interface{}, error)
	DeleteSvc(c context.Context, p DeleteSvcInputParams) (interface{}, error)
}

type svcService struct{}

var svcs svcService

func SvcService() *svcService {
	return &svcs
}

const (
	SVC_API_VERSION = "v1"
	SVC_KIND        = "Service"
)

type OutputSvcList struct {
	Result    []corev1.Service
	Resource  string
	Remaining int64
	Total     int
}

type GetSvcListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputSvcList
}

func (p *GetSvcListInputParams) Find(c context.Context, svcClient v1.ServiceInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "svc")
	filteredSvc := []corev1.Service{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		svcList, err := svcClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get svc list", "err", err.Error())
			return err
		}

		for _, svc := range svcList.Items {
			if strings.Contains(svc.Name, p.Search) {
				filteredSvc = append(filteredSvc, svc)
			}
		}
		length = len(filteredSvc)
		nextPageToken = svcList.Continue
		if svcList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		svcList, err := svcClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get svc list", "err", err.Error())
			return err
		}
		for _, svc := range svcList.Items {
			if strings.Contains(svc.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredSvc
	p.output.Total = len(filteredSvc)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetSvcListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching svc list")
	svcClient := GetKubeClientSet().CoreV1().Services(p.NamespaceName)
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
	var svcList *corev1.ServiceList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, svcClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get svc list", "err", err.Error())
			return err
		}
		return nil
	} else {
		svcList, err = svcClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get svc list", "err", err.Error())
			return err
		}
		remaining := svcList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: svcList.Continue}
				res, err := svcClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get svc list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = svcList.Items
		p.output.Total = len(svcList.Items)
		p.output.Resource = svcList.Continue
	}
	return nil
}

func (svc *svcService) GetSvcList(c context.Context, p GetSvcListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetSvcDetailsInputParams struct {
	NamespaceName string
	SvcName       string
	output        corev1.Service
}

func (p *GetSvcDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching svc details of ....", p.NamespaceName)
	svcClient := GetKubeClientSet().CoreV1().Services(p.NamespaceName)
	output, err := svcClient.Get(context.Background(), p.SvcName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get svc ", p.SvcName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.APIVersion = SVC_API_VERSION
	p.output.Kind = SVC_KIND
	return nil
}

func (svc *svcService) GetSvcDetails(c context.Context, p GetSvcDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeploySvcInputParams struct {
	Svc    *corev1.Service
	output *corev1.Service
}

func (p *DeploySvcInputParams) Process(c context.Context) error {
	svcClient := GetKubeClientSet().CoreV1().Services(p.Svc.Namespace)
	existingService, err := svcClient.Get(context.Background(), p.Svc.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating Service in namespace "+p.Svc.Namespace, "value", p.Svc.Name)
		p.output, err = svcClient.Create(context.Background(), p.Svc, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create c service in namespace "+p.Svc.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("svc created")
	} else {
		log.Logger.Infow("Service exist in namespace "+p.Svc.Namespace, "value", p.Svc.Name)
		log.Logger.Infow("Updating Service in namespace "+p.Svc.Namespace, "value", p.Svc.Name)
		p.Svc.ResourceVersion = existingService.ResourceVersion
		p.Svc.Spec.ClusterIP = existingService.Spec.ClusterIP
		p.output, err = svcClient.Update(context.Background(), p.Svc, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update svc ", p.Svc.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("svc updated")
	}
	return nil
}

func (svc *svcService) DeploySvc(c context.Context, p DeploySvcInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteSvcInputParams struct {
	NamespaceName string
	SvcName       string
}

func (p *DeleteSvcInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting svc of ....", p.NamespaceName)
	svcClient := GetKubeClientSet().CoreV1().Services(p.NamespaceName)
	_, err := svcClient.Get(context.Background(), p.SvcName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get svc ", p.SvcName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = svcClient.Delete(context.Background(), p.SvcName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete svc ", p.SvcName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *svcService) DeleteSvc(c context.Context, p DeleteSvcInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
