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

type NamespaceServiceInterface interface {
	GetNamespaceList(c context.Context, p GetNamespaceListInputParams) (interface{}, error)
	GetNamespaceNameList(c context.Context, p GetNamespaceNamesInputParams) (interface{}, error)
	GetNamespaceDetails(c context.Context, p GetNamespaceInputParams) (interface{}, error)
	DeployNamespace(c context.Context, p DeployNamespaceInputParams) (interface{}, error)
	DeleteNamespace(c context.Context, p DeleteNamespaceInputParams) (interface{}, error)
}

type namespaceService struct {
}

var nss namespaceService

func NamespaceService() *namespaceService {
	return &nss
}

func (p *GetNamespaceInputParams) setClient() {
	p.Client = config.GetKubeClientSet().CoreV1().Namespaces()
}

func getNamespaceClient() v1.NamespaceInterface {
	return config.GetKubeClientSet().CoreV1().Namespaces()
}

func (p *GetNamespaceListInputParams) removeNamespaceListFields() interface{} {
	namespaceList := p.output.Result.([]corev1.Namespace)
	for idx, _ := range namespaceList {
		namespaceList[idx].ManagedFields = nil
	}
	return namespaceList
}

func removeNamespaceFields(namespace corev1.Namespace) corev1.Namespace {
	namespace.ManagedFields = nil
	return namespace
}

type OutputNamespaceList struct {
	Result    interface{}
	Resource  string
	Remaining int64
	Total     int
}

type GetNamespaceListInputParams struct {
	Search   string
	Labels   map[string]string
	Continue string
	Limit    string
	output   OutputNamespaceList
}

func (p *GetNamespaceListInputParams) Find(c context.Context, namespaceClient v1.NamespaceInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "namespace")
	filteredNamespaces := []corev1.Namespace{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		namespaceList, err := namespaceClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get namespace list", "err", err.Error())
			return err
		}

		for _, namespace := range namespaceList.Items {
			if strings.Contains(namespace.Name, p.Search) {
				filteredNamespaces = append(filteredNamespaces, namespace)
			}
		}
		length = len(filteredNamespaces)
		nextPageToken = namespaceList.Continue
		if namespaceList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		namespaceList, err := namespaceClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get namespace list", "err", err.Error())
			return err
		}
		for _, namespace := range namespaceList.Items {
			if strings.Contains(namespace.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredNamespaces
	p.output.Total = len(filteredNamespaces)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetNamespaceListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching namespace list")
	namespacesClient := config.GetKubeClientSet().CoreV1().Namespaces()
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
	namespaceList, err := namespacesClient.List(context.Background(), listOptions)
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, namespacesClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get namespace list", "err", err.Error())
			return err
		}
		return nil
	} else {
		namespaceList, err = namespacesClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get namespace list", "err", err.Error())
			return err
		}
		remaining := namespaceList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: namespaceList.Continue}
				res, err := namespacesClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get namespace list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = namespaceList.Items
		p.output.Total = len(namespaceList.Items)
		p.output.Resource = namespaceList.Continue
	}
	return nil
}

func (namespace *namespaceService) GetNamespaceList(c context.Context, p GetNamespaceListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}
	p.output.Result = p.removeNamespaceListFields()
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetNamespaceNamesInputParams struct {
	output []string
}

func (p *GetNamespaceNamesInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching namespace details of ....", "namespace name list")
	namespacesClient := config.GetKubeClientSet().CoreV1().Namespaces()
	listOptions := metav1.ListOptions{}
	namespaceList, err := namespacesClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to get namespace list", "err", err.Error())
		return err
	}
	for _, namespace := range namespaceList.Items {
		p.output = append(p.output, namespace.Name)
	}
	return nil
}

func (namespace *namespaceService) GetNamespaceNameList(c context.Context, p GetNamespaceNamesInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetNamespaceInputParams struct {
	Client        v1.NamespaceInterface
	NamespaceName string
	output        corev1.Namespace
}

func (p *GetNamespaceInputParams) GetClient() v1.NamespaceInterface {
	if p.Client != nil {
		return p.Client
	}
	return getNamespaceClient()
}

func (p *GetNamespaceInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching namespace details of ....", p.NamespaceName)
	output, err := p.GetClient().Get(context.Background(), p.NamespaceName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get namespace "+p.NamespaceName, "err", err.Error())
		return err
	}
	output.ManagedFields = nil
	p.output = *output
	return nil
}

func (namespace *namespaceService) GetNamespaceDetails(c context.Context, p GetNamespaceInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}
	p.output = removeNamespaceFields(p.output)
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployNamespaceInputParams struct {
	Namespace *corev1.Namespace
	output    *corev1.Namespace
	Client    v1.NamespaceInterface
}

func (p *DeployNamespaceInputParams) GetClient() v1.NamespaceInterface {
	if p.Client != nil {
		return p.Client
	}
	return getNamespaceClient()
}

func (p *DeployNamespaceInputParams) Process(c context.Context) error {
	_, err := p.GetClient().Get(context.Background(), p.Namespace.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating namespace "+p.Namespace.Name, "value", p.Namespace.Name)
		p.output, err = p.GetClient().Create(context.Background(), p.Namespace, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create namespace "+p.Namespace.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("namespace created")
	} else {
		log.Logger.Infow("Namespace exists "+p.Namespace.Name, "value", p.Namespace.Name)
		log.Logger.Infow("Updating namespace "+p.Namespace.Name, "value", p.Namespace.Name)
		p.output, err = p.GetClient().Update(context.Background(), p.Namespace, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update namespace ", p.Namespace.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("namespace updated")
	}
	return nil
}

func (namespace *namespaceService) DeployNamespace(c context.Context, p DeployNamespaceInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}
	p.output.ManagedFields = nil
	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteNamespaceInputParams struct {
	NamespaceName string
	Client        v1.NamespaceInterface
}

func (p *DeleteNamespaceInputParams) GetClient() v1.NamespaceInterface {
	if p.Client != nil {
		return p.Client
	}
	return getNamespaceClient()
}

func (p *DeleteNamespaceInputParams) Process(c context.Context) error {
	log.Logger.Debugw("Deleting Namespace ....", "value", p.NamespaceName)
	err := p.GetClient().Delete(context.Background(), p.NamespaceName, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to delete namespace "+p.NamespaceName, "err", err.Error())
		return err
	}
	return nil
}

func (namespace *namespaceService) DeleteNamespace(c context.Context, p DeleteNamespaceInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Msg:    "deleted namespace " + p.NamespaceName,
	}, nil
}
