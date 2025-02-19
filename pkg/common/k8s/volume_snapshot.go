package k8s

import (
	"context"
	cfg "github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/dto"
	"github.com/krack8/lighthouse/pkg/common/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
)

type VolumeSnapshotServiceInterface interface {
	GetVolumeSnapshotList(c context.Context, p GetVolumeSnapshotListInputParams) (interface{}, error)
	GetVolumeSnapshotDetails(c context.Context, p GetVolumeSnapshotDetailsInputParams) (interface{}, error)
	DeployVolumeSnapshot(c context.Context, p DeployVolumeSnapshotInputParams) (interface{}, error)
	DeleteVolumeSnapshot(c context.Context, p DeleteVolumeSnapshotInputParams) (interface{}, error)
}

type volumeSnapshotService struct{}

var vss volumeSnapshotService

func VolumeSnapshotService() *volumeSnapshotService {
	return &vss
}

type GetVolumeSnapshotListInputParams struct {
	NamespaceName string
	Search        string
	Labels        map[string]string
	output        []dto.VolumeSnapshotV1
}

func (p *GetVolumeSnapshotListInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching volume snapshot list")
	var volumeSnapshotList []dto.VolumeSnapshotV1

	listOptions := metav1.ListOptions{}
	if p.Labels != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: p.Labels}
		listOptions = metav1.ListOptions{
			LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		}
	}
	if p.Search != "" {
		listOptions.FieldSelector = fields.OneTermEqualSelector("metadata.name", p.Search).String()
	}
	volumeSnapshotsClient := cfg.GetDynamicClientSet().Resource(cfg.VolumeSnapshotSGVR)
	unstructuredVolumeSnapshotList, err := volumeSnapshotsClient.Namespace(p.NamespaceName).List(context.Background(), listOptions)
	if err != nil {
		log.Logger.Errorw("Failed to get volumeSnapshot list", "err", err.Error())
		return err
	} else {
		if unstructuredVolumeSnapshotList.Items != nil {
			for _, unstructured := range unstructuredVolumeSnapshotList.Items {
				var volumeSnapshot dto.VolumeSnapshotV1
				err := runtime.DefaultUnstructuredConverter.
					FromUnstructured(unstructured.Object, &volumeSnapshot)
				if err != nil {
					log.Logger.Errorw("Get volumeSnapshot list structured conversion", "err", err.Error())
					return err
				}
				volumeSnapshotList = append(volumeSnapshotList, volumeSnapshot)
			}
		}
	}
	p.output = volumeSnapshotList
	return nil
}

func (svc *volumeSnapshotService) GetVolumeSnapshotList(c context.Context, p GetVolumeSnapshotListInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type GetVolumeSnapshotDetailsInputParams struct {
	NamespaceName      string
	VolumeSnapshotName string
	output             dto.VolumeSnapshotV1
}

func (p *GetVolumeSnapshotDetailsInputParams) Process(c context.Context) error {
	log.Logger.Debugw("fetching volumeSnapshot details of ....", p.NamespaceName)
	var volumeSnapshot dto.VolumeSnapshotV1

	volumeSnapshotsClient := cfg.GetDynamicClientSet().Resource(cfg.VolumeSnapshotSGVR)
	unstructuredVolumeSnapshot, err := volumeSnapshotsClient.Namespace(p.NamespaceName).Get(context.Background(), p.VolumeSnapshotName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to get volumeSnapshot ", p.VolumeSnapshotName, "err", err.Error())
		return err
	} else {
		err := runtime.DefaultUnstructuredConverter.
			FromUnstructured(unstructuredVolumeSnapshot.Object, &volumeSnapshot)
		if err != nil {
			log.Logger.Errorw("Get volumeSnapshot details", "err", err.Error())
			return err
		}
	}
	p.output = volumeSnapshot
	return nil
}

func (svc *volumeSnapshotService) GetVolumeSnapshotDetails(c context.Context, p GetVolumeSnapshotDetailsInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   p.output,
	}, nil
}

