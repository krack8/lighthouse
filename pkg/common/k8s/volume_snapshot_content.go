package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/dto"
	"github.com/krack8/lighthouse/pkg/common/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
)

type VolumeSnapshotContentServiceInterface interface {
	GetVolumeSnapshotContentList(c context.Context, p GetVolumeSnapshotContentListInputParams) (interface{}, error)
	GetVolumeSnapshotContentDetails(c context.Context, p GetVolumeSnapshotContentDetailsInputParams) (interface{}, error)
}

type volumeSnapshotContentService struct{}

var vscs volumeSnapshotContentService

func VolumeSnapshotContentService() *volumeSnapshotContentService {
	return &vscs
}

type GetVolumeSnapshotContentListInputParams struct {
	Search string
	Labels map[string]string
	output []dto.VolumeSnapshotContent
}

func (p *GetVolumeSnapshotContentListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching volume snapshot content list")
	var volumeSnapshotContentList []dto.VolumeSnapshotContent

	listOptions := metav1.ListOptions{}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions.LabelSelector = labels.Set(labelSelector.MatchLabels).String()
	}
	if p.Search != "" {
		listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
	}
	volumeSnapshotContentsClient := GetDynamicClientSet().Resource(VolumeSnapshotContentSGVR)
	unstructuredVolumeSnapshotContentList, err := volumeSnapshotContentsClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to get volumeSnapshotContent list", "err", err.Error())
		return err
	} else {
		if unstructuredVolumeSnapshotContentList.Items != nil {
			for _, unstructured := range unstructuredVolumeSnapshotContentList.Items {
				var volumeSnapshotContent dto.VolumeSnapshotContent
				err := runtime.DefaultUnstructuredConverter.
					FromUnstructured(unstructured.Object, &volumeSnapshotContent)
				if err != nil {
					log.Logger.Errorw("Get volumeSnapshotContent list structured conversion", "err", err.Error())
					return err
				}
				volumeSnapshotContentList = append(volumeSnapshotContentList, volumeSnapshotContent)
			}
		}
	}
	p.output = volumeSnapshotContentList
	return nil
}

func (svc *volumeSnapshotContentService) GetVolumeSnapshotContentList(c context.Context, p GetVolumeSnapshotContentListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetVolumeSnapshotContentDetailsInputParams struct {
	VolumeSnapshotContentName string
	output                    dto.VolumeSnapshotContent
}

func (p *GetVolumeSnapshotContentDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching volumeSnapshotContent details of ....", p.VolumeSnapshotContentName)
	var volumeSnapshotContent dto.VolumeSnapshotContent

	volumeSnapshotContentsClient := GetDynamicClientSet().Resource(VolumeSnapshotContentSGVR)
	unstructuredVolumeSnapshotContent, err := volumeSnapshotContentsClient.Get(context.Background(), p.VolumeSnapshotContentName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get volumeSnapshotContent ", p.VolumeSnapshotContentName, "err", err.Error())
		return err
	} else {
		err := runtime.DefaultUnstructuredConverter.
			FromUnstructured(unstructuredVolumeSnapshotContent.Object, &volumeSnapshotContent)
		if err != nil {
			log.Logger.Errorw("Get volumeSnapshotContent details", "err", err.Error())
			return err
		}
	}
	p.output = volumeSnapshotContent
	return nil
}

func (svc *volumeSnapshotContentService) GetVolumeSnapshotContentDetails(c context.Context, p GetVolumeSnapshotContentDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}
