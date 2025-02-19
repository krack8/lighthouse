package k8s

import (
	"context"
	"errors"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/dto"
	"github.com/krack8/lighthouse/pkg/common/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"strconv"
	"strings"
)

type CustomResourceServiceInterface interface {
	GetCustomResourceList(c context.Context, p GetCustomResourceListInputParams) (interface{}, error)
	GetCustomResourceDetails(c context.Context, p GetCustomResourceDetailsInputParams) (interface{}, error)
	DeployCustomResource(c context.Context, p DeployCustomResourceInputParams) (interface{}, error)
	DeleteCustomResource(c context.Context, p DeleteCustomResourceInputParams) (interface{}, error)
}

type customResourceService struct{}

var cres customResourceService

func CustomResourceService() *customResourceService {
	return &cres
}

type OutputCustomResourceList struct {
	Result    []*dto.CustomResource
	Resource  string
	Remaining int64
	Total     int
}

type GetCustomResourceListInputParams struct {
	NamespaceName      string
	Search             string
	CustomResourceSGVR schema.GroupVersionResource
	Labels             map[string]string
	Continue           string
	Limit              string
	output             OutputCustomResourceList
}

func (p *GetCustomResourceListInputParams) PreProcess(c context.Context) error {
	if p.CustomResourceSGVR.Resource == "" || p.CustomResourceSGVR.Group == "" {
		if p.CustomResourceSGVR.Resource == "" {
			ErrorFieldEmptyStr = ",resource" + ErrorFieldEmptyStr
		}
		if p.CustomResourceSGVR.Version == "" {
			ErrorFieldEmptyStr = ",version" + ErrorFieldEmptyStr
		}
		if p.CustomResourceSGVR.Group == "" {
			ErrorFieldEmptyStr = ",group" + ErrorFieldEmptyStr
		}
		return errors.New(ErrorFieldEmptyStr[1:])
	}
	return nil
}

func (p *GetCustomResourceListInputParams) Find(c context.Context, customResourceClient dynamic.NamespaceableResourceInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "custom resource")
	filteredCustomResource := []*dto.CustomResource{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		unstructuredCustomResource, err := customResourceClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get custom resource list", "err", err.Error())
			return err
		} else {
			log.Logger.Info("get custom resource list")
			if unstructuredCustomResource.Items != nil {
				for _, unstructured := range unstructuredCustomResource.Items {
					var cr dto.CustomResource
					err := runtime.DefaultUnstructuredConverter.
						FromUnstructured(unstructured.Object, &cr)
					if err != nil {
						log.Logger.Errorw("Get custom resource list structured conversion", "err", err.Error())
						return err
					}
					if strings.Contains(cr.Name, p.Search) {
						filteredCustomResource = append(filteredCustomResource, &cr)
					}
				}
			}
		}
		length = len(filteredCustomResource)
		nextPageToken = unstructuredCustomResource.GetContinue()
		if nextPageToken == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		unstructuredCustomResource, err := customResourceClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get custom resource list", "err", err.Error())
			return err
		}
		if unstructuredCustomResource.Items != nil {
			for _, unstructured := range unstructuredCustomResource.Items {
				var cr dto.CustomResource
				err := runtime.DefaultUnstructuredConverter.
					FromUnstructured(unstructured.Object, &cr)
				if err != nil {
					log.Logger.Errorw("Get custom resource list structured conversion", "err", err.Error())
					return err
				}
				if strings.Contains(cr.Name, p.Search) {
					remaining = remaining + 1
				}
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredCustomResource
	p.output.Total = len(filteredCustomResource)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetCustomResourceListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching customResource list")
	var customResourceList []*dto.CustomResource
	var err error
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

	customResourcesClient := config.GetDynamicClientSet().Resource(p.CustomResourceSGVR)
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, customResourcesClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get custom resource list", "err", err.Error())
			return err
		}
		return nil
	} else {
		unstructuredCustomResourceList, err := customResourcesClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get custom resource list", "err", err.Error())
			return err
		}
		if unstructuredCustomResourceList.Items != nil {
			for _, unstructured := range unstructuredCustomResourceList.Items {
				var cr dto.CustomResource
				err := runtime.DefaultUnstructuredConverter.
					FromUnstructured(unstructured.Object, &cr)
				if err != nil {
					log.Logger.Errorw("Get custom resource list structured conversion", "err", err.Error())
					return err
				}
				customResourceList = append(customResourceList, &cr)
			}
		}
		remaining := unstructuredCustomResourceList.GetRemainingItemCount()
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: unstructuredCustomResourceList.GetContinue()}
				res, err := customResourcesClient.List(context.Background(), listOptions)
				if err != nil {
					log.Logger.Errorw("failed to get custom resource remaining count", "err", err.Error())
					return err
				}
				p.output.Remaining = int64(len(res.Items))
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Resource = unstructuredCustomResourceList.GetContinue()
		p.output.Result = customResourceList
		p.output.Total = len(customResourceList)
	}
	return nil
}

