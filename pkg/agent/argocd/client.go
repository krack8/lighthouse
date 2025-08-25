package argocd

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// RESTClient is the main ArgoCD client using REST API
type RESTClient struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// Config holds the ArgoCD client configuration
type Config struct {
	ServerAddr string
	AuthToken  string
	Insecure   bool
}

// NewClientFromEnv creates a new client from environment variables
func NewClientFromEnv() (*RESTClient, error) {
	config := &Config{
		ServerAddr: os.Getenv("ARGOCD_SERVER"),
		AuthToken:  os.Getenv("ARGOCD_AUTH_TOKEN"),
		Insecure:   os.Getenv("ARGOCD_INSECURE") == "true",
	}

	if config.ServerAddr == "" {
		return nil, fmt.Errorf("ARGOCD_SERVER environment variable is required")
	}

	if config.AuthToken == "" {
		return nil, fmt.Errorf("ARGOCD_AUTH_TOKEN environment variable is required")
	}

	return NewClient(config)
}

// NewClient creates a new ArgoCD REST client
func NewClient(config *Config) (*RESTClient, error) {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	if config.Insecure {
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	return &RESTClient{
		baseURL:    fmt.Sprintf("https://%s/api/v1", config.ServerAddr),
		token:      config.AuthToken,
		httpClient: httpClient,
	}, nil
}

// NewRESTClient creates a new REST client (alternative constructor for compatibility)
func NewRESTClient(serverURL, token string, insecure bool) *RESTClient {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	if insecure {
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	return &RESTClient{
		baseURL:    fmt.Sprintf("https://%s/api/v1", serverURL),
		token:      token,
		httpClient: httpClient,
	}
}

// doRequest performs an HTTP request to the ArgoCD API
func (c *RESTClient) doRequest(method, path string, body interface{}) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, c.baseURL+path, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %s - %s", resp.Status, string(responseBody))
	}

	return responseBody, nil
}

// Close closes the HTTP client (implements io.Closer interface if needed)
func (c *RESTClient) Close() error {
	// HTTP client doesn't need explicit closing
	// This method is here for interface compatibility
	return nil
}

// GetBaseURL returns the base URL of the ArgoCD server
func (c *RESTClient) GetBaseURL() string {
	return c.baseURL
}

// GetToken returns the authentication token (useful for debugging, use carefully)
func (c *RESTClient) GetToken() string {
	return c.token
}
