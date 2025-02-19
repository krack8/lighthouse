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

type ReplicationControllerServiceInterface interface {
	GetReplicationControllerList(c context.Context, p GetReplicationControllerListInputParams) (interface{}, error)
	GetReplicationControllerDetails(c context.Context, p GetReplicationControllerDetailsInputParams) (interface{}, error)
	DeployReplicationController(c context.Context, p DeployReplicationControllerInputParams) (interface{}, error)
	DeleteReplicationController(c context.Context, p DeleteReplicationControllerInputParams) (interface{}, error)
}

type replicationControllerService struct{}

var rcs replicationControllerService

func ReplicationControllerService() *replicationControllerService {
	return &rcs
}

type OutputReplicationControllerList struct {
	Result    []corev1.ReplicationController
	Resource  string
	Remaining int64
	Total     int
}

type GetReplicationControllerListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputReplicationControllerList
}

func (p *GetReplicationControllerListInputParams) Find(c context.Context, replicationControllerClient v1.ReplicationControllerInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "configmap")
	filteredconfigmaps := []corev1.ReplicationController{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		configmapList, err := replicationControllerClient.List(context.Background(), listOptions)
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
		configmapList, err := replicationControllerClient.List(context.Background(), listOptions)
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

func (p *GetReplicationControllerListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching config map list")
	replicationControllerClient := config.GetKubeClientSet().CoreV1().ReplicationControllers(p.NamespaceName)
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
	var replicationControllerList *corev1.ReplicationControllerList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, replicationControllerClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
			return err
		}
		return nil
	} else {
		replicationControllerList, err = replicationControllerClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
			return err
		}

		replicationControllerList, err = replicationControllerClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
			return err
		}
		remaining := replicationControllerList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: replicationControllerList.Continue}
				res, err := replicationControllerClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = replicationControllerList.Items
		p.output.Total = len(replicationControllerList.Items)
		p.output.Resource = replicationControllerList.Continue
	}
	return nil
}

func (svc *replicationControllerService) GetReplicationControllerList(c context.Context, p GetReplicationControllerListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetReplicationControllerDetailsInputParams struct {
	NamespaceName             string
	ReplicationControllerName string
	output                    corev1.ReplicationController
}

func (p *GetReplicationControllerDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching replicationController details of ....", p.NamespaceName)
	replicationControllerClient := config.GetKubeClientSet().CoreV1().ReplicationControllers(p.NamespaceName)
	output, err := replicationControllerClient.Get(context.Background(), p.ReplicationControllerName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get replicationController ", p.ReplicationControllerName, "err", err.Error())
		return err
	}
	p.output = *output
	return nil
}

func (svc *replicationControllerService) GetReplicationControllerDetails(c context.Context, p GetReplicationControllerDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployReplicationControllerInputParams struct {
	ReplicationController *corev1.ReplicationController
	output                *corev1.ReplicationController
}

func (p *DeployReplicationControllerInputParams) Process(c context.Context) error {
	replicationControllerClient := config.GetKubeClientSet().CoreV1().ReplicationControllers(p.ReplicationController.Namespace)
	_, err := replicationControllerClient.Get(context.Background(), p.ReplicationController.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating Config Map in namespace "+p.ReplicationController.Namespace, "value", p.ReplicationController.Name)
		p.output, err = replicationControllerClient.Create(context.Background(), p.ReplicationController, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create config map in namespace "+p.ReplicationController.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("config map created")
	} else {
		log.Logger.Infow("config map exist in namespace "+p.ReplicationController.Namespace, "value", p.ReplicationController.Name)
		log.Logger.Infow("Updating config map in namespace "+p.ReplicationController.Namespace, "value", p.ReplicationController.Name)
		p.output, err = replicationControllerClient.Update(context.Background(), p.ReplicationController, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update config map ", p.ReplicationController.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("config map updated")
	}
	return nil
}

func (svc *replicationControllerService) DeployReplicationController(c context.Context, p DeployReplicationControllerInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteReplicationControllerInputParams struct {
	NamespaceName             string
	ReplicationControllerName string
}

func (p *DeleteReplicationControllerInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting replicationController of ....", p.NamespaceName)
	replicationControllerClient := config.GetKubeClientSet().CoreV1().ReplicationControllers(p.NamespaceName)
	_, err := replicationControllerClient.Get(context.Background(), p.ReplicationControllerName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get replicationController ", p.ReplicationControllerName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = replicationControllerClient.Delete(context.Background(), p.ReplicationControllerName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete replicationController ", p.ReplicationControllerName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *replicationControllerService) DeleteReplicationController(c context.Context, p DeleteReplicationControllerInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
