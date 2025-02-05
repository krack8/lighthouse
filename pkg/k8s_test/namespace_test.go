package k8s_test

//
//import (
//	"context"
//	"github.com/krack8/lighthouse/pkg/controller/api"
//	"github.com/krack8/lighthouse/pkg/k8s"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/suite"
//	v1 "k8s.io/api/core/v1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/client-go/kubernetes/fake"
//	"testing"
//)
//
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
//
//func (s *NamespaceTestSuite) TestGetNamespaceDetails(t *testing.T) {
//	// Mock the Kubernetes client
//	clientset := fake.NewSimpleClientset(&corev1.Namespace{
//		ObjectMeta: metav1.ObjectMeta{
//			Name: "test-namespace",
//		},
//	})
//	cfg.kubeClientSet = clientset // Assign the fake clientset
//
//	// Test case 1: Successful retrieval
//	t.Run("Success", func(t *testing.T) {
//		p := GetNamespaceInputParams{
//			NamespaceName: "test-namespace",
//		}
//		expectedNamespace := corev1.Namespace{
//			ObjectMeta: metav1.ObjectMeta{
//				Name: "test-namespace",
//			},
//		}
//		expectedNamespace = removeNamespaceFields(expectedNamespace) // Apply the function
//
//		result, err := namespaceService{}.GetNamespaceDetails(context.Background(), p)
//		assert.NoError(t, err)
//		assert.NotNil(t, result)
//
//		response, ok := result.(ResponseDTO)
//		assert.True(t, ok)
//		assert.Equal(t, "success", response.Status)
//		assert.Equal(t, expectedNamespace, response.Data)
//
//	})
//
//	// Test case 2: Namespace not found
//	t.Run("NotFound", func(t *testing.T) {
//		p := GetNamespaceInputParams{
//			NamespaceName: "non-existent-namespace",
//		}
//
//		result, err := namespaceService{}.GetNamespaceDetails(context.Background(), p)
//
//		assert.Error(t, err)
//		assert.Nil(t, result)
//
//	})
//}
//
//
//// A dummy removeNamespaceFields function for testing purposes
//func removeNamespaceFields(ns corev1.Namespace) corev1.Namespace{
//	return ns
//}
//
//type ResponseDTO struct {
//	Status string      `json:"status"`
//	Data   interface{} `json:"data"`
//}
//var cfg config
//
//type config struct {
//	KubeClientSet *fake.Clientset
//}
//
//func (c *config) GetKubeClientSet() *fake.Clientset {
//	return c.KubeClientSet
//}
