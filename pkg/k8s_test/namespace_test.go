package k8s_test

import (
	"context"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

//type NamespaceTestSuite struct {
//	suite.Suite
//	clientSet *fake.Clientset
//}
//
//func TestNamespace(t *testing.T) {
//	suite.Run(t, new(NamespaceTestSuite))
//}
//
//func (s *NamespaceTestSuite) SetupNamespaceSuite() {
//}
//
//func (s *NamespaceTestSuite) SetupTest() {
//	s.clientSet = fake.NewClientset()
//	// Set up before *each* test runs (e.g., reset mocks, clear databases)
//}
//
//func (s *NamespaceTestSuite) TearDownTest() {
//	// Clean up after *each* test runs
//}

func TestGetNamespaceDetails(t *testing.T) {
	log.InitializeLogger()
	// Mock the Kubernetes client
	clientset := fake.NewClientset(&corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-namespace",
		},
	})

	cfg.KubeClientSet = clientset // Assign the fake clientset
	// Test case 1: Successful retrieval
	t.Run("Success", func(t *testing.T) {
		p := k8s.GetNamespaceInputParams{
			NamespaceName: "test-namespace",
			Client:        clientset.CoreV1().Namespaces(),
		}
		expectedNamespace := corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-namespace",
			},
		}
		expectedNamespace = removeNamespaceFields(expectedNamespace) // Apply the function

		result, err := k8s.NamespaceService().GetNamespaceDetails(context.Background(), p)
		assert := assert.New(t)
		assert.NoError(err)
		assert.NotNil(result)

		response, ok := result.(k8s.ResponseDTO)
		assert.True(ok)
		assert.Equal("success", response.Status)
		assert.Equal(expectedNamespace, response.Data)

	})

	// Test case 2: Namespace not found
	t.Run("Error", func(t *testing.T) {
		p := k8s.GetNamespaceInputParams{
			NamespaceName: "non-existent-namespace",
			Client:        clientset.CoreV1().Namespaces(),
		}

		result, err := k8s.NamespaceService().GetNamespaceDetails(context.Background(), p)
		assert := assert.New(t)
		assert.Error(err)
		assert.Nil(result)

	})
}

// A dummy removeNamespaceFields function for testing purposes
func removeNamespaceFields(ns corev1.Namespace) corev1.Namespace {
	return ns
}

var cfg config

type config struct {
	KubeClientSet *fake.Clientset
}

func (c *config) GetKubeClientSet() *fake.Clientset {
	return c.KubeClientSet
}
