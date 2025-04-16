package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"strconv"
	"strings"
)

type RoleBindingServiceInterface interface {
	GetRoleBindingList(c context.Context, p GetRoleBindingListInputParams) (interface{}, error)
	GetRoleBindingDetails(c context.Context, p GetRoleBindingDetailsInputParams) (interface{}, error)
	DeployRoleBinding(c context.Context, p DeployRoleBindingInputParams) (interface{}, error)
	DeleteRoleBinding(c context.Context, p DeleteRoleBindingInputParams) (interface{}, error)
}

type roleBindingService struct{}

var rbs roleBindingService

func RoleBindingService() *roleBindingService {
	return &rbs
}

const (
	ROLE_BINDING_API_VERSION = "rbac.authorization.k8s.io/v1"
	ROLE_BINDING_KIND        = "RoleBinding"
)

type OutputRoleBindingList struct {
	Result    []rbacv1.RoleBinding
	Resource  string
	Remaining int64
	Total     int
}

type GetRoleBindingListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputRoleBindingList
}

func (p *GetRoleBindingListInputParams) Find(c context.Context, roleBindingClient v1.RoleBindingInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "role binding")
	filteredRoleBinding := []rbacv1.RoleBinding{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		roleBindingList, err := roleBindingClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get role binding list", "err", err.Error())
			return err
		}

		for _, role := range roleBindingList.Items {
			if strings.Contains(role.Name, p.Search) {
				filteredRoleBinding = append(filteredRoleBinding, role)
			}
		}
		length = len(filteredRoleBinding)
		nextPageToken = roleBindingList.Continue
		if roleBindingList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		roleBindingList, err := roleBindingClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get role binding list", "err", err.Error())
			return err
		}
		for _, roleBinding := range roleBindingList.Items {
			if strings.Contains(roleBinding.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredRoleBinding
	p.output.Total = len(filteredRoleBinding)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetRoleBindingListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching roleBinding list")
	roleBindingClient := GetKubeClientSet().RbacV1().RoleBindings(p.NamespaceName)
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
	var roleBindingList *rbacv1.RoleBindingList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, roleBindingClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get role binding list", "err", err.Error())
			return err
		}
		return nil
	} else {
		roleBindingList, err = roleBindingClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get role binding list", "err", err.Error())
			return err
		}

		roleBindingList, err = roleBindingClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get role binding list", "err", err.Error())
			return err
		}
		remaining := roleBindingList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: roleBindingList.Continue}
				res, err := roleBindingClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get role binding list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = roleBindingList.Items
		p.output.Total = len(roleBindingList.Items)
		p.output.Resource = roleBindingList.Continue
	}
	return nil
}

func (svc *roleBindingService) GetRoleBindingList(c context.Context, p GetRoleBindingListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetRoleBindingDetailsInputParams struct {
	NamespaceName   string
	RoleBindingName string
	output          rbacv1.RoleBinding
}

func (p *GetRoleBindingDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching roleBinding details of ....", p.NamespaceName)
	roleBindingsClient := GetKubeClientSet().RbacV1().RoleBindings(p.NamespaceName)
	output, err := roleBindingsClient.Get(context.Background(), p.RoleBindingName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get roleBinding ", p.RoleBindingName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.APIVersion = ROLE_BINDING_API_VERSION
	p.output.Kind = ROLE_BINDING_KIND
	return nil
}

func (svc *roleBindingService) GetRoleBindingDetails(c context.Context, p GetRoleBindingDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployRoleBindingInputParams struct {
	RoleBinding *rbacv1.RoleBinding
	output      *rbacv1.RoleBinding
}

func (p *DeployRoleBindingInputParams) Process(c context.Context) error {
	RoleBindingClient := GetKubeClientSet().RbacV1().RoleBindings(p.RoleBinding.Namespace)
	_, err := RoleBindingClient.Get(context.Background(), p.RoleBinding.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating roleBinding in namespace "+p.RoleBinding.Namespace, "value", p.RoleBinding.Name)
		p.output, err = RoleBindingClient.Create(context.Background(), p.RoleBinding, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create roleBinding in namespace "+p.RoleBinding.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("roleBinding created")
	} else {
		log.Logger.Infow("roleBinding exist in namespace "+p.RoleBinding.Namespace, "value", p.RoleBinding.Name)
		log.Logger.Infow("Updating roleBinding in namespace "+p.RoleBinding.Namespace, "value", p.RoleBinding.Name)
		p.output, err = RoleBindingClient.Update(context.Background(), p.RoleBinding, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update roleBinding ", p.RoleBinding.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("roleBinding updated")
	}
	return nil
}

func (svc *roleBindingService) DeployRoleBinding(c context.Context, p DeployRoleBindingInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteRoleBindingInputParams struct {
	NamespaceName   string
	RoleBindingName string
}

func (p *DeleteRoleBindingInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting RoleBinding of ....", p.NamespaceName)
	RoleBindingClient := GetKubeClientSet().RbacV1().RoleBindings(p.NamespaceName)
	_, err := RoleBindingClient.Get(context.Background(), p.RoleBindingName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get RoleBinding ", p.RoleBindingName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = RoleBindingClient.Delete(context.Background(), p.RoleBindingName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete RoleBinding ", p.RoleBindingName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *roleBindingService) DeleteRoleBinding(c context.Context, p DeleteRoleBindingInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
