package k8s

import (
	"context"
	k8s2 "github.com/krack8/lighthouse/pkg/common/k8s"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

type ConfigMapTestSuite struct {
	suite.Suite
	clientSet *fake.Clientset
	configMap *corev1.ConfigMap
}

func TestConfigMap(t *testing.T) {
	suite.Run(t, new(ConfigMapTestSuite))
}

func (s *ConfigMapTestSuite) SetupSuite() { // Setup for the entire suite
	log.InitializeTestLogger()
	// Initialize Test logger once for the suite
}

func (s *ConfigMapTestSuite) TearDownSuite() { // Teardown for the entire suite
	// Close resources, etc. if needed.
}

func (s *ConfigMapTestSuite) SetupTest() {
	s.configMap = &corev1.ConfigMap{}
	s.configMap.Name = "test-configmap"
	s.configMap.Namespace = "test-namespace"
	s.clientSet = fake.NewClientset(s.configMap)
	// Setup before each test
	// Reset mocks, clear databases, etc. if needed for *each* test
}

func (s *ConfigMapTestSuite) TearDownTest() { // Teardown after each test
	// Clean up after *each* test runs
}

func (s *ConfigMapTestSuite) TestGetConfigMapDetails() {
	s.T().Run("Success GetConfigMapDetails", func(t *testing.T) {
		t.Log("Test CASE: Get existing ConfigMap with existing namespace")
		p := k8s2.GetConfigMapDetailsInputParams{
			ConfigMapName: s.configMap.Name,
			NamespaceName: s.configMap.Namespace,
			Client:        s.clientSet.CoreV1().ConfigMaps(s.configMap.Namespace), // Use the clientSet from SetupSuite
		}
		expectedConfigMap := corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      s.configMap.Name,
				Namespace: s.configMap.Namespace,
			},
		}
		expectedConfigMap = removeConfigMapFields(expectedConfigMap)

		result, err := k8s2.ConfigMapService().GetConfigMapDetails(context.Background(), p)
		assert := assert.New(s.T())
		assert.NoError(err)
		assert.NotNil(result)

		response, ok := result.(k8s2.ResponseDTO)
		assert.True(ok)
		assert.Equal("success", response.Status)
		assert.Equal(expectedConfigMap, response.Data)
	})
	s.T().Run("Error GetConfigMapDetails", func(t *testing.T) {
		t.Log("Test CASE: Get ConfigMap with non-existent namespace")
		p := k8s2.GetConfigMapDetailsInputParams{
			ConfigMapName: "non-existent-namespace",
			Client:        s.clientSet.CoreV1().ConfigMaps(s.configMap.Namespace), // Use the clientSet from SetupSuite
		}

		result, err := k8s2.ConfigMapService().GetConfigMapDetails(context.Background(), p)
		assert := assert.New(s.T())
		assert.Error(err)
		assert.Nil(result)
	})
}

// A dummy removeConfigMapFields function for testing purposes
func removeConfigMapFields(ns corev1.ConfigMap) corev1.ConfigMap {
	return ns
}

