package dto

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeletionPolicy string

//const (
//	// volumeSnapshotContentDelete means the snapshot will be deleted from the
//	// underlying storage system on release from its volume snapshot.
//	VolumeSnapshotContentDelete DeletionPolicy = "Delete"
//
//	// volumeSnapshotContentRetain means the snapshot will be left in its current
//	// state on release from its volume snapshot.
//	VolumeSnapshotContentRetain DeletionPolicy = "Retain"
//)

type VolumeSnapshotContent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              VolumeSnapshotContentSpec    `json:"spec"`
	Status            *VolumeSnapshotContentStatus `json:"status"`
}

type VolumeSnapshotContentSpec struct {
	VolumeSnapshotRef       corev1.ObjectReference      `json:"volumeSnapshotRef" protobuf:"bytes,1,opt,name=volumeSnapshotRef"`
	Source                  VolumeSnapshotContentSource `json:"source"`
	DeletionPolicy          DeletionPolicy              `json:"deletionPolicy" protobuf:"bytes,2,opt,name=deletionPolicy"`
	Driver                  string                      `json:"driver"`
	VolumeSnapshotClassName *string                     `json:"volumeSnapshotClassName"`
}

type VolumeSnapshotContentSource struct {
	VolumeHandle   *string `json:"volumeHandle"`
	SnapshotHandle *string `json:"snapshotHandle"`
}

type VolumeSnapshotContentStatus struct {
	CreationTime   *int64  `json:"creationTime"`
	ReadyToUse     *bool   `json:"readyToUse"`
	RestoreSize    *int64  `json:"restoreSize"`
	SnapshotHandle *string `json:"snapshotHandle"`
}
