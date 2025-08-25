package argocd

import (
	"encoding/json"
	"fmt"
)

// ListProjects retrieves all projects
func (c *RESTClient) ListProjects() (*ProjectList, error) {
	data, err := c.doRequest("GET", "/projects", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}

	var projects ProjectList
	if err := json.Unmarshal(data, &projects); err != nil {
		return nil, err
	}

	return &projects, nil
}

// GetProject retrieves a specific project by name
func (c *RESTClient) GetProject(name string) (*Project, error) {
	path := fmt.Sprintf("/projects/%s", name)

	data, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get project %s: %w", name, err)
	}

	var project Project
	if err := json.Unmarshal(data, &project); err != nil {
		return nil, err
	}

	return &project, nil
}

// CreateProject creates a new project
func (c *RESTClient) CreateProject(spec *ProjectSpec) (*Project, error) {
	project := &Project{
		Metadata: ProjectMetadata{
			Name: spec.Name,
		},
		Spec: *spec,
	}

	data, err := c.doRequest("POST", "/projects", project)
	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	var created Project
	if err := json.Unmarshal(data, &created); err != nil {
		return nil, err
	}

	return &created, nil
}

// UpdateProject updates an existing project
func (c *RESTClient) UpdateProject(name string, spec *ProjectSpec) (*Project, error) {
	project := &Project{
		Metadata: ProjectMetadata{
			Name: name,
		},
		Spec: *spec,
	}

	path := fmt.Sprintf("/projects/%s", name)
	data, err := c.doRequest("PUT", path, project)
	if err != nil {
		return nil, fmt.Errorf("failed to update project %s: %w", name, err)
	}

	var updated Project
	if err := json.Unmarshal(data, &updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

// DeleteProject deletes a project
func (c *RESTClient) DeleteProject(name string) error {
	path := fmt.Sprintf("/projects/%s", name)
	_, err := c.doRequest("DELETE", path, nil)
	if err != nil {
		return fmt.Errorf("failed to delete project %s: %w", name, err)
	}
	return nil
}

// GetProjectRole gets a specific role in a project
func (c *RESTClient) GetProjectRole(projectName, roleName string) (*ProjectRole, error) {
	path := fmt.Sprintf("/projects/%s/roles/%s", projectName, roleName)

	data, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get role %s in project %s: %w", roleName, projectName, err)
	}

	var role ProjectRole
	if err := json.Unmarshal(data, &role); err != nil {
		return nil, err
	}

	return &role, nil
}

// CreateProjectToken creates a new token for a project role
func (c *RESTClient) CreateProjectToken(projectName, roleName string, expiresIn string) (string, error) {
	path := fmt.Sprintf("/projects/%s/roles/%s/token", projectName, roleName)

	tokenRequest := map[string]string{
		"expiresIn": expiresIn,
	}

	data, err := c.doRequest("POST", path, tokenRequest)
	if err != nil {
		return "", fmt.Errorf("failed to create token: %w", err)
	}

	var response struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return "", err
	}

	return response.Token, nil
}
