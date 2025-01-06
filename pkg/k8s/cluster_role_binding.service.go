package k8s

import (
	"context"
	cfg "github.com/krack8/lighthouse/pkg/config"
	"github.com/krack8/lighthouse/pkg/log"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"strconv"
	"strings"
)

type ClusterRoleBindingServiceInterface interface {
	GetClusterRoleBindingList(c context.Context, p GetClusterRoleBindingListInputParams) (interface{}, error)
	GetClusterRoleBindingDetails(c context.Context, p GetClusterRoleBindingDetailsInputParams) (interface{}, error)
	DeployClusterRoleBinding(c context.Context, p DeployClusterRoleBindingInputParams) (interface{}, error)
	DeleteClusterRoleBinding(c context.Context, p DeleteClusterRoleBindingInputParams) (interface{}, error)
}

type clusterRoleBindingService struct{}

var crbs clusterRoleBindingService

func ClusterRoleBindingService() *clusterRoleBindingService {
	return &crbs
}

type OutputClusterRoleBindingList struct {
	Result    []rbacv1.ClusterRoleBinding
	Resource  string
	Remaining int64
	Total     int
}

type GetClusterRoleBindingListInputParams struct {
	Search   string
	Labels   map[string]string
	Continue string
	Limit    string
	output   OutputClusterRoleBindingList
}

func (p *GetClusterRoleBindingListInputParams) Find(c context.Context, clusterRoleBindingClient v1.ClusterRoleBindingInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "clusterRoleBinding")
	filteredClusterRoleBindings := []rbacv1.ClusterRoleBinding{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		clusterRoleBindingList, err := clusterRoleBindingClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get clusterRoleBinding list", "err", err.Error())
			return err
		}

		for _, clusterRoleBinding := range clusterRoleBindingList.Items {
			if strings.Contains(clusterRoleBinding.Name, p.Search) {
				filteredClusterRoleBindings = append(filteredClusterRoleBindings, clusterRoleBinding)
			}
		}
		length = len(filteredClusterRoleBindings)
		nextPageToken = clusterRoleBindingList.Continue
		if clusterRoleBindingList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		clusterRoleBindingList, err := clusterRoleBindingClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get clusterRoleBinding list", "err", err.Error())
			return err
		}
		for _, clusterRoleBinding := range clusterRoleBindingList.Items {
			if strings.Contains(clusterRoleBinding.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredClusterRoleBindings
	p.output.Total = len(filteredClusterRoleBindings)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetClusterRoleBindingListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching cluster role binding list")
	clusterRoleBindingClient := cfg.GetKubeClientSet().RbacV1().ClusterRoleBindings()
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
	var clusterRoleBindingList *rbacv1.ClusterRoleBindingList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, clusterRoleBindingClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get cluster role list", "err", err.Error())
			return err
		}
		return nil
	} else {
		clusterRoleBindingList, err = clusterRoleBindingClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to cluster role list", "err", err.Error())
			return err
		}

		clusterRoleBindingList, err = clusterRoleBindingClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get cluster role list", "err", err.Error())
			return err
		}
		remaining := clusterRoleBindingList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: clusterRoleBindingList.Continue}
				res, err := clusterRoleBindingClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get cluster role list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = clusterRoleBindingList.Items
		p.output.Total = len(clusterRoleBindingList.Items)
		p.output.Resource = clusterRoleBindingList.Continue
	}
	return nil
}

func (svc *clusterRoleBindingService) GetClusterRoleBindingList(c context.Context, p GetClusterRoleBindingListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetClusterRoleBindingDetailsInputParams struct {
	ClusterRoleBindingName string
	output                 rbacv1.ClusterRoleBinding
}

func (p *GetClusterRoleBindingDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching cluster role details ....")
	clusterRoleBindingsClient := cfg.GetKubeClientSet().RbacV1().ClusterRoleBindings()
	output, err := clusterRoleBindingsClient.Get(context.Background(), p.ClusterRoleBindingName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get cluster role binding", p.ClusterRoleBindingName, "err", err.Error())
		return err
	}
	p.output = *output
	return nil
}

func (svc *clusterRoleBindingService) GetClusterRoleBindingDetails(c context.Context, p GetClusterRoleBindingDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployClusterRoleBindingInputParams struct {
	ClusterRoleBinding *rbacv1.ClusterRoleBinding
	output             *rbacv1.ClusterRoleBinding
}

func (p *DeployClusterRoleBindingInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *DeployClusterRoleBindingInputParams) Process(c context.Context) error {
	clusterRoleBindingClient := cfg.GetKubeClientSet().RbacV1().ClusterRoleBindings()
	_, err := clusterRoleBindingClient.Get(context.Background(), p.ClusterRoleBinding.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating clusterRoleBinding ", "value", p.ClusterRoleBinding.Name)
		p.output, err = clusterRoleBindingClient.Create(context.Background(), p.ClusterRoleBinding, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create clusterRoleBinding ", "err", err.Error())
			return err
		}
		log.Logger.Infow("clusterRoleBinding created")
	} else {
		log.Logger.Infow("ClusterRoleBinding exist ", "value", p.ClusterRoleBinding.Name)
		log.Logger.Infow("Updating clusterRoleBinding ", "value", p.ClusterRoleBinding.Name)
		p.output, err = clusterRoleBindingClient.Update(context.Background(), p.ClusterRoleBinding, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update clusterRoleBinding ", p.ClusterRoleBinding.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("clusterRoleBinding updated")
	}
	return nil
}

func (svc *clusterRoleBindingService) DeployClusterRoleBinding(c context.Context, p DeployClusterRoleBindingInputParams) (interface{}, error) {
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

type DeleteClusterRoleBindingInputParams struct {
	ClusterRoleBindingName string
}

func (p *DeleteClusterRoleBindingInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting clusterRoleBinding ....", p.ClusterRoleBindingName)
	clusterRoleBindingClient := cfg.GetKubeClientSet().RbacV1().ClusterRoleBindings()
	_, err := clusterRoleBindingClient.Get(context.Background(), p.ClusterRoleBindingName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get clusterRoleBinding ", p.ClusterRoleBindingName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = clusterRoleBindingClient.Delete(context.Background(), p.ClusterRoleBindingName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete clusterRoleBinding ", p.ClusterRoleBindingName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *clusterRoleBindingService) DeleteClusterRoleBinding(c context.Context, p DeleteClusterRoleBindingInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
