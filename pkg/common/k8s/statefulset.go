package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"strconv"
	"strings"
)

type StatefulSetServiceInterface interface {
	GetStatefulSetList(c context.Context, p GetStatefulSetListInputParams) (interface{}, error)
	GetStatefulSetDetails(c context.Context, p GetStatefulSetDetailsInputParams) (interface{}, error)
	DeployStatefulSet(c context.Context, p DeployStatefulSetInputParams) (interface{}, error)
	DeleteStatefulSet(c context.Context, p DeleteStatefulSetInputParams) (interface{}, error)
	GetStatefulSetStats(c context.Context, p GetStatefulSetStatsInputParams) (interface{}, error)
	GetStatefulSetPodList(c context.Context, p GetStatefulSetPodListInputParams) (interface{}, error)
}

type statefulSetService struct{}

var sfss statefulSetService

func StatefulSetService() *statefulSetService {
	return &sfss
}

const (
	StatefulSetApiVersion = "apps/v1"
	StatefulSetKind       = "StatefulSet"
	PodLabelKey           = "controller-revision-hash"
)

type OutputStatefulSetList struct {
	Result    []appsv1.StatefulSet
	Resource  string
	Remaining int64
	Total     int
}

type GetStatefulSetListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputStatefulSetList
}

func (p *GetStatefulSetListInputParams) PostProcess(c context.Context) error {
	for idx, _ := range p.output.Result {
		p.output.Result[idx].ManagedFields = nil
		p.output.Result[idx].APIVersion = StatefulSetApiVersion
		p.output.Result[idx].Kind = StatefulSetKind
	}
	return nil
}

func (p *GetStatefulSetListInputParams) Find(c context.Context, statefulsetClient v1.StatefulSetInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "statefulset")
	filteredStatefulsets := []appsv1.StatefulSet{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		statefulsetList, err := statefulsetClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get statefulset list", "err", err.Error())
			return err
		}

		for _, statefulset := range statefulsetList.Items {
			if strings.Contains(statefulset.Name, p.Search) {
				filteredStatefulsets = append(filteredStatefulsets, statefulset)
			}
		}
		length = len(filteredStatefulsets)
		nextPageToken = statefulsetList.Continue
		if statefulsetList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		statefulsetList, err := statefulsetClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get statefulset list", "err", err.Error())
			return err
		}
		for _, statefulset := range statefulsetList.Items {
			if strings.Contains(statefulset.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredStatefulsets
	p.output.Total = len(filteredStatefulsets)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetStatefulSetListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching stateful set list")
	statefulSetClient := GetKubeClientSet().AppsV1().StatefulSets(p.NamespaceName)
	limit := config.PageLimit
	if p.Limit != "" {
		limit, _ = strconv.ParseInt(p.Limit, 10, 64)
	}
	listOptions := metav1.ListOptions{Limit: limit, Continue: p.Continue}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions.LabelSelector = labels.Set(labelSelector.MatchLabels).String()
	}
	var err error
	var statefulSetList *appsv1.StatefulSetList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, statefulSetClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get statefulset list", "err", err.Error())
			return err
		}
		return nil
	} else {
		statefulSetList, err = statefulSetClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get statefulset list", "err", err.Error())
			return err
		}

		statefulSetList, err = statefulSetClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get statefulset list", "err", err.Error())
			return err
		}
		remaining := statefulSetList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: statefulSetList.Continue}
				res, err := statefulSetClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get statefulset list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = statefulSetList.Items
		p.output.Total = len(statefulSetList.Items)
		p.output.Resource = statefulSetList.Continue
	}
	return nil
}

