package k8s

import (
	"context"
	"errors"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ResponseDTO struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

var clientset *kubernetes.Clientset
var (
	ErrorUnstructuredNil   = errors.New("unstructured CustomResource is nil")
	ErrorFieldEmptyStr     = " field empty"
	ErrorFieldMismatch     = " field mismatched"
	ErrorNamespaceMismatch = "namespace mismatched"
)

const (
	RUNNING = "Running"
	FAILED  = "Failed"
	PENDING = "Pending"
)

func InitClient() {
	// Load kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to load kubeconfig: %v", err))
	}

	// Create Kubernetes client
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("Failed to create Kubernetes client: %v", err))
	}
}

func GetClientset() *kubernetes.Clientset {
	return clientset
}

func ListNamespaces(clientset kubernetes.Interface) ([]string, error) {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list namespaces: %w", err)
	}

	var namespaceNames []string
	for _, ns := range namespaces.Items {
		namespaceNames = append(namespaceNames, ns.Name)
	}
	return namespaceNames, nil
}
