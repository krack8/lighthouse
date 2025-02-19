package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	_v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"strconv"
	"strings"
)

type EndpointsServiceInterface interface {
	GetEndpointsList(c context.Context, p GetEndpointsListInputParams) (interface{}, error)
	GetEndpointsDetails(c context.Context, p GetEndpointsDetailsInputParams) (interface{}, error)
	DeployEndpoints(c context.Context, p DeployEndpointsInputParams) (interface{}, error)
	DeleteEndpoints(c context.Context, p DeleteEndpointsInputParams) (interface{}, error)
}

type endpointsService struct{}

var eps endpointsService

func EndpointsService() *endpointsService {
	return &eps
}

type OutputEndpointsList struct {
	Result    []v1.Endpoints
	Resource  string
	Remaining int64
	Total     int
}

type GetEndpointsListInputParams struct {
	Search    string
	Labels    map[string]string
	Namespace string
	Continue  string
	Limit     string
	output    OutputEndpointsList
}

func (p *GetEndpointsListInputParams) Find(c context.Context, endpointsClient _v1.EndpointsInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "endpoints")
	filteredEndpointss := []v1.Endpoints{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		endpointsList, err := endpointsClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get endpoints list", "err", err.Error())
			return err
		}

		for _, endpoints := range endpointsList.Items {
			if strings.Contains(endpoints.Name, p.Search) {
				filteredEndpointss = append(filteredEndpointss, endpoints)
			}
		}
		length = len(filteredEndpointss)
		nextPageToken = endpointsList.Continue
		if endpointsList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		endpointsList, err := endpointsClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get endpoints list", "err", err.Error())
			return err
		}
		for _, endpoints := range endpointsList.Items {
			if strings.Contains(endpoints.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredEndpointss
	p.output.Total = len(filteredEndpointss)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetEndpointsListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching cluster role binding list")
	endpointsClient := GetKubeClientSet().CoreV1().Endpoints(p.Namespace)
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
	var endpointsList *v1.EndpointsList
	if p.Search != "" {
		err = p.Find(c, endpointsClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get cluster role list", "err", err.Error())
			return err
		}
		return nil
	} else {
		endpointsList, err = endpointsClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to cluster role list", "err", err.Error())
			return err
		}

		endpointsList, err = endpointsClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get cluster role list", "err", err.Error())
			return err
		}
		remaining := endpointsList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: endpointsList.Continue}
				res, err := endpointsClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get cluster role list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = endpointsList.Items
		p.output.Total = len(endpointsList.Items)
		p.output.Resource = endpointsList.Continue
	}
	return nil
}

func (svc *endpointsService) GetEndpointsList(c context.Context, p GetEndpointsListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetEndpointsDetailsInputParams struct {
	EndpointsName string
	Namespace     string
	output        v1.Endpoints
}

func (p *GetEndpointsDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching cluster role details ....")
	endpointsClient := GetKubeClientSet().CoreV1().Endpoints(p.Namespace)
	output, err := endpointsClient.Get(context.Background(), p.EndpointsName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get cluster role binding", p.EndpointsName, "err", err.Error())
		return err
	}
	p.output = *output
	return nil
}

func (svc *endpointsService) GetEndpointsDetails(c context.Context, p GetEndpointsDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployEndpointsInputParams struct {
	Endpoints *v1.Endpoints
	Namespace string
	output    *v1.Endpoints
}

func (p *DeployEndpointsInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *DeployEndpointsInputParams) Process(c context.Context) error {
	if p.Namespace == "" {
		p.Namespace = "default"
	}
	endpointsClient := GetKubeClientSet().CoreV1().Endpoints(p.Namespace)
	_, err := endpointsClient.Get(context.Background(), p.Endpoints.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating endpoints ", "value", p.Endpoints.Name)
		p.output, err = endpointsClient.Create(context.Background(), p.Endpoints, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create endpoints ", "err", err.Error())
			return err
		}
		log.Logger.Infow("endpoints created")
	} else {
		log.Logger.Infow("Endpoints exist ", "value", p.Endpoints.Name)
		log.Logger.Infow("Updating endpoints ", "value", p.Endpoints.Name)
		p.output, err = endpointsClient.Update(context.Background(), p.Endpoints, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update endpoints ", p.Endpoints.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("endpoints updated")
	}
	return nil
}

func (svc *endpointsService) DeployEndpoints(c context.Context, p DeployEndpointsInputParams) (interface{}, error) {
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

type DeleteEndpointsInputParams struct {
	EndpointsName string
	Namespace     string
}

func (p *DeleteEndpointsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting endpoints ....", p.EndpointsName)
	endpointsClient := GetKubeClientSet().CoreV1().Endpoints(p.Namespace)
	_, err := endpointsClient.Get(context.Background(), p.EndpointsName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get endpoints ", p.EndpointsName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = endpointsClient.Delete(context.Background(), p.EndpointsName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete endpoints ", p.EndpointsName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *endpointsService) DeleteEndpoints(c context.Context, p DeleteEndpointsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