func (svc *statefulSetService) GetStatefulSetList(c context.Context, p GetStatefulSetListInputParams) (interface{}, error) {
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

type GetStatefulSetDetailsInputParams struct {
	NamespaceName   string
	StatefulSetName string
	output          appsv1.StatefulSet
}

func (p *GetStatefulSetDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching statefulSet details of ....", p.NamespaceName)
	statefulSetsClient := GetKubeClientSet().AppsV1().StatefulSets(p.NamespaceName)
	output, err := statefulSetsClient.Get(context.Background(), p.StatefulSetName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get statefulSet ", p.StatefulSetName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.ManagedFields = nil
	p.output.APIVersion = StatefulSetApiVersion
	p.output.Kind = StatefulSetKind
	return nil
}

func (svc *statefulSetService) GetStatefulSetDetails(c context.Context, p GetStatefulSetDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployStatefulSetInputParams struct {
	StatefulSet *appsv1.StatefulSet
	output      *appsv1.StatefulSet
}

func (p *DeployStatefulSetInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *DeployStatefulSetInputParams) Process(c context.Context) error {
	statefulSetClient := GetKubeClientSet().AppsV1().StatefulSets(p.StatefulSet.Namespace)
	_, err := statefulSetClient.Get(context.Background(), p.StatefulSet.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating statefulSet in namespace "+p.StatefulSet.Namespace, "value", p.StatefulSet.Name)
		p.output, err = statefulSetClient.Create(context.Background(), p.StatefulSet, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create statefulSet in namespace "+p.StatefulSet.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("statefulSet created")
	} else {
		log.Logger.Infow("StatefulSet exist in namespace "+p.StatefulSet.Namespace, "value", p.StatefulSet.Name)
		log.Logger.Infow("Updating statefulSet in namespace "+p.StatefulSet.Namespace, "value", p.StatefulSet.Name)
		p.output, err = statefulSetClient.Update(context.Background(), p.StatefulSet, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update statefulSet ", p.StatefulSet.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("statefulSet updated")
	}
	return nil
}

func (svc *statefulSetService) DeployStatefulSet(c context.Context, p DeployStatefulSetInputParams) (interface{}, error) {
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

type DeleteStatefulSetInputParams struct {
	NamespaceName   string
	StatefulSetName string
}

func (p *DeleteStatefulSetInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting statefulSet of ....", p.NamespaceName)
	statefulSetClient := GetKubeClientSet().AppsV1().StatefulSets(p.NamespaceName)
	_, err := statefulSetClient.Get(context.Background(), p.StatefulSetName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get statefulSet ", p.StatefulSetName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = statefulSetClient.Delete(context.Background(), p.StatefulSetName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete statefulSet ", p.StatefulSetName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *statefulSetService) DeleteStatefulSet(c context.Context, p DeleteStatefulSetInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}

type Stats struct {
	Total    int
	Ready    int
	NotReady int
}

func (s *Stats) New() *Stats {
	return &Stats{Total: 0, Ready: 0, NotReady: 0}
}

type GetStatefulSetStatsInputParams struct {
	NamespaceName string
	Labels        map[string]string
	Search        string
	output        *Stats
}

func (p *GetStatefulSetStatsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching statefulSet list stats")
	statefulSetClient := GetKubeClientSet().AppsV1().StatefulSets(p.NamespaceName)

	listOptions := metav1.ListOptions{}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions.LabelSelector = labels.Set(labelSelector.MatchLabels).String()
	}

	statefulSetList, err := statefulSetClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to get statefulSet list stats", "err", err.Error())
		return err
	}

	if p.Search != "" {
		listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		filteredStatefulSet := []appsv1.StatefulSet{}

		for _, statefulSet := range statefulSetList.Items {
			if strings.Contains(statefulSet.Name, p.Search) {
				filteredStatefulSet = append(filteredStatefulSet, statefulSet)
			}
		}

		p.output = p.output.New()
		for _, obj := range filteredStatefulSet {
			p.output.Total += int(obj.Status.Replicas)
			p.output.Ready += int(obj.Status.ReadyReplicas)
		}

		p.output.NotReady = p.output.Total - p.output.Ready
		return nil
	}

	p.output = p.output.New()

	for _, obj := range statefulSetList.Items {
		p.output.Total += int(obj.Status.Replicas)
		p.output.Ready += int(obj.Status.ReadyReplicas)
	}

	p.output.NotReady = p.output.Total - p.output.Ready
	return nil
}

func (svc *statefulSetService) GetStatefulSetStats(c context.Context, p GetStatefulSetStatsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type StatefulSetPodOutput struct {
	PodList   []corev1.Pod
	Resource  string
	Remaining int64
}

type GetStatefulSetPodListInputParams struct {
	NamespaceName   string
	StatefulSetName string
	Labels          map[string]string
	output          StatefulSetPodOutput
}

func (p *GetStatefulSetPodListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching statefulset pods list of ...."+p.StatefulSetName, "service", "statefulSet-pod-list")
	statefulSetsClient := GetKubeClientSet().AppsV1().StatefulSets(p.NamespaceName)
	statefulSet, err := statefulSetsClient.Get(context.Background(), p.StatefulSetName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get statefulSet ", p.StatefulSetName, "err", err.Error())
		return err
	}
	podClient := GetKubeClientSet().CoreV1().Pods(p.NamespaceName)
	listOptions := metav1.ListOptions{}
	if p.Labels == nil {
		p.Labels = make(map[string]string)
	}

	p.Labels[PodLabelKey] = statefulSet.Status.CurrentRevision

	labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
	listOptions = metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	}
	podList, err := podClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to get pod list", "err", err.Error())
		return err
	}
	p.output.PodList = podList.Items
	for idx, _ := range p.output.PodList {
		p.output.PodList[idx].ManagedFields = nil
	}
	return nil
}

func (svc *statefulSetService) GetStatefulSetPodList(c context.Context, p GetStatefulSetPodListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}
