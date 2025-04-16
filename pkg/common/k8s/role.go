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

type RoleServiceInterface interface {
	GetRoleList(c context.Context, p GetRoleListInputParams) (interface{}, error)
	GetRoleDetails(c context.Context, p GetRoleDetailsInputParams) (interface{}, error)
	DeployRole(c context.Context, p DeployRoleInputParams) (interface{}, error)
	DeleteRole(c context.Context, p DeleteRoleInputParams) (interface{}, error)
}

type roleService struct{}

var rs roleService

func RoleService() *roleService {
	return &rs
}

const (
	ROLE_API_VERSION = "rbac.authorization.k8s.io/v1"
	ROLE_KIND        = "Role"
)

type OutputRoleList struct {
	Result    []rbacv1.Role
	Resource  string
	Remaining int64
	Total     int
}

type GetRoleListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputRoleList
}

func (p *GetRoleListInputParams) Find(c context.Context, roleClient v1.RoleInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "role")
	filteredRole := []rbacv1.Role{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		roleList, err := roleClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get role list", "err", err.Error())
			return err
		}

		for _, role := range roleList.Items {
			if strings.Contains(role.Name, p.Search) {
				filteredRole = append(filteredRole, role)
			}
		}
		length = len(filteredRole)
		nextPageToken = roleList.Continue
		if roleList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		roleList, err := roleClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get role list", "err", err.Error())
			return err
		}
		for _, role := range roleList.Items {
			if strings.Contains(role.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredRole
	p.output.Total = len(filteredRole)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetRoleListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching role list")
	roleClient := GetKubeClientSet().RbacV1().Roles(p.NamespaceName)
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
	var roleList *rbacv1.RoleList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, roleClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get role list", "err", err.Error())
			return err
		}
		return nil
	} else {
		roleList, err = roleClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get role list", "err", err.Error())
			return err
		}

		roleList, err = roleClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get role list", "err", err.Error())
			return err
		}
		remaining := roleList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: roleList.Continue}
				res, err := roleClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get role list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = roleList.Items
		p.output.Total = len(roleList.Items)
		p.output.Resource = roleList.Continue
	}
	return nil
}

func (svc *roleService) GetRoleList(c context.Context, p GetRoleListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetRoleDetailsInputParams struct {
	NamespaceName string
	RoleName      string
	output        rbacv1.Role
}

func (p *GetRoleDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching role details of ....", p.NamespaceName)
	rolesClient := GetKubeClientSet().RbacV1().Roles(p.NamespaceName)
	output, err := rolesClient.Get(context.Background(), p.RoleName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get role ", p.RoleName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.APIVersion = ROLE_API_VERSION
	p.output.Kind = ROLE_KIND
	return nil
}

func (svc *roleService) GetRoleDetails(c context.Context, p GetRoleDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployRoleInputParams struct {
	Role   *rbacv1.Role
	output *rbacv1.Role
}

func (p *DeployRoleInputParams) Process(c context.Context) error {
	RoleClient := GetKubeClientSet().RbacV1().Roles(p.Role.Namespace)
	_, err := RoleClient.Get(context.Background(), p.Role.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating role in namespace "+p.Role.Namespace, "value", p.Role.Name)
		p.output, err = RoleClient.Create(context.Background(), p.Role, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create role in namespace "+p.Role.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("role created")
	} else {
		log.Logger.Infow("role exist in namespace "+p.Role.Namespace, "value", p.Role.Name)
		log.Logger.Infow("Updating role in namespace "+p.Role.Namespace, "value", p.Role.Name)
		p.output, err = RoleClient.Update(context.Background(), p.Role, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update role ", p.Role.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("role updated")
	}
	return nil
}

func (svc *roleService) DeployRole(c context.Context, p DeployRoleInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteRoleInputParams struct {
	NamespaceName string
	RoleName      string
}

func (p *DeleteRoleInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting Role of ....", p.NamespaceName)
	RoleClient := GetKubeClientSet().RbacV1().Roles(p.NamespaceName)
	_, err := RoleClient.Get(context.Background(), p.RoleName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get Role ", p.RoleName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = RoleClient.Delete(context.Background(), p.RoleName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete Role ", p.RoleName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *roleService) DeleteRole(c context.Context, p DeleteRoleInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
