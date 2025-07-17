package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"strconv"
	"strings"
)

type ReplicaSetServiceInterface interface {
	GetReplicaSetList(c context.Context, p GetReplicaSetListInputParams) (interface{}, error)
	GetReplicaSetDetails(c context.Context, p GetReplicaSetDetailsInputParams) (interface{}, error)
	DeployReplicaSet(c context.Context, p DeployReplicaSetInputParams) (interface{}, error)
	DeleteReplicaSet(c context.Context, p DeleteReplicaSetInputParams) (interface{}, error)
}

type replicaSetService struct{}

var rss replicaSetService

func ReplicaSetService() *replicaSetService {
	return &rss
}

const (
	ReplicaSetApiVersion = "apps/v1"
	ReplicaSetKind       = "ReplicaSet"
)

type OutputReplicaSetList struct {
	Result    []appv1.ReplicaSet
	Resource  string
	Remaining int64
	Total     int
}

type GetReplicaSetListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputReplicaSetList
}

func (p *GetReplicaSetListInputParams) Find(c context.Context, replicasetClient v1.ReplicaSetInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "replicaset")
	filteredreplicasets := []appv1.ReplicaSet{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		replicasetList, err := replicasetClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get replicaset list", "err", err.Error())
			return err
		}

		for _, replicaset := range replicasetList.Items {
			if strings.Contains(replicaset.Name, p.Search) {
				filteredreplicasets = append(filteredreplicasets, replicaset)
			}
		}
		length = len(filteredreplicasets)
		nextPageToken = replicasetList.Continue
		if replicasetList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		replicasetList, err := replicasetClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get replicaset list", "err", err.Error())
			return err
		}
		for _, replicaset := range replicasetList.Items {
			if strings.Contains(replicaset.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredreplicasets
	p.output.Total = len(filteredreplicasets)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetReplicaSetListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching replicaset list")
	replicaSetClient := GetKubeClientSet().AppsV1().ReplicaSets(p.NamespaceName)
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
	var replicaSetList *appv1.ReplicaSetList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, replicaSetClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get replicaset list", "err", err.Error())
			return err
		}
		return nil
	} else {
		replicaSetList, err = replicaSetClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get replicaset list", "err", err.Error())
			return err
		}

		replicaSetList, err = replicaSetClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get replicaset list", "err", err.Error())
			return err
		}
		remaining := replicaSetList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: replicaSetList.Continue}
				res, err := replicaSetClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get replicaset list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = replicaSetList.Items
		p.output.Total = len(replicaSetList.Items)
		p.output.Resource = replicaSetList.Continue
	}
	return nil
}

func (p *GetReplicaSetListInputParams) PostProcess(ctx context.Context) error {
	for i := 0; i < len(p.output.Result); i++ {
		p.output.Result[i].ManagedFields = nil
		p.output.Result[i].APIVersion = ReplicaSetApiVersion
		p.output.Result[i].Kind = ReplicaSetKind
	}
	return nil
}

func (svc *replicaSetService) GetReplicaSetList(c context.Context, p GetReplicaSetListInputParams) (interface{}, error) {
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

type GetReplicaSetDetailsInputParams struct {
	NamespaceName  string
	ReplicaSetName string
	output         appv1.ReplicaSet
}

func (p *GetReplicaSetDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching replicaSet details of ....", p.NamespaceName)
	replicaSetsClient := GetKubeClientSet().AppsV1().ReplicaSets(p.NamespaceName)
	output, err := replicaSetsClient.Get(context.Background(), p.ReplicaSetName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get replicaSet ", p.ReplicaSetName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.ManagedFields = nil
	p.output.APIVersion = ReplicaSetApiVersion
	p.output.Kind = ReplicaSetKind
	return nil
}

func (svc *replicaSetService) GetReplicaSetDetails(c context.Context, p GetReplicaSetDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployReplicaSetInputParams struct {
	ReplicaSet *appv1.ReplicaSet
	output     *appv1.ReplicaSet
}

func (p *DeployReplicaSetInputParams) Process(c context.Context) error {
	ReplicaSetClient := GetKubeClientSet().AppsV1().ReplicaSets(p.ReplicaSet.Namespace)
	_, err := ReplicaSetClient.Get(context.Background(), p.ReplicaSet.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating replicaSet in namespace "+p.ReplicaSet.Namespace, "value", p.ReplicaSet.Name)
		p.output, err = ReplicaSetClient.Create(context.Background(), p.ReplicaSet, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create replicaSet in namespace "+p.ReplicaSet.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("replicaSet created")
	} else {
		log.Logger.Infow("replicaSet exist in namespace "+p.ReplicaSet.Namespace, "value", p.ReplicaSet.Name)
		log.Logger.Infow("Updating replicaSet in namespace "+p.ReplicaSet.Namespace, "value", p.ReplicaSet.Name)
		p.output, err = ReplicaSetClient.Update(context.Background(), p.ReplicaSet, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update replicaSet ", p.ReplicaSet.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("replicaSet updated")
	}
	return nil
}

func (svc *replicaSetService) DeployReplicaSet(c context.Context, p DeployReplicaSetInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteReplicaSetInputParams struct {
	NamespaceName  string
	ReplicaSetName string
}

func (p *DeleteReplicaSetInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting ReplicaSet of ....", p.NamespaceName)
	ReplicaSetClient := GetKubeClientSet().AppsV1().ReplicaSets(p.NamespaceName)
	_, err := ReplicaSetClient.Get(context.Background(), p.ReplicaSetName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get ReplicaSet ", p.ReplicaSetName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = ReplicaSetClient.Delete(context.Background(), p.ReplicaSetName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete ReplicaSet ", p.ReplicaSetName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *replicaSetService) DeleteReplicaSet(c context.Context, p DeleteReplicaSetInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
