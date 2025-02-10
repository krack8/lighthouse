package k8s

//import (
//	"github.com/krack8/lighthouse/pkg/log"
//	"github.com/stretchr/testify/suite"
//	corev1 "k8s.io/api/core/v1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/client-go/kubernetes/fake"
//	"testing"
//)
//
//type ConfigMapTestSuite struct {
//	suite.Suite
//	clientSet *fake.Clientset
//}
//
//func TestConfigMap(t *testing.T) {
//	suite.Run(t, new(ConfigMapTestSuite))
//}
//
//func (s *ConfigMapTestSuite) SetupSuite() { // Setup for the entire suite
//	log.InitializeTestLogger()
//	// Initialize Test logger once for the suite
//}
//
//func (s *ConfigMapTestSuite) TearDownSuite() { // Teardown for the entire suite
//	// Close resources, etc. if needed.
//}
//
//func (s *ConfigMapTestSuite) SetupTest() {
//	s.clientSet = fake.NewClientset(&corev1.ConfigMap{
//		ObjectMeta: metav1.ObjectMeta{
//			Name: "test-namespace",
//		},
//	})
//	// Setup before each test
//	// Reset mocks, clear databases, etc. if needed for *each* test
//}
//
//func (s *ConfigMapTestSuite) TearDownTest() { // Teardown after each test
//	// Clean up after *each* test runs
//}
//
//func (s *ConfigMapTestSuite) TestGetConfigMapDetailsSuccess() {
//	p := k8s.GetConfigMapDetailsInputParams{
//		ConfigMapName: "test-namespace",
//		Client:        s.clientSet.CoreV1().ConfigMaps(), // Use the clientSet from SetupSuite
//	}
//	expectedConfigMap := corev1.ConfigMap{
//		ObjectMeta: metav1.ObjectMeta{
//			Name: "test-namespace",
//		},
//	}
//	expectedConfigMap = removeConfigMapFields(expectedConfigMap)
//
//	result, err := k8s.ConfigMapService().GetConfigMapDetails(context.Background(), p)
//	assert := assert.New(s.T())
//	assert.NoError(err)
//	assert.NotNil(result)
//
//	response, ok := result.(k8s.ResponseDTO)
//	assert.True(ok)
//	assert.Equal("success", response.Status)
//	assert.Equal(expectedConfigMap, response.Data)
//}
//
//func (s *ConfigMapTestSuite) TestGetConfigMapDetailsError() {
//	p := k8s.GetConfigMapInputParams{
//		ConfigMapName: "non-existent-namespace",
//		Client:        s.clientSet.CoreV1().ConfigMaps(), // Use the clientSet from SetupSuite
//	}
//
//	result, err := k8s.ConfigMapService().GetConfigMapDetails(context.Background(), p)
//	assert := assert.New(s.T())
//	assert.Error(err)
//	assert.Nil(result)
//}
//
//// A dummy removeConfigMapFields function for testing purposes
//func removeConfigMapFields(ns corev1.ConfigMap) corev1.ConfigMap {
//	return ns
//}
//
//func (s *ConfigMapTestSuite) TestDeleteConfigMapSuccess() {
//
//	p := k8s.DeleteConfigMapInputParams{
//		ConfigMapName: "test-namespace",
//		Client:        s.clientSet.CoreV1().ConfigMaps(),
//	}
//
//	result, err := k8s.ConfigMapService().DeleteConfigMap(context.Background(), p)
//	assert := assert.New(s.T())
//	assert.NoError(err)
//	assert.NotNil(result)
//
//	response, ok := result.(k8s.ResponseDTO)
//	assert.True(ok)
//	assert.Equal("success", response.Status)
//	assert.Equal("deleted namespace test-namespace", response.Msg)
//
//	_, err = s.clientSet.CoreV1().ConfigMaps().Get(context.Background(), "test-namespace", metav1.GetOptions{})
//	assert.Error(err)
//
//	s.T().Run("Test ConfigMap Error", func(t *testing.T) {
//		t.Log("Test CASE: Delete ConfigMap with non-existent namespace")
//		p := k8s.DeleteConfigMapInputParams{
//			ConfigMapName: "non-existent-namespace",
//			Client:        s.clientSet.CoreV1().ConfigMaps(),
//		}
//
//		result, err := k8s.ConfigMapService().DeleteConfigMap(context.Background(), p)
//		assert.Error(err)
//		assert.Nil(result)
//	})
//}
//
//func (s *ConfigMapTestSuite) TestDeleteConfigMapError() {
//	p := k8s.DeleteConfigMapInputParams{
//		ConfigMapName: "non-existent-namespace",
//		Client:        s.clientSet.CoreV1().ConfigMaps(),
//	}
//
//	result, err := k8s.ConfigMapService().DeleteConfigMap(context.Background(), p)
//
//	assert := assert.New(s.T())
//	assert.Error(err)
//	assert.Nil(result)
//}
//
//func (s *ConfigMapTestSuite) TestDeployConfigMap() {
//	s.T().Run("Success CreateConfigMap", func(t *testing.T) {
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
