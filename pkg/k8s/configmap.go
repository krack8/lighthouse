package k8s

import (
	"context"
	cfg "github.com/krack8/lighthouse/pkg/config"
	"github.com/krack8/lighthouse/pkg/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"strconv"
	"strings"
)

type ConfigMapServiceInterface interface {
	GetConfigMapList(c context.Context, p GetConfigMapListInputParams) (interface{}, error)
	GetConfigMapDetails(c context.Context, p GetConfigMapDetailsInputParams) (interface{}, error)
	DeployConfigMap(c context.Context, p DeployConfigMapInputParams) (interface{}, error)
	DeleteConfigMap(c context.Context, p DeleteConfigMapInputParams) (interface{}, error)
}

type configMapService struct{}

var cms configMapService

func ConfigMapService() *configMapService {
	return &cms
}

func getConfigMapClient(namespace string) v1.ConfigMapInterface {
	return cfg.GetKubeClientSet().CoreV1().ConfigMaps(namespace)
}

type OutputConfigMapList struct {
	Result    []corev1.ConfigMap
	Resource  string
	Remaining int64
	Total     int
}

type GetConfigMapListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputConfigMapList
}

func (p *GetConfigMapListInputParams) Find(c context.Context, configMapClient v1.ConfigMapInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "configmap")
	filteredconfigmaps := []corev1.ConfigMap{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		configmapList, err := configMapClient.List(context.Background(), listOptions)
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
		configmapList, err := configMapClient.List(context.Background(), listOptions)
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

func (p *GetConfigMapListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching config map list")
	configMapClient := cfg.GetKubeClientSet().CoreV1().ConfigMaps(p.NamespaceName)
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
	var configMapList *corev1.ConfigMapList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, configMapClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
			return err
		}
		return nil
	} else {
		configMapList, err = configMapClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
			return err
		}

		configMapList, err = configMapClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
			return err
		}
		remaining := configMapList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: configMapList.Continue}
				res, err := configMapClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get configmap list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = configMapList.Items
		p.output.Total = len(configMapList.Items)
		p.output.Resource = configMapList.Continue
	}
	return nil
}

func (svc *configMapService) GetConfigMapList(c context.Context, p GetConfigMapListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetConfigMapDetailsInputParams struct {
	NamespaceName string
	ConfigMapName string
	output        corev1.ConfigMap
	Client        v1.ConfigMapInterface
}

func (p *GetConfigMapDetailsInputParams) GetClient() v1.ConfigMapInterface {
	if p.Client != nil {
		return p.Client
	}
	return getConfigMapClient(p.NamespaceName)
}

func (p *GetConfigMapDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching configMap details of ....", p.NamespaceName)
	output, err := p.GetClient().Get(context.Background(), p.ConfigMapName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get configMap ", p.ConfigMapName, "err", err.Error())
		return err
	}
	p.output = *output
	return nil
}

func (svc *configMapService) GetConfigMapDetails(c context.Context, p GetConfigMapDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployConfigMapInputParams struct {
	ConfigMap *corev1.ConfigMap
	output    *corev1.ConfigMap
}

func (p *DeployConfigMapInputParams) Process(c context.Context) error {
	configMapClient := cfg.GetKubeClientSet().CoreV1().ConfigMaps(p.ConfigMap.Namespace)
	_, err := configMapClient.Get(context.Background(), p.ConfigMap.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating Config Map in namespace "+p.ConfigMap.Namespace, "value", p.ConfigMap.Name)
		p.output, err = configMapClient.Create(context.Background(), p.ConfigMap, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create config map in namespace "+p.ConfigMap.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("config map created")
	} else {
		log.Logger.Infow("config map exist in namespace "+p.ConfigMap.Namespace, "value", p.ConfigMap.Name)
		log.Logger.Infow("Updating config map in namespace "+p.ConfigMap.Namespace, "value", p.ConfigMap.Name)
		p.output, err = configMapClient.Update(context.Background(), p.ConfigMap, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update config map ", p.ConfigMap.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("config map updated")
	}
	return nil
}

func (svc *configMapService) DeployConfigMap(c context.Context, p DeployConfigMapInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteConfigMapInputParams struct {
	NamespaceName string
	ConfigMapName string
}

func (p *DeleteConfigMapInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting configMap of ....", p.NamespaceName)
	configMapClient := cfg.GetKubeClientSet().CoreV1().ConfigMaps(p.NamespaceName)
	_, err := configMapClient.Get(context.Background(), p.ConfigMapName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get configMap ", p.ConfigMapName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = configMapClient.Delete(context.Background(), p.ConfigMapName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete configMap ", p.ConfigMapName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *configMapService) DeleteConfigMap(c context.Context, p DeleteConfigMapInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
