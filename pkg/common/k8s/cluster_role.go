package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/typed/rbac/v1"
	"strconv"
	"strings"
)

type ClusterRoleServiceInterface interface {
	GetClusterRoleList(c context.Context, p GetClusterRoleListInputParams) (interface{}, error)
	GetClusterRoleDetails(c context.Context, p GetClusterRoleDetailsInputParams) (interface{}, error)
	DeployClusterRole(c context.Context, p DeployClusterRoleInputParams) (interface{}, error)
	DeleteClusterRole(c context.Context, p DeleteClusterRoleInputParams) (interface{}, error)
}

type clusterRoleService struct{}

var croles clusterRoleService

func ClusterRoleService() *clusterRoleService {
	return &croles
}

const (
	ClusterRoleApiVersion = "rbac.authorization.k8s.io/v1"
	ClusterRoleKind       = "ClusterRole"
)

type OutputClusterRoleList struct {
	Result    []rbacv1.ClusterRole
	Resource  string
	Remaining int64
	Total     int
}

type GetClusterRoleListInputParams struct {
	Search   string
	Labels   map[string]string
	Limit    string
	Continue string
	output   OutputClusterRoleList
}

func (p *GetClusterRoleListInputParams) Find(c context.Context, clusterRoleClient v1.ClusterRoleInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "clusterRole")
	filteredClusterRoles := []rbacv1.ClusterRole{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		clusterRoleList, err := clusterRoleClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get clusterRole list", "err", err.Error())
			return err
		}

		for _, clusterRole := range clusterRoleList.Items {
			if strings.Contains(clusterRole.Name, p.Search) {
				filteredClusterRoles = append(filteredClusterRoles, clusterRole)
			}
		}
		length = len(filteredClusterRoles)
		nextPageToken = clusterRoleList.Continue
		if clusterRoleList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		clusterRoleList, err := clusterRoleClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get clusterRole list", "err", err.Error())
			return err
		}
		for _, clusterRole := range clusterRoleList.Items {
			if strings.Contains(clusterRole.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredClusterRoles
	p.output.Total = len(filteredClusterRoles)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetClusterRoleListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching cluster role list")
	clusterRoleClient := GetKubeClientSet().RbacV1().ClusterRoles()
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
	var clusterRoleList *rbacv1.ClusterRoleList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, clusterRoleClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get cluster role list", "err", err.Error())
			return err
		}
		return nil
	} else {
		clusterRoleList, err = clusterRoleClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to cluster role list", "err", err.Error())
			return err
		}

		clusterRoleList, err = clusterRoleClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get cluster role list", "err", err.Error())
			return err
		}
		remaining := clusterRoleList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: clusterRoleList.Continue}
				res, err := clusterRoleClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get cluster role list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = clusterRoleList.Items
		p.output.Total = len(clusterRoleList.Items)
		p.output.Resource = clusterRoleList.Continue
	}
	return nil
}

func (p *GetClusterRoleListInputParams) PostProcess() error {
	for i := 0; i < len(p.output.Result); i++ {
		p.output.Result[i].ManagedFields = nil
		p.output.Result[i].APIVersion = ClusterRoleApiVersion
		p.output.Result[i].Kind = ClusterRoleKind
	}
	return nil
}

func (svc *clusterRoleService) GetClusterRoleList(c context.Context, p GetClusterRoleListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}
	_ = p.PostProcess()
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetClusterRoleDetailsInputParams struct {
	ClusterRoleName string
	output          rbacv1.ClusterRole
}

func (p *GetClusterRoleDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching cluster role details ....")
	clusterRolesClient := GetKubeClientSet().RbacV1().ClusterRoles()
	output, err := clusterRolesClient.Get(context.Background(), p.ClusterRoleName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get cluster role ", p.ClusterRoleName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.ManagedFields = nil
	p.output.APIVersion = ClusterRoleApiVersion
	p.output.Kind = ClusterRoleKind
	return nil
}

func (svc *clusterRoleService) GetClusterRoleDetails(c context.Context, p GetClusterRoleDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployClusterRoleInputParams struct {
	ClusterRole *rbacv1.ClusterRole
	output      *rbacv1.ClusterRole
}

func (p *DeployClusterRoleInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *DeployClusterRoleInputParams) Process(c context.Context) error {
	clusterRoleClient := GetKubeClientSet().RbacV1().ClusterRoles()
	returnedClusterRole, err := clusterRoleClient.Get(context.Background(), p.ClusterRole.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating clusterRole ", "value", p.ClusterRole.Name)
		p.output, err = clusterRoleClient.Create(context.Background(), p.ClusterRole, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create clusterRole ", "err", err.Error())
			return err
		}
		log.Logger.Infow("clusterRole created")
	} else {
		log.Logger.Infow("ClusterRole exist ", "value", p.ClusterRole.Name)
		log.Logger.Infow("Updating clusterRole ", "value", p.ClusterRole.Name)
		p.ClusterRole.SetResourceVersion(returnedClusterRole.ResourceVersion)
		p.output, err = clusterRoleClient.Update(context.Background(), p.ClusterRole, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update clusterRole ", p.ClusterRole.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("clusterRole updated")
	}
	return nil
}

func (svc *clusterRoleService) DeployClusterRole(c context.Context, p DeployClusterRoleInputParams) (interface{}, error) {
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

type DeleteClusterRoleInputParams struct {
	ClusterRoleName string
}

func (p *DeleteClusterRoleInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting clusterRole ....", p.ClusterRoleName)
	clusterRoleClient := GetKubeClientSet().RbacV1().ClusterRoles()
	_, err := clusterRoleClient.Get(context.Background(), p.ClusterRoleName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get clusterRole ", p.ClusterRoleName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = clusterRoleClient.Delete(context.Background(), p.ClusterRoleName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete clusterRole ", p.ClusterRoleName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *clusterRoleService) DeleteClusterRole(c context.Context, p DeleteClusterRoleInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
