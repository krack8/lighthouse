package k8s

import (
	"context"
	"fmt"
	cfg "github.com/krack8/lighthouse/pkg/config"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/types"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	_v1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"strconv"
	"strings"
)

type CrdServiceInterface interface {
	GetCrdList(c context.Context, p GetCrdListInputParams) (interface{}, error)
	GetCrdDetails(c context.Context, p GetCrdDetailsInputParams) (interface{}, error)
	DeployCrd(c context.Context, p DeployCrdInputParams) (interface{}, error)
	DeleteCrd(c context.Context, p DeleteCrdInputParams) (interface{}, error)
}

type crdService struct{}

var cs crdService

func CrdService() *crdService {
	return &cs
}

type OutputCrdList struct {
	CrdForList []types.CrdForList
	Resource   string
	Total      int
}

type GetCrdListInputParams struct {
	Labels   map[string]string
	Continue string
	Search   string
	CrdList  []v1.CustomResourceDefinition
	Limit    string
	output   OutputCrdList
}

func (p *GetCrdListInputParams) Find(c context.Context, crdClient _v1.CustomResourceDefinitionInterface, pageSize int64) (string, []v1.CustomResourceDefinition, error) {
	log.Logger.Debugw("Entering Search mode....", "src", "crd")
	filteredCrds := []v1.CustomResourceDefinition{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		crdList, err := crdClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get crd list", "err", err.Error())
			return "", nil, err
		}

		for _, crd := range crdList.Items {
			if strings.Contains(crd.Name, p.Search) {
				filteredCrds = append(filteredCrds, crd)
			}
		}
		length = len(filteredCrds)
		nextPageToken = crdList.Continue
		if crdList.Continue == "" {
			break
		}
	}
	log.Logger.Info("crds", "count", len(filteredCrds))
	return nextPageToken, filteredCrds, nil
}

func (p *GetCrdListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching crd list...", "debug", "crd list")
	limit := cfg.PageLimit
	if p.Limit != "" {
		limit, _ = strconv.ParseInt(p.Limit, 10, 64)
	}
	crdClient := cfg.GetApiExtensionClientSet().ApiextensionsV1().CustomResourceDefinitions()
	listOptions := metav1.ListOptions{Limit: cfg.PageLimit, Continue: p.Continue}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions = metav1.ListOptions{
			LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		}
	}
	var err error
	var crdList *v1.CustomResourceDefinitionList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		p.output.Resource, p.CrdList, err = p.Find(c, crdClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get crd list", "err", err.Error())
			return err
		}
		count := 0
		all, _ := crdClient.List(context.Background(), metav1.ListOptions{})
		for _, crd := range all.Items {
			if strings.Contains(crd.Name, p.Search) {
				count = count + 1
			}
		}
		p.output.Total = count
		return nil
	} else {
		crdList, err = crdClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get crd list", "err", err.Error())
			return err
		}
		p.CrdList = crdList.Items
		p.output.Resource = crdList.Continue
		listOptions = metav1.ListOptions{}
		if p.Labels != nil {
			labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
			listOptions = metav1.ListOptions{
				LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
			}
		}
		all, _ := crdClient.List(context.Background(), listOptions)
		log.Logger.Infow(fmt.Sprintf("crd total %v", len(all.Items)), "value", "")
		p.output.Total = len(all.Items)
	}
	return nil
}

func (p *GetCrdListInputParams) PostProcess(c context.Context) error {
	var crd types.CrdForList
	var keys map[string]int
	for _, obj := range p.CrdList {
		keys = make(map[string]int)
		crd.CreationTimestamp = obj.CreationTimestamp
		crd.Kind = obj.Spec.Names.Kind
		crd.Name = obj.Name
		crd.Scope = obj.Spec.Scope
		crd.NamePlural = obj.Spec.Names.Plural
		crd.Group = obj.Spec.Group
		for _, ver := range obj.Spec.Versions {
			_, exists := keys[ver.Name]
			if !exists {
				crd.Version = append(crd.Version, ver.Name)
				keys[ver.Name] = 1
			}
		}
		p.output.CrdForList = append(p.output.CrdForList, crd)
		crd.Version = nil
	}
	return nil
}

func (svc *crdService) GetCrdList(c context.Context, p GetCrdListInputParams) (interface{}, error) {
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

type GetCrdDetailsInputParams struct {
	CrdName string
	output  *v1.CustomResourceDefinition
}

func (p *GetCrdDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching crd details of ....", p.CrdName)

	crdClient := cfg.GetApiExtensionClientSet().ApiextensionsV1().CustomResourceDefinitions()
	crd, err := crdClient.Get(context.Background(), p.CrdName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get crd ", p.CrdName, "err", err.Error())
		return err
	}
	crd.ManagedFields = nil
	p.output = crd
	return nil
}

func (svc *crdService) GetCrdDetails(c context.Context, p GetCrdDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployCrdInputParams struct {
	Crd *v1.CustomResourceDefinition
}

func (p *DeployCrdInputParams) Process(c context.Context) error {
	log.Logger.Debugw("creating custom resource definition ....")
	crdClient := cfg.GetApiExtensionClientSet().ApiextensionsV1().CustomResourceDefinitions()

	returnCrd, err := crdClient.Get(context.Background(), p.Crd.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating crd ", "value", p.Crd.Name)
		_, err := crdClient.Create(context.Background(), p.Crd, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create crd in namespace "+p.Crd.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("crd created")
	} else {
		log.Logger.Infow("crd exist ", "value", p.Crd.Name)
		log.Logger.Infow("Updating crd ", "value", p.Crd.Name)
		p.Crd.SetResourceVersion(returnCrd.GetResourceVersion())
		_, err = crdClient.Update(context.Background(), p.Crd, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update crd ", p.Crd.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("crd updated")
	}
	return nil
}

func (svc *crdService) DeployCrd(c context.Context, p DeployCrdInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}

type DeleteCrdInputParams struct {
	CrdName string
}

func (p *DeleteCrdInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting Crd ", "value", p.CrdName)
	crdClient := cfg.GetApiExtensionClientSet().ApiextensionsV1().CustomResourceDefinitions()
	_, err := crdClient.Get(context.Background(), p.CrdName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get Crd ", p.CrdName, "err", err.Error())
		return err
	}
	err = crdClient.Delete(context.Background(), p.CrdName, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to delete Crd ", p.CrdName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *crdService) DeleteCrd(c context.Context, p DeleteCrdInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
