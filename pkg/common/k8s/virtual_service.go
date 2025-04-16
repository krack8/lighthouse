package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
	_v1beta1 "istio.io/client-go/pkg/clientset/versioned/typed/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"strconv"
	"strings"
)

type VirtualServiceServiceInterface interface {
	GetVirtualServiceList(c context.Context, p GetVirtualServiceListInputParams) (interface{}, error)
	GetVirtualServiceDetails(c context.Context, p GetVirtualServiceDetailsInputParams) (interface{}, error)
	DeployVirtualService(c context.Context, p DeployVirtualServiceInputParams) (interface{}, error)
	DeleteVirtualService(c context.Context, p DeleteVirtualServiceInputParams) (interface{}, error)
}

type virtualServiceService struct{}

var is virtualServiceService

func VirtualServiceService() *virtualServiceService {
	return &is
}

const (
	VIRTUAL_SERVICE_API_VERSION = "v1beta1"
	VIRTUAL_SERVICE_KIND        = "VirtualService"
)

type OutputVirtualServiceList struct {
	Result    []*v1beta1.VirtualService
	Resource  string
	Remaining int64
	Total     int
}

type GetVirtualServiceListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputVirtualServiceList
}

func (p *GetVirtualServiceListInputParams) Find(c context.Context, virtualServiceClient _v1beta1.VirtualServiceInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "virtualService")
	filteredVirtualService := []*v1beta1.VirtualService{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		virtualServiceList, err := virtualServiceClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get virtualService list", "err", err.Error())
			return err
		}

		for _, virtualService := range virtualServiceList.Items {
			if strings.Contains(virtualService.Name, p.Search) {
				filteredVirtualService = append(filteredVirtualService, virtualService)
			}
		}
		length = len(filteredVirtualService)
		nextPageToken = virtualServiceList.Continue
		if virtualServiceList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		virtualServiceList, err := virtualServiceClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get virtualService list", "err", err.Error())
			return err
		}
		for _, virtualService := range virtualServiceList.Items {
			if strings.Contains(virtualService.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredVirtualService
	p.output.Total = len(filteredVirtualService)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetVirtualServiceListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching virtualService list")
	virtualServiceClient := GetNetworkingV1Beta1ClientSet().VirtualServices(p.NamespaceName)
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
	var virtualServiceList *v1beta1.VirtualServiceList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, virtualServiceClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get virtualService list", "err", err.Error())
			return err
		}
		return nil
	} else {
		virtualServiceList, err = virtualServiceClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get virtualService list", "err", err.Error())
			return err
		}

		virtualServiceList, err = virtualServiceClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get virtualService list", "err", err.Error())
			return err
		}
		remaining := virtualServiceList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: virtualServiceList.Continue}
				res, err := virtualServiceClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get virtualService list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = virtualServiceList.Items
		p.output.Total = len(virtualServiceList.Items)
		p.output.Resource = virtualServiceList.Continue
	}
	return nil
}

func (svc *virtualServiceService) GetVirtualServiceList(c context.Context, p GetVirtualServiceListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetVirtualServiceDetailsInputParams struct {
	NamespaceName      string
	VirtualServiceName string
	output             v1beta1.VirtualService
}

func (p *GetVirtualServiceDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching virtualService details of ....", p.NamespaceName)
	virtualServicesClient := GetNetworkingV1Beta1ClientSet().VirtualServices(p.NamespaceName)
	output, err := virtualServicesClient.Get(context.Background(), p.VirtualServiceName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get virtualService ", p.VirtualServiceName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.APIVersion = VIRTUAL_SERVICE_API_VERSION
	p.output.Kind = VIRTUAL_SERVICE_KIND
	return nil
}

func (svc *virtualServiceService) GetVirtualServiceDetails(c context.Context, p GetVirtualServiceDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployVirtualServiceInputParams struct {
	VirtualService *v1beta1.VirtualService
	output         *v1beta1.VirtualService
}

func (p *DeployVirtualServiceInputParams) Process(c context.Context) error {
	VirtualServiceClient := GetNetworkingV1Beta1ClientSet().VirtualServices(p.VirtualService.Namespace)
	_, err := VirtualServiceClient.Get(context.Background(), p.VirtualService.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating virtualService in namespace "+p.VirtualService.Namespace, "value", p.VirtualService.Name)
		p.output, err = VirtualServiceClient.Create(context.Background(), p.VirtualService, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create virtualService in namespace "+p.VirtualService.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("virtualService created")
	} else {
		log.Logger.Infow("virtualService exist in namespace "+p.VirtualService.Namespace, "value", p.VirtualService.Name)
		log.Logger.Infow("Updating virtualService in namespace "+p.VirtualService.Namespace, "value", p.VirtualService.Name)
		p.output, err = VirtualServiceClient.Update(context.Background(), p.VirtualService, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update virtualService ", p.VirtualService.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("virtualService updated")
	}
	return nil
}

func (svc *virtualServiceService) DeployVirtualService(c context.Context, p DeployVirtualServiceInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteVirtualServiceInputParams struct {
	NamespaceName      string
	VirtualServiceName string
}

func (p *DeleteVirtualServiceInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting VirtualService of ....", p.NamespaceName)
	VirtualServiceClient := GetNetworkingV1Beta1ClientSet().VirtualServices(p.NamespaceName)
	_, err := VirtualServiceClient.Get(context.Background(), p.VirtualServiceName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get VirtualService ", p.VirtualServiceName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = VirtualServiceClient.Delete(context.Background(), p.VirtualServiceName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete VirtualService ", p.VirtualServiceName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *virtualServiceService) DeleteVirtualService(c context.Context, p DeleteVirtualServiceInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