type DeployVolumeSnapshotInputParams struct {
	VolumeSnapshot *dto.VolumeSnapshotV1
}

func (p *DeployVolumeSnapshotInputParams) Process(c context.Context) error {
	volumeSnapshotsClient := cfg.GetDynamicClientSet().Resource(cfg.VolumeSnapshotSGVR)
	unstructuredVolumeSnapshot := p.VolumeSnapshot.GenerateUnstructured()
	if unstructuredVolumeSnapshot == nil {
		log.Logger.Errorw("unstructured VolumeSnapshot is nil")
		return ErrorUnstructuredNil
	}
	returnVolumeSnapshot, err := volumeSnapshotsClient.Namespace(p.VolumeSnapshot.Namespace).Get(context.Background(), p.VolumeSnapshot.Name, metav1.GetOptions{})
	if err != nil {
		log.Logger.Infow("Creating volumeSnapshot in namespace "+p.VolumeSnapshot.Namespace, "value", p.VolumeSnapshot.Name)
		_, err := volumeSnapshotsClient.Namespace(p.VolumeSnapshot.Namespace).Create(context.Background(), unstructuredVolumeSnapshot, metav1.CreateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to create volumeSnapshot in namespace "+p.VolumeSnapshot.Namespace, "err", err.Error())
			return err
		}
		log.Logger.Infow("volumeSnapshot created")
	} else {
		log.Logger.Infow("volumeSnapshot exist in namespace "+p.VolumeSnapshot.Namespace, "value", p.VolumeSnapshot.Name)
		log.Logger.Infow("Updating volumeSnapshot in namespace "+p.VolumeSnapshot.Namespace, "value", p.VolumeSnapshot.Name)
		unstructuredVolumeSnapshot.SetResourceVersion(returnVolumeSnapshot.GetResourceVersion())
		_, err = volumeSnapshotsClient.Namespace(p.VolumeSnapshot.Namespace).Update(context.Background(), unstructuredVolumeSnapshot, metav1.UpdateOptions{})
		if err != nil {
			log.Logger.Errorw("failed to update volumeSnapshot ", p.VolumeSnapshot.Name, "err", err.Error())
			return err
		}
		log.Logger.Infow("volumeSnapshot updated")
	}
	return nil
}

func (svc *volumeSnapshotService) DeployVolumeSnapshot(c context.Context, p DeployVolumeSnapshotInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}

type DeleteVolumeSnapshotInputParams struct {
	NamespaceName      string
	VolumeSnapshotName string
}

func (p *DeleteVolumeSnapshotInputParams) Process(c context.Context) error {
	log.Logger.Debugw("deleting VolumeSnapshot of ....", p.NamespaceName)
	volumeSnapshotsClient := cfg.GetDynamicClientSet().Resource(cfg.VolumeSnapshotSGVR)
	_, err := volumeSnapshotsClient.Namespace(p.NamespaceName).Get(context.Background(), p.VolumeSnapshotName, metav1.GetOptions{})
	if err != nil {
		log.Logger.Errorw("get VolumeSnapshot ", p.VolumeSnapshotName, "err", err.Error())
		return err
	}
	err = volumeSnapshotsClient.Namespace(p.NamespaceName).Delete(context.Background(), p.VolumeSnapshotName, metav1.DeleteOptions{})
	if err != nil {
		log.Logger.Errorw("Failed to delete VolumeSnapshot ", p.VolumeSnapshotName, "err", err.Error())
		return err
	}
	return nil
}

func (svc *volumeSnapshotService) DeleteVolumeSnapshot(c context.Context, p DeleteVolumeSnapshotInputParams) (interface{}, error) {
	err := p.Process(c)
	if err != nil {
		return nil, err
	}

	return ResponseDTO{
		Status: "success",
		Data:   nil,
	}, nil
}
