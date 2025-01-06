package k8s

import (
	"context"
	"errors"
	"fmt"
	cfg "github.com/krack8/lighthouse/pkg/config"
	"github.com/krack8/lighthouse/pkg/dto"
	"github.com/krack8/lighthouse/pkg/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"strings"
)

type ManifestServiceInterface interface {
	DeployManifest(c context.Context, p DeployManifestInputParams) (interface{}, error)
}

type manifestService struct{}

var mans manifestService

var namespaceErr string = "the server does not allow this method on the requested resource"

func ManifestService() *manifestService {
	return &mans
}

type DeployManifestInputParams struct {
	Manifest *dto.ManifestDto
	Resource string
}

func (p *DeployManifestInputParams) applyManifest(c context.Context, manifestsClient dynamic.NamespaceableResourceInterface, unstructuredManifest *unstructured.Unstructured) error {
	returnManifest, err := manifestsClient.Namespace(p.Manifest.Namespace).Get(context.Background(), p.Manifest.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating manifest", "value", p.Manifest.Name)
		_, err := manifestsClient.Namespace(p.Manifest.Namespace).Create(context.TODO(), unstructuredManifest, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create manifest", "err", err.Error())
			return err
		}
		log.Logger.Infow("manifest created")
	} else {
		log.Logger.Infow("manifest exist", "value", p.Manifest.Name)
		log.Logger.Infow("Updating manifest", "value", p.Manifest.Name)
		unstructuredManifest.SetResourceVersion(returnManifest.GetResourceVersion())
		_, err = manifestsClient.Namespace(p.Manifest.Namespace).Update(context.Background(), unstructuredManifest, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update manifest ", p.Manifest.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("manifest updated")
	}
	return nil
}

func (p *DeployManifestInputParams) Process(c context.Context) error {
	temp := strings.Split(p.Manifest.APIVersion, "/")
	var group string
	var apiVersion string
	if len(temp) > 1 {
		group = temp[0]
		apiVersion = temp[1]
	} else {
		group = ""
		apiVersion = temp[0]
	}
	ManifestSGVR := schema.GroupVersionResource{
		Group:    group,
		Version:  apiVersion,
		Resource: p.Resource,
	}
	manifestsClient := cfg.GetDynamicClientSet().Resource(ManifestSGVR)
	unstructuredManifest := p.Manifest.GenerateUnstructured()
	fmt.Println("##########")
	fmt.Println(unstructuredManifest)
	if unstructuredManifest == nil {
		log.Logger.Errorw("unstructured Manifest is nil")
		return errors.New("unstructured Manifest is nil")
	}
	err := p.applyManifest(c, manifestsClient, unstructuredManifest)
	if err != nil {
		log.Logger.Errorw("failed to create manifest", "err", err.Error())
		if err.Error() == namespaceErr && p.Manifest.Namespace == "" {
			log.Logger.Infow("retrying to apply manifest on default namespace")
			p.Manifest.Namespace = "default"
			err = p.applyManifest(c, manifestsClient, unstructuredManifest)
			if err != nil {
				log.Logger.Errorw("failed to create manifest", "err", err.Error())
				return errors.New(namespaceErr)
			}
		}
		return err
	}
	return nil
}

func (svc *manifestService) DeployManifest(c context.Context, p DeployManifestInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
