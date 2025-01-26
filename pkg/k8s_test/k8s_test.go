package k8s_test

import (
	"context"
	"github.com/krack8/lighthouse/pkg/k8s"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestListNamespaces(t *testing.T) {
	// Mock the Kubernetes client using a fake client
	clientset := fake.NewClientset()

	// Create mock namespaces
	namespaceNames := []string{"default", "kube-system", "test-namespace"}
	for _, ns := range namespaceNames {
		_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{Name: ns},
		}, metav1.CreateOptions{})
		if err != nil {
			t.Fatalf("Failed to create namespace: %v", err)
		}
	}

	// Fetch namespaces using FetchNamespaces function
	fetchedNamespaces, err := k8s.ListNamespaces(clientset)
	if err != nil {
		t.Fatalf("FetchNamespaces failed: %v", err)
	}

	// Validate the fetched namespaces
	for _, expected := range namespaceNames {
		found := false
		for _, actual := range fetchedNamespaces {
			if expected == actual {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected namespace %q not found", expected)
		}
	}
}
