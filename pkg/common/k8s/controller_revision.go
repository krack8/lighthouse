package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	_v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"strconv"
	"strings"
)

type ControllerRevisionServiceInterface interface {
	GetControllerRevisionList(c context.Context, p GetControllerRevisionListInputParams) (interface{}, error)
	GetControllerRevisionDetails(c context.Context, p GetControllerRevisionDetailsInputParams) (interface{}, error)
	DeployControllerRevision(c context.Context, p DeployControllerRevisionInputParams) (interface{}, error)
	DeleteControllerRevision(c context.Context, p DeleteControllerRevisionInputParams) (interface{}, error)
}

type controllerRevisionService struct{}

var ctrlrs controllerRevisionService

func ControllerRevisionService() *controllerRevisionService {
	return &ctrlrs
}

const (
	CONTROLLER_REVISION_API_VERSION = "apps/v1"
	CONTROLLER_REVISION_KIND        = "ControllerRevision"
)

type OutputControllerRevisionList struct {
	Result    []appsv1.ControllerRevision
	Resource  string
	Remaining int64
	Total     int
}

type GetControllerRevisionListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputControllerRevisionList
}

func (p *GetControllerRevisionListInputParams) Find(c context.Context, controllerRevisionClient _v1.ControllerRevisionInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "configmap")
	filteredconfigmaps := []appsv1.ControllerRevision{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		configmapList, err := controllerRevisionClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
			return err
		}

		for _, configmap := range configmapList.Items {
			if strings.Contains(configmap.Name, p.Search) {
				filteredconfigmaps = append(filteredconfigmaps, configmap)
			}
		}
		length = len(filteredconfigmaps)
		nextPageToken = configmapList.Continue
		if configmapList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		configmapList, err := controllerRevisionClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
			return err
		}
		for _, configmap := range configmapList.Items {
			if strings.Contains(configmap.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredconfigmaps
	p.output.Total = len(filteredconfigmaps)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetControllerRevisionListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching config map list")
	controllerRevisionClient := GetKubeClientSet().AppsV1().ControllerRevisions(p.NamespaceName)
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
	var controllerRevisionList *appsv1.ControllerRevisionList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, controllerRevisionClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
			return err
		}
		return nil
	} else {
		controllerRevisionList, err = controllerRevisionClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
			return err
		}

		controllerRevisionList, err = controllerRevisionClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
			return err
		}
		remaining := controllerRevisionList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: controllerRevisionList.Continue}
				res, err := controllerRevisionClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = controllerRevisionList.Items
		p.output.Total = len(controllerRevisionList.Items)
		p.output.Resource = controllerRevisionList.Continue
	}
	return nil
}

func (svc *controllerRevisionService) GetControllerRevisionList(c context.Context, p GetControllerRevisionListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetControllerRevisionDetailsInputParams struct {
	NamespaceName          string
	ControllerRevisionName string
	output                 appsv1.ControllerRevision
}

func (p *GetControllerRevisionDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching controllerRevision details of ....", p.NamespaceName)
	controllerRevisionClient := GetKubeClientSet().AppsV1().ControllerRevisions(p.NamespaceName)
	output, err := controllerRevisionClient.Get(context.Background(), p.ControllerRevisionName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get controllerRevision ", p.ControllerRevisionName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.APIVersion = CONTROLLER_REVISION_API_VERSION
	p.output.Kind = CONTROLLER_REVISION_KIND
	return nil
}

func (svc *controllerRevisionService) GetControllerRevisionDetails(c context.Context, p GetControllerRevisionDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployControllerRevisionInputParams struct {
	ControllerRevision *appsv1.ControllerRevision
	output             *appsv1.ControllerRevision
}

func (p *DeployControllerRevisionInputParams) Process(c context.Context) error {
	controllerRevisionClient := GetKubeClientSet().AppsV1().ControllerRevisions(p.ControllerRevision.Namespace)
	_, err := controllerRevisionClient.Get(context.Background(), p.ControllerRevision.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating Config Map in namespace "+p.ControllerRevision.Namespace, "value", p.ControllerRevision.Name)
		p.output, err = controllerRevisionClient.Create(context.Background(), p.ControllerRevision, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create config map in namespace "+p.ControllerRevision.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("config map created")
	} else {
		log.Logger.Infow("config map exist in namespace "+p.ControllerRevision.Namespace, "value", p.ControllerRevision.Name)
		log.Logger.Infow("Updating config map in namespace "+p.ControllerRevision.Namespace, "value", p.ControllerRevision.Name)
		p.output, err = controllerRevisionClient.Update(context.Background(), p.ControllerRevision, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update config map ", p.ControllerRevision.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("config map updated")
	}
	return nil
}

func (svc *controllerRevisionService) DeployControllerRevision(c context.Context, p DeployControllerRevisionInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteControllerRevisionInputParams struct {
	NamespaceName          string
	ControllerRevisionName string
}

func (p *DeleteControllerRevisionInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting controllerRevision of ....", p.NamespaceName)
	controllerRevisionClient := GetKubeClientSet().AppsV1().ControllerRevisions(p.NamespaceName)
	_, err := controllerRevisionClient.Get(context.Background(), p.ControllerRevisionName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get controllerRevision ", p.ControllerRevisionName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = controllerRevisionClient.Delete(context.Background(), p.ControllerRevisionName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete controllerRevision ", p.ControllerRevisionName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *controllerRevisionService) DeleteControllerRevision(c context.Context, p DeleteControllerRevisionInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
