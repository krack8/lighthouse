package argocd

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// ListRepositories retrieves all repositories
func (c *RESTClient) ListRepositories() (*RepositoryList, error) {
	data, err := c.doRequest("GET", "/repositories", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list repositories: %w", err)
	}

	var repos RepositoryList
	if err := json.Unmarshal(data, &repos); err != nil {
		return nil, err
	}

	return &repos, nil
}

// GetRepository retrieves a specific repository
func (c *RESTClient) GetRepository(repoURL string) (*Repository, error) {
	path := fmt.Sprintf("/repositories/%s", url.QueryEscape(repoURL))

	data, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get repository %s: %w", repoURL, err)
	}

	var repo Repository
	if err := json.Unmarshal(data, &repo); err != nil {
		return nil, err
	}

	return &repo, nil
}

// CreateRepository creates a new repository
func (c *RESTClient) CreateRepository(repo *Repository) (*Repository, error) {
	data, err := c.doRequest("POST", "/repositories", repo)
	if err != nil {
		return nil, fmt.Errorf("failed to create repository: %w", err)
	}

	var created Repository
	if err := json.Unmarshal(data, &created); err != nil {
		return nil, err
	}

	return &created, nil
}

// UpdateRepository updates an existing repository
func (c *RESTClient) UpdateRepository(repoURL string, repo *Repository) (*Repository, error) {
	path := fmt.Sprintf("/repositories/%s", url.QueryEscape(repoURL))

	data, err := c.doRequest("PUT", path, repo)
	if err != nil {
		return nil, fmt.Errorf("failed to update repository: %w", err)
	}

	var updated Repository
	if err := json.Unmarshal(data, &updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

// DeleteRepository deletes a repository
func (c *RESTClient) DeleteRepository(repoURL string) error {
	path := fmt.Sprintf("/repositories/%s", url.QueryEscape(repoURL))
	_, err := c.doRequest("DELETE", path, nil)
	if err != nil {
		return fmt.Errorf("failed to delete repository %s: %w", repoURL, err)
	}
	return nil
}

// ValidateRepository validates repository credentials
func (c *RESTClient) ValidateRepository(repo *Repository) error {
	_, err := c.doRequest("POST", "/repositories/validate", repo)
	if err != nil {
		return fmt.Errorf("repository validation failed: %w", err)
	}
	return nil
}

// ListRepositoryApps lists all apps in a repository
func (c *RESTClient) ListRepositoryApps(repoURL, revision string) ([]string, error) {
	path := fmt.Sprintf("/repositories/%s/apps?revision=%s",
		url.QueryEscape(repoURL), revision)

	data, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list apps in repository: %w", err)
	}

	var apps struct {
		Items []string `json:"items"`
	}
	if err := json.Unmarshal(data, &apps); err != nil {
		return nil, err
	}

	return apps.Items, nil
}