func (svc *customResourceService) GetCustomResourceList(c context.Context, p GetCustomResourceListInputParams) (interface{}, error) {
	err := p.PreProcess(c)
	if err != nil {
		return nil, err
	}

	err = p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetCustomResourceDetailsInputParams struct {
	Name               string
	NamespaceName      string
	CustomResourceSGVR schema.GroupVersionResource
	output             dto.CustomResource
}

func (p *GetCustomResourceDetailsInputParams) PreProcess(c context.Context) error {
	if p.CustomResourceSGVR.Resource == "" || p.CustomResourceSGVR.Group == "" || p.CustomResourceSGVR.Version == "" {
		if p.CustomResourceSGVR.Resource == "" {
			ErrorFieldEmptyStr = ",resource" + ErrorFieldEmptyStr
		}
		if p.CustomResourceSGVR.Version == "" {
			ErrorFieldEmptyStr = ",version" + ErrorFieldEmptyStr
		}
		if p.CustomResourceSGVR.Group == "" {
			ErrorFieldEmptyStr = ",group" + ErrorFieldEmptyStr
		}
		return errors.New(ErrorFieldEmptyStr[1:])
	}
	return nil
}

func (p *GetCustomResourceDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching customResource ", p.Name)
	var customResource dto.CustomResource
	var unstructuredCustomResource *unstructured.Unstructured
	var err error

	customResourcesClient := config.GetDynamicClientSet().Resource(p.CustomResourceSGVR)
	if p.NamespaceName != "" {
		unstructuredCustomResource, err = customResourcesClient.Namespace(p.NamespaceName).Get(context.Background(), p.Name, metav1.GetOptions{})
	} else {
		unstructuredCustomResource, err = customResourcesClient.Get(context.Background(), p.Name, metav1.GetOptions{})
	}
	if err != nil {
		log.Logger.Errorw("Failed to get customResource "+p.Name, "err", err.Error())
		return err
	} else {
		log.Logger.Info("get customResource ", p.Name)
		err := runtime.DefaultUnstructuredConverter.
			FromUnstructured(unstructuredCustomResource.Object, &customResource)
		if err != nil {
			log.Logger.Error("Get customResource structured conversion", "err", err.Error())
			log.Logger.Errorw("Get customResource structured conversion", "err", err.Error())
			return err
		}
	}
	p.output = customResource
	return nil
}

func (svc *customResourceService) GetCustomResourceDetails(c context.Context, p GetCustomResourceDetailsInputParams) (interface{}, error) {
	err := p.PreProcess(c)
	if err != nil {
		return nil, err
	}

	err = p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployCustomResourceInputParams struct {
	NamespaceName      string
	Kind               string
	CustomResourceSGVR schema.GroupVersionResource
	CustomResource     *dto.CustomResource
}

func (p *DeployCustomResourceInputParams) PreProcess(c context.Context) error {
	if p.CustomResourceSGVR.Resource == "" || p.CustomResourceSGVR.Group == "" || p.CustomResourceSGVR.Version == "" {
		if p.CustomResourceSGVR.Resource == "" {
			ErrorFieldEmptyStr = ",resource" + ErrorFieldEmptyStr
		}
		if p.CustomResourceSGVR.Version == "" {
			ErrorFieldEmptyStr = ",version" + ErrorFieldEmptyStr
		}
		if p.CustomResourceSGVR.Group == "" {
			ErrorFieldEmptyStr = ",group" + ErrorFieldEmptyStr
		}
		return errors.New(ErrorFieldEmptyStr[1:])
	}
	//if p.NamespaceName != "" && p.NamespaceName != p.CustomResource.Namespace {
	//	return errors.New(ErrorNamespaceMismatch)
	//}
	//if p.NamespaceName == "" && p.CustomResource.Namespace != "" {
	//	return errors.New(ErrorNamespaceMismatch)
	//}
	if p.Kind != p.CustomResource.Kind || p.CustomResourceSGVR.Group+"/"+p.CustomResourceSGVR.Version != p.CustomResource.APIVersion {
		if p.Kind != p.CustomResource.Kind {
			ErrorFieldMismatch = ",kind"
		}
		if p.CustomResourceSGVR.Group+"/"+p.CustomResourceSGVR.Version != p.CustomResource.APIVersion {
			ErrorFieldMismatch = ",version"
		}
		return errors.New(ErrorFieldMismatch[1:])
	}
	return nil
}

func (p *DeployCustomResourceInputParams) Process(c context.Context) error {
	customResourcesClient := config.GetDynamicClientSet().Resource(p.CustomResourceSGVR)
	unstructuredCustomResource := p.CustomResource.GenerateUnstructured()
	if unstructuredCustomResource == nil {
		log.Logger.Errorw("unstructured CustomResource is nil")
		return ErrorUnstructuredNil
	}
	var returnCustomResource *unstructured.Unstructured
	var err error
	if p.CustomResource.Namespace != "" {
		returnCustomResource, err = customResourcesClient.Namespace(p.CustomResource.Namespace).Get(context.Background(), p.CustomResource.Name, metav1.GetOptions{})
		if err != nil {
			log.Logger.Infow("Creating customResource in namespace "+p.CustomResource.Namespace, "value", p.CustomResource.Name)
			_, err := customResourcesClient.Namespace(p.CustomResource.Namespace).Create(context.Background(), unstructuredCustomResource, metav1.CreateOptions{})
			if err != nil {
				log.Logger.Errorw("failed to create customResource in namespace "+p.CustomResource.Namespace, "err", err.Error())
				return err
			}
			log.Logger.Infow("customResource created")
		} else {
			log.Logger.Infow("customResource exist in namespace "+p.CustomResource.Namespace, "value", p.CustomResource.Name)
			log.Logger.Infow("Updating customResource in namespace "+p.CustomResource.Namespace, "value", p.CustomResource.Name)
			unstructuredCustomResource.SetResourceVersion(returnCustomResource.GetResourceVersion())
			_, err = customResourcesClient.Namespace(p.CustomResource.Namespace).Update(context.Background(), unstructuredCustomResource, metav1.UpdateOptions{})
			if err != nil {
				log.Logger.Errorw("failed to update customResource ", p.CustomResource.Name, "err", err.Error())
				return err
			}
			log.Logger.Infow("customResource updated")
		}
	} else {
		returnCustomResource, err = customResourcesClient.Get(context.Background(), p.CustomResource.Name, metav1.GetOptions{})
		if err != nil {
			log.Logger.Infow("Creating customResource ", "value", p.CustomResource.Name)
			_, err := customResourcesClient.Create(context.Background(), unstructuredCustomResource, metav1.CreateOptions{})
			if err != nil {
				log.Logger.Errorw("failed to create customResource ", "err", err.Error())
				return err
			}
			log.Logger.Infow("customResource created")
		} else {
			log.Logger.Infow("customResource exist ", "value", p.CustomResource.Name)
			log.Logger.Infow("Updating customResource ", "value", p.CustomResource.Name)
			unstructuredCustomResource.SetResourceVersion(returnCustomResource.GetResourceVersion())
			_, err = customResourcesClient.Update(context.Background(), unstructuredCustomResource, metav1.UpdateOptions{})
			if err != nil {
				log.Logger.Errorw("failed to update customResource ", p.CustomResource.Name, "err", err.Error())
				return err
			}
			log.Logger.Infow("customResource updated")
		}
	}
	return nil
}

func (svc *customResourceService) DeployCustomResource(c context.Context, p DeployCustomResourceInputParams) (interface{}, error) {
	err := p.PreProcess(c)
	if err != nil {
		return nil, err
	}
	err = p.Process(c)
	if err != nil {
		return nil, err
	}
	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}

type DeleteCustomResourceInputParams struct {
	NamespaceName      string
	CustomResourceName string
	CustomResourceSGVR schema.GroupVersionResource
}

func (p *DeleteCustomResourceInputParams) PreProcess(c context.Context) error {
	if p.CustomResourceSGVR.Resource == "" || p.CustomResourceSGVR.Group == "" || p.CustomResourceSGVR.Version == "" {
		if p.CustomResourceSGVR.Resource == "" {
			ErrorFieldEmptyStr = ",resource" + ErrorFieldEmptyStr
		}
		if p.CustomResourceSGVR.Version == "" {
			ErrorFieldEmptyStr = ",version" + ErrorFieldEmptyStr
		}
		if p.CustomResourceSGVR.Group == "" {
			ErrorFieldEmptyStr = ",group" + ErrorFieldEmptyStr
		}
		return errors.New(ErrorFieldEmptyStr[1:])
	}
	return nil
}

func (p *DeleteCustomResourceInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting CustomResource ", p.CustomResourceName)
	customResourcesClient := config.GetDynamicClientSet().Resource(p.CustomResourceSGVR)
	if p.NamespaceName != "" {
		_, err := customResourcesClient.Namespace(p.NamespaceName).Get(context.Background(), p.CustomResourceName, metav1.GetOptions{})
		if err != nil {
			log.Logger.Errorw("get CustomResource ", p.CustomResourceName, "err", err.Error())
			return err
		}
		err = customResourcesClient.Namespace(p.NamespaceName).Delete(context.Background(), p.CustomResourceName, metav1.DeleteOptions{})
		if err != nil {
			log.Logger.Errorw("Failed to delete CustomResource ", p.CustomResourceName, "err", err.Error())
			return err
		}
	} else {
		_, err := customResourcesClient.Get(context.Background(), p.CustomResourceName, metav1.GetOptions{})
		if err != nil {
			log.Logger.Errorw("get CustomResource ", p.CustomResourceName, "err", err.Error())
			return err
		}
		err = customResourcesClient.Delete(context.Background(), p.CustomResourceName, metav1.DeleteOptions{})
		if err != nil {
			log.Logger.Errorw("Failed to delete CustomResource ", p.CustomResourceName, "err", err.Error())
			return err
		}
	}
	return nil
}

func (svc *customResourceService) DeleteCustomResource(c context.Context, p DeleteCustomResourceInputParams) (interface{}, error) {
	err := p.PreProcess(c)
	if err != nil {
		return nil, err
	}
	err = p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
