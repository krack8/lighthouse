package dto

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type VolumeSnapshotV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              VolumeSnapshotSpec   `json:"spec"`
	Status            VolumeSnapshotStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type VolumeSnapshotSpec struct {
	VolumeSnapshotClassName string               `json:"volumeSnapshotClassName"`
	Source                  VolumeSnapshotSource `json:"source"`
}

type VolumeSnapshotSource struct {
	PersistentVolumeClaimName *string `json:"persistentVolumeClaimName"`
	VolumeSnapshotContentName *string `json:"volumeSnapshotContentName"`
}

type VolumeSnapshotStatus struct {
	ReadyToUse                     *bool              `json:"readyToUse,omitempty" protobuf:"varint,4,opt,name=readyToUse"`
	BoundVolumeSnapshotContentName *string            `json:"boundVolumeSnapshotContentName,omitempty"`
	RestoreSize                    *resource.Quantity `json:"restoreSize,omitempty" protobuf:"bytes,3,opt,name=restoreSize"` // integer will return in byte
	CreationTime                   *metav1.Time       `json:"creationTime,omitempty" protobuf:"varint,2,opt,name=creationTime"`
}

func (volumeSnapshot *VolumeSnapshotV1) GenerateUnstructured() *unstructured.Unstructured {
	if volumeSnapshot == nil {
		return nil
	}
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       volumeSnapshot.Kind,
			"apiVersion": volumeSnapshot.APIVersion,
			"metadata": map[string]interface{}{
				"name":      volumeSnapshot.Name,
				"namespace": volumeSnapshot.Namespace,
			},
			"spec": map[string]interface{}{
				"volumeSnapshotClassName": volumeSnapshot.Spec.VolumeSnapshotClassName,
				"source": map[string]interface{}{
					"persistentVolumeClaimName": volumeSnapshot.Spec.Source.PersistentVolumeClaimName,
				},
			},
		},
	}
}
