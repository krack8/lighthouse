package k8s_test

import (
	"context"
	v1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestListNamespaces(t *testing.T) {
	// Create a fake Kubernetes client
	clientset := fake.NewSimpleClientset()

	// Add mock namespaces to the fake client
	mockNamespaces := []string{"default", "kube-system", "test-namespace"}
	for _, ns := range mockNamespaces {
		_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{Name: ns},
		}, metav1.CreateOptions{})
		if err != nil {
			t.Fatalf("Failed to create mock namespace: %v", err)
		}
	}

	// List namespaces using the fake client
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Failed to list namespaces: %v", err)
	}

	// Verify that the namespaces match the expected mock namespaces
	var namespaceNames []string
	for _, ns := range namespaces.Items {
		namespaceNames = append(namespaceNames, ns.Name)
	}

	for _, expected := range mockNamespaces {
		found := false
		for _, actual := range namespaceNames {
			if expected == actual {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected namespace %s not found in the list", expected)
		}
	}
}
