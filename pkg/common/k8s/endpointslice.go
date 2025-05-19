package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	"k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	_v1 "k8s.io/client-go/kubernetes/typed/discovery/v1"
	"strconv"
	"strings"
)

type EndpointSliceServiceInterface interface {
	GetEndpointSliceList(c context.Context, p GetEndpointSliceListInputParams) (interface{}, error)
	GetEndpointSliceDetails(c context.Context, p GetEndpointSliceDetailsInputParams) (interface{}, error)
	DeployEndpointSlice(c context.Context, p DeployEndpointSliceInputParams) (interface{}, error)
	DeleteEndpointSlice(c context.Context, p DeleteEndpointSliceInputParams) (interface{}, error)
}

type endpointSliceService struct{}

var crs endpointSliceService

func EndpointSliceService() *endpointSliceService {
	return &crs
}

const (
	EndpointslicesApiVersion = "discovery.k8s.io/v1"
	EndpointslicesKind       = "EndpointSlice"
)

type OutputEndpointSliceList struct {
	Result    []v1.EndpointSlice
	Resource  string
	Remaining int64
	Total     int
}

type GetEndpointSliceListInputParams struct {
	Search    string
	Labels    map[string]string
	Namespace string
	Limit     string
	Continue  string
	output    OutputEndpointSliceList
}

func (p *GetEndpointSliceListInputParams) Find(c context.Context, endpointSliceClient _v1.EndpointSliceInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "endpointSlice")
	filteredEndpointSlices := []v1.EndpointSlice{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		endpointSliceList, err := endpointSliceClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get endpointSlice list", "err", err.Error())
			return err
		}

		for _, endpointSlice := range endpointSliceList.Items {
			if strings.Contains(endpointSlice.Name, p.Search) {
				filteredEndpointSlices = append(filteredEndpointSlices, endpointSlice)
			}
		}
		length = len(filteredEndpointSlices)
		nextPageToken = endpointSliceList.Continue
		if endpointSliceList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		endpointSliceList, err := endpointSliceClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get endpointSlice list", "err", err.Error())
			return err
		}
		for _, endpointSlice := range endpointSliceList.Items {
			if strings.Contains(endpointSlice.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredEndpointSlices
	p.output.Total = len(filteredEndpointSlices)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetEndpointSliceListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching cluster role list")
	endpointSliceClient := GetKubeClientSet().DiscoveryV1().EndpointSlices(p.Namespace)
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
	var endpointSliceList *v1.EndpointSliceList
	if p.Search != "" {
		err = p.Find(c, endpointSliceClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get cluster role list", "err", err.Error())
			return err
		}
		return nil
	} else {
		endpointSliceList, err = endpointSliceClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to cluster role list", "err", err.Error())
			return err
		}

		endpointSliceList, err = endpointSliceClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get cluster role list", "err", err.Error())
			return err
		}
		remaining := endpointSliceList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: endpointSliceList.Continue}
				res, err := endpointSliceClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get cluster role list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = endpointSliceList.Items
		p.output.Total = len(endpointSliceList.Items)
		p.output.Resource = endpointSliceList.Continue
	}
	return nil
}

func (p *GetEndpointSliceListInputParams) PostProcess(ctx context.Context) error {
	for i := 0; i < len(p.output.Result); i++ {
		p.output.Result[i].ManagedFields = nil
		p.output.Result[i].APIVersion = EndpointsApiVersion
		p.output.Result[i].Kind = EndpointslicesKind
	}
	return nil
}

func (svc *endpointSliceService) GetEndpointSliceList(c context.Context, p GetEndpointSliceListInputParams) (interface{}, error) {
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

type GetEndpointSliceDetailsInputParams struct {
	EndpointSliceName string
	Namespace         string
	output            v1.EndpointSlice
}

func (p *GetEndpointSliceDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching cluster role details ....")
	endpointSlicesClient := GetKubeClientSet().DiscoveryV1().EndpointSlices(p.Namespace)
	output, err := endpointSlicesClient.Get(context.Background(), p.EndpointSliceName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get cluster role ", p.EndpointSliceName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.ManagedFields = nil
	p.output.APIVersion = EndpointslicesApiVersion
	p.output.Kind = EndpointslicesKind
	return nil
}

func (svc *endpointSliceService) GetEndpointSliceDetails(c context.Context, p GetEndpointSliceDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployEndpointSliceInputParams struct {
	EndpointSlice *v1.EndpointSlice
	Namespace     string
	output        *v1.EndpointSlice
}

func (p *DeployEndpointSliceInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *DeployEndpointSliceInputParams) Process(c context.Context) error {
	if p.Namespace == "" {
		p.Namespace = "default"
	}
	endpointSliceClient := GetKubeClientSet().DiscoveryV1().EndpointSlices(p.Namespace)
	returnedEndpointSlice, err := endpointSliceClient.Get(context.Background(), p.EndpointSlice.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating endpointSlice ", "value", p.EndpointSlice.Name)
		p.output, err = endpointSliceClient.Create(context.Background(), p.EndpointSlice, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create endpointSlice ", "err", err.Error())
			return err
		}
		log.Logger.Infow("endpointSlice created")
	} else {
		log.Logger.Infow("EndpointSlice exist ", "value", p.EndpointSlice.Name)
		log.Logger.Infow("Updating endpointSlice ", "value", p.EndpointSlice.Name)
		p.EndpointSlice.SetResourceVersion(returnedEndpointSlice.ResourceVersion)
		p.output, err = endpointSliceClient.Update(context.Background(), p.EndpointSlice, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update endpointSlice ", p.EndpointSlice.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("endpointSlice updated")
	}
	return nil
}

func (svc *endpointSliceService) DeployEndpointSlice(c context.Context, p DeployEndpointSliceInputParams) (interface{}, error) {
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

type DeleteEndpointSliceInputParams struct {
	EndpointSliceName string
	Namespace         string
}

func (p *DeleteEndpointSliceInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting endpointSlice ....", p.EndpointSliceName)
	endpointSliceClient := GetKubeClientSet().DiscoveryV1().EndpointSlices(p.Namespace)
	_, err := endpointSliceClient.Get(context.Background(), p.EndpointSliceName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get endpointSlice ", p.EndpointSliceName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = endpointSliceClient.Delete(context.Background(), p.EndpointSliceName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete endpointSlice ", p.EndpointSliceName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *endpointSliceService) DeleteEndpointSlice(c context.Context, p DeleteEndpointSliceInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
