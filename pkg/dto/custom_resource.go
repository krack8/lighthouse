package dto

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type CustomResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              interface{} `json:"spec"`
	Status            interface{} `json:"status"`
}

func (config *CustomResource) GenerateUnstructured() *unstructured.Unstructured {
	if config == nil {
		return nil
	}
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       config.Kind,
			"apiVersion": config.APIVersion,
			"metadata": map[string]interface{}{
				"name":      config.Name,
				"namespace": config.Namespace,
			},
			"spec": config.Spec,
		},
	}
}
