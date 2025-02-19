package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/log"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	"strconv"
	"strings"
)

type StorageClassServiceInterface interface {
	GetStorageClassList(c context.Context, p GetStorageClassListInputParams) (interface{}, error)
	GetStorageClassDetails(c context.Context, p GetStorageClassDetailsInputParams) (interface{}, error)
	DeployStorageClass(c context.Context, p DeployStorageClassInputParams) (interface{}, error)
	DeleteStorageClass(c context.Context, p DeleteStorageClassInputParams) (interface{}, error)
}

type storageClassService struct{}

var scs storageClassService

func StorageClassService() *storageClassService {
	return &scs
}

type OutputStorageClassList struct {
	Result    []storagev1.StorageClass
	Resource  string
	Remaining int64
	Total     int
}

type GetStorageClassListInputParams struct {
	NamespaceName string
	Search        string
	Continue      string
	Limit         string
	Labels        map[string]string
	output        OutputStorageClassList
}

func (p *GetStorageClassListInputParams) Find(c context.Context, storageClassClient v1.StorageClassInterface, pageSize int64) error {
	log.Logger.Debugw("Entering Search mode....", "src", "storageClass")
	filteredStorageClasses := []storagev1.StorageClass{}
	length := 0
	var nextPageToken string
	nextPageToken = p.Continue
	//limit := int(pageSize)
	for length < int(pageSize) {
		listOptions := metav1.ListOptions{Limit: pageSize, Continue: nextPageToken}
		storageClassList, err := storageClassClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get storageClass list", "err", err.Error())
			return err
		}

		for _, storageClass := range storageClassList.Items {
			if strings.Contains(storageClass.Name, p.Search) {
				filteredStorageClasses = append(filteredStorageClasses, storageClass)
			}
		}
		length = len(filteredStorageClasses)
		nextPageToken = storageClassList.Continue
		if storageClassList.Continue == "" {
			break
		}
	}
	remaining := 0
	if nextPageToken != "" {
		listOptions := metav1.ListOptions{Continue: nextPageToken}
		storageClassList, err := storageClassClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get storageClass list", "err", err.Error())
			return err
		}
		for _, storageClass := range storageClassList.Items {
			if strings.Contains(storageClass.Name, p.Search) {
				remaining = remaining + 1
			}
		}
	}
	p.output.Resource = nextPageToken
	p.output.Result = filteredStorageClasses
	p.output.Total = len(filteredStorageClasses)
	p.output.Remaining = int64(remaining)
	return nil
}

func (p *GetStorageClassListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching config map list")
	storageClassClient := GetKubeClientSet().StorageV1().StorageClasses()
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
	var storageClassList *storagev1.StorageClassList
	if p.Search != "" {
		//listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
		err = p.Find(c, storageClassClient, limit)
		if err != nil {
			log.Logger.Errorw("Failed to get storageClass list", "err", err.Error())
			return err
		}
		return nil
	} else {
		storageClassList, err = storageClassClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get storageClass list", "err", err.Error())
			return err
		}

		storageClassList, err = storageClassClient.List(context.Background(), listOptions)
		if err != nil {
			log.Logger.Errorw("Failed to get storageClass list", "err", err.Error())
			return err
		}
		remaining := storageClassList.RemainingItemCount
		if remaining != nil {
			p.output.Remaining = *remaining
			if p.output.Remaining == 1 {
				listOptions = metav1.ListOptions{Continue: storageClassList.Continue}
				res, err := storageClassClient.List(context.Background(), listOptions)
				p.output.Remaining = int64(len(res.Items))
				if err != nil {
					log.Logger.Errorw("Failed to get storageClass list", "err", err.Error())
					return err
				}
			}
		} else {
			p.output.Remaining = 0
		}
		p.output.Result = storageClassList.Items
		p.output.Total = len(storageClassList.Items)
		p.output.Resource = storageClassList.Continue
	}
	return nil
}

func (svc *storageClassService) GetStorageClassList(c context.Context, p GetStorageClassListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetStorageClassDetailsInputParams struct {
	StorageClassName string
	output           storagev1.StorageClass
}

func (p *GetStorageClassDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching storageClass details of ....", p.StorageClassName)
	storageClasssClient := GetKubeClientSet().StorageV1().StorageClasses()
	output, err := storageClasssClient.Get(context.Background(), p.StorageClassName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get storageClass ", p.StorageClassName, "err", err.Error())
		return err
	}
	p.output = *output
	return nil
}

func (svc *storageClassService) GetStorageClassDetails(c context.Context, p GetStorageClassDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployStorageClassInputParams struct {
	StorageClass *storagev1.StorageClass
	output       *storagev1.StorageClass
}

func (p *DeployStorageClassInputParams) Process(c context.Context) error {
	StorageClassClient := GetKubeClientSet().StorageV1().StorageClasses()
	_, err := StorageClassClient.Get(context.Background(), p.StorageClass.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating storageClass ", "value", p.StorageClass.Name)
		p.output, err = StorageClassClient.Create(context.Background(), p.StorageClass, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create storageClass"+p.StorageClass.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("storageClass created")
	} else {
		log.Logger.Infow("storageClass exists ", "value", p.StorageClass.Name)
		log.Logger.Infow("Updating storageClass ", "value", p.StorageClass.Name)
		p.output, err = StorageClassClient.Update(context.Background(), p.StorageClass, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update storageClass ", p.StorageClass.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("storageClass updated")
	}
	return nil
}

func (svc *storageClassService) DeployStorageClass(c context.Context, p DeployStorageClassInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeleteStorageClassInputParams struct {
	StorageClassName string
}

func (p *DeleteStorageClassInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting StorageClass ", p.StorageClassName)
	StorageClassClient := GetKubeClientSet().StorageV1().StorageClasses()
	_, err := StorageClassClient.Get(context.Background(), p.StorageClassName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get StorageClass ", p.StorageClassName, "err", err.Error())
		return err
	}
	var grace int64 = 1
	err = StorageClassClient.Delete(context.Background(), p.StorageClassName, metav1.DeleteOptions{GracePeriodSeconds: &grace})
	if err != nil {
		log.Logger.Errorw("Failed to delete StorageClass ", p.StorageClassName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *storageClassService) DeleteStorageClass(c context.Context, p DeleteStorageClassInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