//
//func (s *ConfigMapTestSuite) TestDeleteConfigMapSuccess() {
//	s.T().Run("Success DeleteConfigMap", func(t *testing.T) {
//		t.Log("Test CASE: Delete ConfigMap with existing namespace")
//		p := k8s.DeleteConfigMapInputParams{
//			ConfigMapName: "test-namespace",
//			Client:        s.clientSet.CoreV1().ConfigMaps(),
//		}
//
//		result, err := k8s.ConfigMapService().DeleteConfigMap(context.Background(), p)
//		assert := assert.New(s.T())
//		assert.NoError(err)
//		assert.NotNil(result)
//
//		response, ok := result.(k8s.ResponseDTO)
//		assert.True(ok)
//		assert.Equal("success", response.Status)
//		assert.Equal("deleted namespace test-namespace", response.Msg)
//
//		_, err = s.clientSet.CoreV1().ConfigMaps().Get(context.Background(), "test-namespace", metav1.GetOptions{})
//		assert.Error(err)
//	})
//
//	s.T().Run("Error DeleteConfigMap", func(t *testing.T) {
//		t.Log("Test CASE: Delete ConfigMap with non-existent namespace")
//		p := k8s.DeleteConfigMapInputParams{
//			ConfigMapName: "non-existent-namespace",
//			Client:        s.clientSet.CoreV1().ConfigMaps(),
//		}
//
//		result, err := k8s.ConfigMapService().DeleteConfigMap(context.Background(), p)
//		assert := assert.New(s.T())
//		assert.Error(err)
//		assert.Nil(result)
//	})
//}
//
//func (s *ConfigMapTestSuite) TestDeployConfigMap() {
//	s.T().Run("Success CreateConfigMap", func(t *testing.T) {
//		t.Log("Test CASE: Create ConfigMap with new namespace")
//		namespace := &corev1.ConfigMap{
//			ObjectMeta: metav1.ObjectMeta{
//				Name: "new-namespace",
//			},
//			TypeMeta: metav1.TypeMeta{
//				Kind:       "ConfigMap",
//				APIVersion: "v1",
//			},
//		}
//
//		p := k8s.DeployConfigMapInputParams{
//			ConfigMap: namespace,
//			Client:    s.clientSet.CoreV1().ConfigMaps(),
//		}
//
//		result, err := k8s.ConfigMapService().DeployConfigMap(context.Background(), p)
//		assert := assert.New(t)
//		assert.NoError(err)
//		assert.NotNil(result)
//
//		response, ok := result.(k8s.ResponseDTO)
//		assert.True(ok)
//		assert.Equal("success", response.Status)
//		assert.Equal(namespace, response.Data)
//		// Verify namespace exists
//		fetchedConfigMap, err := s.clientSet.CoreV1().ConfigMaps().Get(context.Background(), namespace.Name, metav1.GetOptions{})
//		assert.NoError(err)
//		assert.Equal(namespace.Name, fetchedConfigMap.Name)
//	})
//
//	s.T().Run("Success UpdateConfigMap", func(t *testing.T) {
//		t.Log("Test CASE: Update ConfigMap with existing namespace")
//		updatedLabels := map[string]string{"updated": "true"}
//		namespace := &corev1.ConfigMap{
//			ObjectMeta: metav1.ObjectMeta{
//				Name:   "test-namespace",
//				Labels: updatedLabels,
//			},
//		}
//
//		p := k8s.DeployConfigMapInputParams{
//			ConfigMap: namespace,
//			Client:    s.clientSet.CoreV1().ConfigMaps(),
//		}
//		result, err := k8s.ConfigMapService().DeployConfigMap(context.Background(), p)
//		assert := assert.New(t)
//		assert.NoError(err)
//		assert.NotNil(result)
//
//		response, ok := result.(k8s.ResponseDTO)
//		assert.True(ok)
//		assert.Equal("success", response.Status)
//		assert.Equal(namespace, response.Data)
//
//		// Verify namespace was updated
//		fetchedConfigMap, err := s.clientSet.CoreV1().ConfigMaps().Get(context.Background(), "test-namespace", metav1.GetOptions{})
//		assert.NoError(err)
//		assert.Equal(updatedLabels, fetchedConfigMap.Labels)
//	})
//	s.T().Run("Error CreateConfigMap - Create Fails", func(t *testing.T) {
//		t.Log("Test CASE: Create ConfigMap with new namespace fail")
//		namespace := &corev1.ConfigMap{
//			ObjectMeta: metav1.ObjectMeta{
//				Name: "new-namespace",
//			},
//		}
//		clientSet := fake.NewClientset()
//		// Make the fake client return an error on create
//		clientSet.PrependReactor("create", "namespaces", func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
//			return true, nil, assert.AnError // Return an error
//		})
//
//		p := k8s.DeployConfigMapInputParams{
//			ConfigMap: namespace,
//			Client:    clientSet.CoreV1().ConfigMaps(),
//		}
//
//		result, err := k8s.ConfigMapService().DeployConfigMap(context.Background(), p)
//		assert := assert.New(t)
//		assert.Error(err)
//		assert.Nil(result)
//	})
//	s.T().Run("Error UpdateConfigMap - Update Fails", func(t *testing.T) {
//		t.Log("Test CASE: Update ConfigMap with existing namespace fail")
//		namespace := &corev1.ConfigMap{
//			ObjectMeta: metav1.ObjectMeta{
//				Name: "test-namespace",
//			},
//		}
//
//		clientSet := fake.NewClientset(&corev1.ConfigMap{
//			ObjectMeta: metav1.ObjectMeta{
//				Name: "test-namespace",
//			},
//		})
//		// Make the fake client return an error on update
//		clientSet.PrependReactor("update", "namespaces", func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
//			return true, nil, assert.AnError // Return an error
//		})
//
//		p := k8s.DeployConfigMapInputParams{
//			ConfigMap: namespace,
//			Client:    clientSet.CoreV1().ConfigMaps(),
//		}
//
//		result, err := k8s.ConfigMapService().DeployConfigMap(context.Background(), p)
//		assert := assert.New(t)
//		assert.Error(err)
//		assert.Nil(result)
//	})
//}
