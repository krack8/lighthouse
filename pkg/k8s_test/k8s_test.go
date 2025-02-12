package k8s_test

//func TestListNamespaces(t *testing.T) {
//	// Mock the Kubernetes client using a fake client
//	clientset := fake.NewClientset()
//
//	// Create mock namespaces
//	namespaceNames := []string{"default", "kube-system", "test-namespace"}
//	for _, ns := range namespaceNames {
//		_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{
//			ObjectMeta: metav1.ObjectMeta{Name: ns},
//		}, metav1.CreateOptions{})
//		if err != nil {
//			t.Fatalf("Failed to create namespace: %v", err)
//		}
//	}
//
//	// Fetch namespaces using FetchNamespaces function
//	fetchedNamespaces, err := k8s.ListNamespaces(clientset)
//	if err != nil {
//		t.Fatalf("FetchNamespaces failed: %v", err)
//	}
//
//	// Validate the fetched namespaces
//	for _, expected := range namespaceNames {
//		found := false
//		for _, actual := range fetchedNamespaces {
//			if expected == actual {
//				found = true
//				break
//			}
//		}
//		if !found {
//			t.Errorf("Expected namespace %q not found", expected)
//		}
//	}
//}
//
//var (
//	BaseUrl        = "http://127.0.0.1:" + _cfg.ServerPort
//	HealthEndpoint = "/health"
//	IndexEndpoint  = "/"
//)
//
//type EndToEndSuite struct {
//	suite.Suite
//	client *http.Client
//}
//
//func TestEndToEnd(t *testing.T) {
//	suite.Run(t, new(EndToEndSuite))
//}
//
//func (s *EndToEndSuite) SetupSuite() {
//	s.client = &http.Client{}
//}
//
//func (s *EndToEndSuite) SetupTest() {
//	// Set up before *each* test runs (e.g., reset mocks, clear databases)
//}
//
//func (s *EndToEndSuite) TearDownTest() {
//	// Clean up after *each* test runs
//}
//
//func (s *EndToEndSuite) TestHealthCheck() {
//	resp, _ := s.client.Get(BaseUrl + HealthEndpoint)
//	defer resp.Body.Close()
//	s.Equal(resp.StatusCode, http.StatusOK)
//	body, _ := io.ReadAll(resp.Body)
//	s.Equal("I am live!", string(body))
//}
//
//func (s *EndToEndSuite) TestHealthIndex() {
//	resp, _ := s.client.Get(BaseUrl + IndexEndpoint)
//	defer resp.Body.Close()
//	s.Equal(resp.StatusCode, http.StatusOK)
//	body, _ := io.ReadAll(resp.Body)
//	s.Equal("This is KloverCloud Lighthouse", string(body))
//}
