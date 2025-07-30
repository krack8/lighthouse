package k8s

import (
	"context"
	"github.com/krack8/lighthouse/pkg/common/log"
	csiv1 "github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumesnapshot/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

type VolumeSnapshotClassServiceInterface interface {
	GetVolumeSnapshotClassList(c context.Context, p GetVolumeSnapshotClassListInputParams) (interface{}, error)
	GetVolumeSnapshotClassDetails(c context.Context, p GetVolumeSnapshotClassDetailsInputParams) (interface{}, error)
}

type volumeSnapshotClassService struct{}

var vscls volumeSnapshotClassService

func VolumeSnapshotClassService() *volumeSnapshotClassService {
	return &vscls
}

type GetVolumeSnapshotClassListInputParams struct {
	Search string
	Labels map[string]string
	output []csiv1.VolumeSnapshotClass
}

func (p *GetVolumeSnapshotClassListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching volume snapshot list")
	volumeSnapshotClassClient := GetSnapshotV1ClientSet().VolumeSnapshotClasses()
	listOptions := metav1.ListOptions{}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions.LabelSelector = labels.Set(labelSelector.MatchLabels).String()
	}
	if p.Search != "" {
		listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
	}
	volumeSnapshotClassList, err := volumeSnapshotClassClient.List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to process get volume snapshot list", "err", err.Error())
		return err
	}
	p.output = volumeSnapshotClassList.Items
	return nil
}

func (svc *volumeSnapshotClassService) GetVolumeSnapshotClassList(c context.Context, p GetVolumeSnapshotClassListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetVolumeSnapshotClassDetailsInputParams struct {
	VolumeSnapshotClassName string
	output                  csiv1.VolumeSnapshotClass
}

func (p *GetVolumeSnapshotClassDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching volumeSnapshotClass details of ....", p.VolumeSnapshotClassName)

	volumeSnapshotClassClient := GetSnapshotV1ClientSet().VolumeSnapshotClasses()
	volumeSnapshotClass, err := volumeSnapshotClassClient.Get(context.Background(), p.VolumeSnapshotClassName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get volumeSnapshotClass ", p.VolumeSnapshotClassName, "err", err.Error())
		return err
	}
	p.output = *volumeSnapshotClass
	return nil
}

func (svc *volumeSnapshotClassService) GetVolumeSnapshotClassDetails(c context.Context, p GetVolumeSnapshotClassDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}
