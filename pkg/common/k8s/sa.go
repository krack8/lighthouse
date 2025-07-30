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

type ServiceAccountServiceInterface interface {
	GetServiceAccountList(c context.Context, p GetServiceAccountListInputParams) (interface{}, error)
	GetServiceAccountDetails(c context.Context, p GetServiceAccountDetailsInputParams) (interface{}, error)
	DeployServiceAccount(c context.Context, p DeployServiceAccountInputParams) (interface{}, error)
	DeleteServiceAccount(c context.Context, p DeleteServiceAccountInputParams) (interface{}, error)
}

type serviceAccountService struct{}

var sas serviceAccountService

func ServiceAccountService() *serviceAccountService {
	return &sas
}

const (
	ServiceAccountApiVersion = "v1"
	ServiceAccountKind       = "ServiceAccount"
)

type OutputServiceAccountList struct {
	Result    []corev1.ServiceAccount
	Resource  string
	Remaining int64
	Total     int
}

type GetServiceAccountListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputServiceAccountList
}

func (p *GetServiceAccountListInputParams) Find(c context.Context, serviceAccountClient v1.ServiceAccountInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "service Account")
	filteredServiceAccounts := []corev1.ServiceAccount{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		serviceAccountList, err := serviceAccountClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get service Account list", "err", err.Error())
			return err
		}

		for _, serviceAccount := range serviceAccountList.Items {
			if strings.Contains(serviceAccount.Name, p.Search) {
				filteredServiceAccounts = append(filteredServiceAccounts, serviceAccount)
			}
		}
		length = len(filteredServiceAccounts)
		nextPageToken = serviceAccountList.Continue
		if serviceAccountList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		serviceAccountList, err := serviceAccountClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get service Account list", "err", err.Error())
			return err
		}
		for _, serviceAccount := range serviceAccountList.Items {
			if strings.Contains(serviceAccount.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredServiceAccounts
	p.output.Total = len(filteredServiceAccounts)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetServiceAccountListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching service account list")
	serviceAccountClient := GetKubeClientSet().CoreV1().ServiceAccounts(p.NamespaceName)
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
	var serviceAccountList *corev1.ServiceAccountList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, serviceAccountClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get service Account list", "err", err.Error())
			return err
		}
		return nil
	} else {
		serviceAccountList, err = serviceAccountClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get service Account list", "err", err.Error())
			return err
		}
		remaining := serviceAccountList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: serviceAccountList.Continue}
				res, err := serviceAccountClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get service Account list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = serviceAccountList.Items
		p.output.Total = len(serviceAccountList.Items)
		p.output.Resource = serviceAccountList.Continue
	}
	return nil
}

func (p *GetServiceAccountListInputParams) PostProcess(ctx context.Context) error {
	for i := 0; i < len(p.output.Result); i++ {
		p.output.Result[i].ManagedFields = nil
		p.output.Result[i].TypeMeta.APIVersion = ServiceAccountApiVersion
		p.output.Result[i].TypeMeta.Kind = ServiceAccountKind
	}
	return nil
}

func (svc *serviceAccountService) GetServiceAccountList(c context.Context, p GetServiceAccountListInputParams) (interface{}, error) {
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

type GetServiceAccountDetailsInputParams struct {
	NamespaceName      string
	ServiceAccountName string
	output             corev1.ServiceAccount
}

func (p *GetServiceAccountDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching serviceAccount details of ....", p.NamespaceName)
	serviceAccountsClient := GetKubeClientSet().CoreV1().ServiceAccounts(p.NamespaceName)
	output, err := serviceAccountsClient.Get(context.Background(), p.ServiceAccountName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get serviceAccount ", p.ServiceAccountName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.ManagedFields = nil
	p.output.APIVersion = ServiceAccountApiVersion
	p.output.Kind = ServiceAccountKind
	return nil
}

func (svc *serviceAccountService) GetServiceAccountDetails(c context.Context, p GetServiceAccountDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployServiceAccountInputParams struct {
	ServiceAccount *corev1.ServiceAccount
	output         *corev1.ServiceAccount
}

func (p *DeployServiceAccountInputParams) Process(c context.Context) error {
	ServiceAccountClient := GetKubeClientSet().CoreV1().ServiceAccounts(p.ServiceAccount.Namespace)
	_, err := ServiceAccountClient.Get(context.Background(), p.ServiceAccount.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating serviceAccount in namespace "+p.ServiceAccount.Namespace, "value", p.ServiceAccount.Name)
		p.output, err = ServiceAccountClient.Create(context.Background(), p.ServiceAccount, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create serviceAccount in namespace "+p.ServiceAccount.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("serviceAccount created")
	} else {
		log.Logger.Infow("serviceAccount exist in namespace "+p.ServiceAccount.Namespace, "value", p.ServiceAccount.Name)
		log.Logger.Infow("Updating serviceAccount in namespace "+p.ServiceAccount.Namespace, "value", p.ServiceAccount.Name)
		p.output, err = ServiceAccountClient.Update(context.Background(), p.ServiceAccount, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update serviceAccount ", p.ServiceAccount.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("serviceAccount updated")
	}
	return nil
}

func (svc *serviceAccountService) DeployServiceAccount(c context.Context, p DeployServiceAccountInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteServiceAccountInputParams struct {
	NamespaceName      string
	ServiceAccountName string
}

func (p *DeleteServiceAccountInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting ServiceAccount of ....", p.NamespaceName)
	ServiceAccountClient := GetKubeClientSet().CoreV1().ServiceAccounts(p.NamespaceName)
	_, err := ServiceAccountClient.Get(context.Background(), p.ServiceAccountName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get ServiceAccount ", p.ServiceAccountName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = ServiceAccountClient.Delete(context.Background(), p.ServiceAccountName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete ServiceAccount ", p.ServiceAccountName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *serviceAccountService) DeleteServiceAccount(c context.Context, p DeleteServiceAccountInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
