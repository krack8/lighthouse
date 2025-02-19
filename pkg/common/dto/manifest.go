package dto

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type ManifestDto struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              interface{} `json:"spec"`
	Status            interface{} `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
	Rules             interface{} `json:"rules,omitempty"`
	RoleRef           interface{} `json:"roleRef,omitempty"`
	Subjects          interface{} `json:"subjects,omitempty"`
	Provisioner       interface{} `json:"provisioner,omitempty"`
}

func (manifest *ManifestDto) GenerateUnstructured() *unstructured.Unstructured {
	if manifest == nil {
		return nil
	}
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       manifest.Kind,
			"apiVersion": manifest.APIVersion,
			"metadata": map[string]interface{}{
				"name": manifest.Name,
			},
			"spec":        manifest.Spec,
			"rules":       manifest.Rules,
			"roleRef":     manifest.RoleRef,
			"subjects":    manifest.Subjects,
			"provisioner": manifest.Provisioner,
		},
	}
}
