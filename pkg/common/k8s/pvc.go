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

type PvcServiceInterface interface {
	GetPvcList(c context.Context, p GetPvcListInputParams) (interface{}, error)
	GetPvcDetails(c context.Context, p GetPvcDetailsInputParams) (interface{}, error)
	DeployPvc(c context.Context, p DeployPvcInputParams) (interface{}, error)
	DeletePvc(c context.Context, p DeletePvcInputParams) (interface{}, error)
}

type pvcService struct{}

var pvcs pvcService

func PvcService() *pvcService {
	return &pvcs
}

const (
	PVC_API_VERSION = "v1"
	PVC_KIND        = "PersistentVolumeClaim"
)

type OutputPvcList struct {
	Result    []corev1.PersistentVolumeClaim
	Resource  string
	Remaining int64
	Total     int
}

type GetPvcListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputPvcList
}

func (p *GetPvcListInputParams) PostProcess(c context.Context) error {
	for idx, _ := range p.output.Result {
		p.output.Result[idx].ManagedFields = nil
	}
	return nil
}

func (p *GetPvcListInputParams) Find(c context.Context, pvcClient v1.PersistentVolumeClaimInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "pvc")
	filteredPvc := []corev1.PersistentVolumeClaim{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		pvcList, err := pvcClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get pvc list", "err", err.Error())
			return err
		}

		for _, pvc := range pvcList.Items {
			if strings.Contains(pvc.Name, p.Search) {
				filteredPvc = append(filteredPvc, pvc)
			}
		}
		length = len(filteredPvc)
		nextPageToken = pvcList.Continue
		if pvcList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		pvcList, err := pvcClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get pvc list", "err", err.Error())
			return err
		}
		for _, pvc := range pvcList.Items {
			if strings.Contains(pvc.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredPvc
	p.output.Total = len(filteredPvc)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetPvcListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching config map list")
	pvcClient := GetKubeClientSet().CoreV1().PersistentVolumeClaims(p.NamespaceName)
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
	var pvcList *corev1.PersistentVolumeClaimList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, pvcClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get pvc list", "err", err.Error())
			return err
		}
		return nil
	} else {
		pvcList, err = pvcClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get pvc list", "err", err.Error())
			return err
		}
		remaining := pvcList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: pvcList.Continue}
				res, err := pvcClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get pvc list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = pvcList.Items
		p.output.Total = len(pvcList.Items)
		p.output.Resource = pvcList.Continue
	}
	return nil
}

func (pvc *pvcService) GetPvcList(c context.Context, p GetPvcListInputParams) (interface{}, error) {
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

type GetPvcDetailsInputParams struct {
	NamespaceName string
	PvcName       string
	output        corev1.PersistentVolumeClaim
}

func (p *GetPvcDetailsInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *GetPvcDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching pvc details of ....", p.NamespaceName)
	pvcsClient := GetKubeClientSet().CoreV1().PersistentVolumeClaims(p.NamespaceName)
	output, err := pvcsClient.Get(context.Background(), p.PvcName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get pvc ", p.PvcName, "err", err.Error())
		return err
	}
	p.output = *output
	p.output.APIVersion = PVC_API_VERSION
	p.output.Kind = PVC_KIND
	return nil
}

func (pvc *pvcService) GetPvcDetails(c context.Context, p GetPvcDetailsInputParams) (interface{}, error) {
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

type DeployPvcInputParams struct {
	Pvc    *corev1.PersistentVolumeClaim
	output *corev1.PersistentVolumeClaim
}

func (p *DeployPvcInputParams) PostProcess(c context.Context) error {
	p.output.ManagedFields = nil
	return nil
}

func (p *DeployPvcInputParams) Process(c context.Context) error {
	pvcClient := GetKubeClientSet().CoreV1().PersistentVolumeClaims(p.Pvc.Namespace)
	_, err := pvcClient.Get(context.Background(), p.Pvc.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating pvc in namespace "+p.Pvc.Namespace, "value", p.Pvc.Name)
		p.output, err = pvcClient.Create(context.Background(), p.Pvc, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create pvc in namespace "+p.Pvc.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("pvc created")
	} else {
		log.Logger.Infow("Pvc exist in namespace "+p.Pvc.Namespace, "value", p.Pvc.Name)
		log.Logger.Infow("Updating pvc in namespace "+p.Pvc.Namespace, "value", p.Pvc.Name)
		p.output, err = pvcClient.Update(context.Background(), p.Pvc, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update pvc ", p.Pvc.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("pvc updated")
	}
	return nil
}

func (pvc *pvcService) DeployPvc(c context.Context, p DeployPvcInputParams) (interface{}, error) {
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

type DeletePvcInputParams struct {
	NamespaceName string
	PvcName       string
}

func (p *DeletePvcInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting pvc of ....", p.NamespaceName)
	pvcClient := GetKubeClientSet().CoreV1().PersistentVolumeClaims(p.NamespaceName)
	_, err := pvcClient.Get(context.Background(), p.PvcName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get pvc ", p.PvcName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = pvcClient.Delete(context.Background(), p.PvcName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete pvc ", p.PvcName, "err", err.Error())
		return err
	}
	return nil
}

func (pvc *pvcService) DeletePvc(c context.Context, p DeletePvcInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
