package k8s_test

import (
	"context"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	k8stesting "k8s.io/client-go/testing"
	"testing"
)

type NoopLogger struct{}

func (l NoopLogger) Debugw(msg string, keysAndValues ...interface{}) {}
func (l NoopLogger) Errorw(msg string, keysAndValues ...interface{}) {}

type NamespaceTestSuite struct {
	suite.Suite
	clientSet *fake.Clientset
}

func TestNamespace(t *testing.T) {
	suite.Run(t, new(NamespaceTestSuite))
}

func (s *NamespaceTestSuite) SetupSuite() { // Setup for the entire suite
	log.InitializeTestLogger()
	// Initialize Test logger once for the suite
}

func (s *NamespaceTestSuite) TearDownSuite() { // Teardown for the entire suite
	// Close resources, etc. if needed.
}

func (s *NamespaceTestSuite) SetupTest() {
	s.clientSet = fake.NewClientset(&corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-namespace",
		},
	})
	// Setup before each test
	// Reset mocks, clear databases, etc. if needed for *each* test
}

func (s *NamespaceTestSuite) TearDownTest() { // Teardown after each test
	// Clean up after *each* test runs
}

func (s *NamespaceTestSuite) TestGetNamespaceDetailsSuccess() {
	p := k8s.GetNamespaceInputParams{
		NamespaceName: "test-namespace",
		Client:        s.clientSet.CoreV1().Namespaces(), // Use the clientSet from SetupSuite
	}
	expectedNamespace := corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-namespace",
		},
	}
	expectedNamespace = removeNamespaceFields(expectedNamespace)

	result, err := k8s.NamespaceService().GetNamespaceDetails(context.Background(), p)
	assert := assert.New(s.T())
	assert.NoError(err)
	assert.NotNil(result)

	response, ok := result.(k8s.ResponseDTO)
	assert.True(ok)
	assert.Equal("success", response.Status)
	assert.Equal(expectedNamespace, response.Data)
}

func (s *NamespaceTestSuite) TestGetNamespaceDetailsError() {
	p := k8s.GetNamespaceInputParams{
		NamespaceName: "non-existent-namespace",
		Client:        s.clientSet.CoreV1().Namespaces(), // Use the clientSet from SetupSuite
	}

	result, err := k8s.NamespaceService().GetNamespaceDetails(context.Background(), p)
	assert := assert.New(s.T())
	assert.Error(err)
	assert.Nil(result)
}

// A dummy removeNamespaceFields function for testing purposes
func removeNamespaceFields(ns corev1.Namespace) corev1.Namespace {
	return ns
}

func (s *NamespaceTestSuite) TestDeleteNamespaceSuccess() {

	p := k8s.DeleteNamespaceInputParams{
		NamespaceName: "test-namespace",
		Client:        s.clientSet.CoreV1().Namespaces(),
	}

	result, err := k8s.NamespaceService().DeleteNamespace(context.Background(), p)
	assert := assert.New(s.T())
	assert.NoError(err)
	assert.NotNil(result)

	response, ok := result.(k8s.ResponseDTO)
	assert.True(ok)
	assert.Equal("success", response.Status)
	assert.Equal("deleted namespace test-namespace", response.Msg)

	_, err = s.clientSet.CoreV1().Namespaces().Get(context.Background(), "test-namespace", metav1.GetOptions{})
	assert.Error(err)

	s.T().Run("Test Namespace Error", func(t *testing.T) {
		t.Log("Test CASE: Delete Namespace with non-existent namespace")
		p := k8s.DeleteNamespaceInputParams{
			NamespaceName: "non-existent-namespace",
			Client:        s.clientSet.CoreV1().Namespaces(),
		}

		result, err := k8s.NamespaceService().DeleteNamespace(context.Background(), p)
		assert.Error(err)
		assert.Nil(result)
	})
}

func (s *NamespaceTestSuite) TestDeleteNamespaceError() {
	p := k8s.DeleteNamespaceInputParams{
		NamespaceName: "non-existent-namespace",
		Client:        s.clientSet.CoreV1().Namespaces(),
	}

	result, err := k8s.NamespaceService().DeleteNamespace(context.Background(), p)

	assert := assert.New(s.T())
	assert.Error(err)
	assert.Nil(result)
}

func (s *NamespaceTestSuite) TestDeployNamespace() {
	s.T().Run("Success CreateNamespace", func(t *testing.T) {
		namespace := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "new-namespace",
			},
			TypeMeta: metav1.TypeMeta{
				Kind:       "Namespace",
				APIVersion: "v1",
			},
		}

		p := k8s.DeployNamespaceInputParams{
			Namespace: namespace,
			Client:    s.clientSet.CoreV1().Namespaces(),
		}

		result, err := k8s.NamespaceService().DeployNamespace(context.Background(), p)
		assert := assert.New(t)
		assert.NoError(err)
		assert.NotNil(result)

		response, ok := result.(k8s.ResponseDTO)
		assert.True(ok)
		assert.Equal("success", response.Status)
		assert.Equal(namespace, response.Data)
		// Verify namespace exists
		fetchedNamespace, err := s.clientSet.CoreV1().Namespaces().Get(context.Background(), namespace.Name, metav1.GetOptions{})
		assert.NoError(err)
		assert.Equal(namespace.Name, fetchedNamespace.Name)
	})

	s.T().Run("Success UpdateNamespace", func(t *testing.T) {
		updatedLabels := map[string]string{"updated": "true"}
		namespace := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name:   "test-namespace",
				Labels: updatedLabels,
			},
		}

		p := k8s.DeployNamespaceInputParams{
			Namespace: namespace,
			Client:    s.clientSet.CoreV1().Namespaces(),
		}
		result, err := k8s.NamespaceService().DeployNamespace(context.Background(), p)
		assert := assert.New(t)
		assert.NoError(err)
		assert.NotNil(result)

		response, ok := result.(k8s.ResponseDTO)
		assert.True(ok)
		assert.Equal("success", response.Status)
		assert.Equal(namespace, response.Data)

		// Verify namespace was updated
		fetchedNamespace, err := s.clientSet.CoreV1().Namespaces().Get(context.Background(), "test-namespace", metav1.GetOptions{})
		assert.NoError(err)
		assert.Equal(updatedLabels, fetchedNamespace.Labels)
	})
	s.T().Run("Error CreateNamespace - Create Fails", func(t *testing.T) {
		namespace := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "new-namespace",
			},
		}
		clientSet := fake.NewClientset()
		// Make the fake client return an error on create
		clientSet.PrependReactor("create", "namespaces", func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
			return true, nil, assert.AnError // Return an error
		})

		p := k8s.DeployNamespaceInputParams{
			Namespace: namespace,
			Client:    clientSet.CoreV1().Namespaces(),
		}

		result, err := k8s.NamespaceService().DeployNamespace(context.Background(), p)
		assert := assert.New(t)
		assert.Error(err)
		assert.Nil(result)
	})
	s.T().Run("Error UpdateNamespace - Update Fails", func(t *testing.T) {
		namespace := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-namespace",
			},
		}

		clientSet := fake.NewClientset(&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-namespace",
			},
		})
		// Make the fake client return an error on update
		clientSet.PrependReactor("update", "namespaces", func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
			return true, nil, assert.AnError // Return an error
		})

		p := k8s.DeployNamespaceInputParams{
			Namespace: namespace,
			Client:    clientSet.CoreV1().Namespaces(),
		}

		result, err := k8s.NamespaceService().DeployNamespace(context.Background(), p)
		assert := assert.New(t)
		assert.Error(err)
		assert.Nil(result)
	})
}
