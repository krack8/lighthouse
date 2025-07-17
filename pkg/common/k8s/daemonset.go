package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"strconv"
	"strings"
)

type DaemonSetServiceInterface interface {
	GetDaemonSetList(c context.Context, p GetDaemonSetListInputParams) (interface{}, error)
	GetDaemonSetDetails(c context.Context, p GetDaemonSetDetailsInputParams) (interface{}, error)
	DeployDaemonSet(c context.Context, p DeployDaemonSetInputParams) (interface{}, error)
	DeleteDaemonSet(c context.Context, p DeleteDaemonSetInputParams) (interface{}, error)
	GetDaemonSetStats(c context.Context, p GetDaemonSetStatsInputParams) (interface{}, error)
}

type daemonSetService struct{}

var dss daemonSetService

func DaemonSetService() *daemonSetService {
	return &dss
}

const (
	DaemonSetApiVersion = "apps/v1"
	DaemonSetKind       = "DaemonSet"
)

type OutputDaemonSetList struct {
	Result    []appv1.DaemonSet
	Resource  string
	Remaining int64
	Total     int
}

type GetDaemonSetListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputDaemonSetList
}

func (p *GetDaemonSetListInputParams) Find(c context.Context, daemonsetClient v1.DaemonSetInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "daemonset")
	filteredDaemonsets := []appv1.DaemonSet{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		daemonsetList, err := daemonsetClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get daemonset list", "err", err.Error())
			return err
		}

		for _, daemonset := range daemonsetList.Items {
			if strings.Contains(daemonset.Name, p.Search) {
				filteredDaemonsets = append(filteredDaemonsets, daemonset)
			}
		}
		length = len(filteredDaemonsets)
		nextPageToken = daemonsetList.Continue
		if daemonsetList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		daemonsetList, err := daemonsetClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get daemonset list", "err", err.Error())
			return err
		}
		for _, daemonset := range daemonsetList.Items {
			if strings.Contains(daemonset.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredDaemonsets
	p.output.Total = len(filteredDaemonsets)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetDaemonSetListInputParams) PostProcess(c context.Context) error {
	for idx, _ := range p.output.Result {
		p.output.Result[idx].ManagedFields = nil
		p.output.Result[idx].APIVersion = DaemonSetApiVersion
		p.output.Result[idx].Kind = DaemonSetKind
	}
	return nil
}

func (p *GetDaemonSetListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching config map list")
	daemonSetClient := GetKubeClientSet().AppsV1().DaemonSets(p.NamespaceName)
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
	var daemonsetList *appv1.DaemonSetList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, daemonSetClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get statefulset list", "err", err.Error())
			return err
		}
		return nil
	} else {
		daemonsetList, err = daemonSetClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get statefulset list", "err", err.Error())
			return err
		}

		daemonsetList, err = daemonSetClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get statefulset list", "err", err.Error())
			return err
		}
		remaining := daemonsetList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: daemonsetList.Continue}
				res, err := daemonSetClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get statefulset list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = daemonsetList.Items
		p.output.Total = len(daemonsetList.Items)
		p.output.Resource = daemonsetList.Continue
	}
	return nil
}

func (svc *daemonSetService) GetDaemonSetList(c context.Context, p GetDaemonSetListInputParams) (interface{}, error) {
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

type GetDaemonSetDetailsInputParams struct {
	NamespaceName string
	DaemonSetName string
	output        appv1.DaemonSet
}

func (p *GetDaemonSetDetailsInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *GetDaemonSetDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching daemonSet details of ....", p.NamespaceName)
	daemonSetsClient := GetKubeClientSet().AppsV1().DaemonSets(p.NamespaceName)
	output, err := daemonSetsClient.Get(context.Background(), p.DaemonSetName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get daemonSet ", p.DaemonSetName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.ManagedFields = nil
	p.output.APIVersion = DaemonSetApiVersion
	p.output.Kind = DaemonSetKind
	return nil
}

func (svc *daemonSetService) GetDaemonSetDetails(c context.Context, p GetDaemonSetDetailsInputParams) (interface{}, error) {
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

type DeployDaemonSetInputParams struct {
	DaemonSet *appv1.DaemonSet
	output    *appv1.DaemonSet
}

func (p *DeployDaemonSetInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *DeployDaemonSetInputParams) Process(c context.Context) error {
	daemonSetClient := GetKubeClientSet().AppsV1().DaemonSets(p.DaemonSet.Namespace)
	_, err := daemonSetClient.Get(context.Background(), p.DaemonSet.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating daemonSet in namespace "+p.DaemonSet.Namespace, "value", p.DaemonSet.Name)
		p.output, err = daemonSetClient.Create(context.Background(), p.DaemonSet, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create daemonSet in namespace "+p.DaemonSet.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("daemonSet created")
	} else {
		log.Logger.Infow("DaemonSet exist in namespace "+p.DaemonSet.Namespace, "value", p.DaemonSet.Name)
		log.Logger.Infow("Updating daemonSet in namespace "+p.DaemonSet.Namespace, "value", p.DaemonSet.Name)
		p.output, err = daemonSetClient.Update(context.Background(), p.DaemonSet, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update daemonSet ", p.DaemonSet.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("daemonSet updated")
	}
	return nil
}

func (svc *daemonSetService) DeployDaemonSet(c context.Context, p DeployDaemonSetInputParams) (interface{}, error) {
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

type DeleteDaemonSetInputParams struct {
	NamespaceName string
	DaemonSetName string
}

func (p *DeleteDaemonSetInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting daemonSet of ....", p.NamespaceName)
	daemonSetClient := GetKubeClientSet().AppsV1().DaemonSets(p.NamespaceName)
	_, err := daemonSetClient.Get(context.Background(), p.DaemonSetName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get daemonSet ", p.DaemonSetName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = daemonSetClient.Delete(context.Background(), p.DaemonSetName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete daemonSet ", p.DaemonSetName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *daemonSetService) DeleteDaemonSet(c context.Context, p DeleteDaemonSetInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}

type StatsDaemonSet struct {
	Total    int
	Ready    int
	NotReady int
}

func (s *StatsDaemonSet) New() *StatsDaemonSet {
	return &StatsDaemonSet{Total: 0, Ready: 0, NotReady: 0}
}

type GetDaemonSetStatsInputParams struct {
	NamespaceName string
	Labels        map[string]string
	Search        string
	output        *StatsDaemonSet
}

func (p *GetDaemonSetStatsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching daemonSet list stats")
	daemonSetClient := GetKubeClientSet().AppsV1().DaemonSets(p.NamespaceName)

	listOptions := metav1.ListOptions{}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions.LabelSelector = labels.Set(labelSelector.MatchLabels).String()
	}

	if p.Search != "" {
		listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
	}

	daemonSetList, err := daemonSetClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to get daemonSet list stats", "err", err.Error())
		return err
	}

	p.output = p.output.New()

	for _, obj := range daemonSetList.Items {
		p.output.Total += int(obj.Status.DesiredNumberScheduled)
		p.output.Ready += int(obj.Status.NumberAvailable)
	}

	p.output.NotReady = p.output.Total - p.output.Ready
	return nil
}

func (svc *daemonSetService) GetDaemonSetStats(c context.Context, p GetDaemonSetStatsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}
