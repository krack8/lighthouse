package argocd

import (
	"encoding/json"
	"fmt"
)

// ListApplications retrieves all applications, optionally filtered by project
func (c *RESTClient) ListApplications(project string) (*ApplicationList, error) {
	path := "/applications"
	if project != "" {
		path = fmt.Sprintf("%s?project=%s", path, project)
	}

	data, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list applications: %w", err)
	}

	var apps ApplicationList
	if err := json.Unmarshal(data, &apps); err != nil {
		return nil, err
	}

	return &apps, nil
}

// GetApplication retrieves a specific application by name
func (c *RESTClient) GetApplication(name string) (*Application, error) {
	path := fmt.Sprintf("/applications/%s", name)

	data, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get application %s: %w", name, err)
	}

	var app Application
	if err := json.Unmarshal(data, &app); err != nil {
		return nil, err
	}

	return &app, nil
}

// CreateApplication creates a new application
func (c *RESTClient) CreateApplication(spec *ApplicationSpec) (*Application, error) {
	app := &Application{
		Metadata: ApplicationMetadata{
			Name:      spec.Name,
			Namespace: spec.Namespace,
		},
		Spec: *spec,
	}

	data, err := c.doRequest("POST", "/applications", app)
	if err != nil {
		return nil, fmt.Errorf("failed to create application: %w", err)
	}

	var created Application
	if err := json.Unmarshal(data, &created); err != nil {
		return nil, err
	}

	return &created, nil
}

// UpdateApplication updates an existing application
func (c *RESTClient) UpdateApplication(name string, spec *ApplicationSpec) (*Application, error) {
	app := &Application{
		Metadata: ApplicationMetadata{
			Name:      name,
			Namespace: spec.Namespace,
		},
		Spec: *spec,
	}

	path := fmt.Sprintf("/applications/%s", name)
	data, err := c.doRequest("PUT", path, app)
	if err != nil {
		return nil, fmt.Errorf("failed to update application %s: %w", name, err)
	}

	var updated Application
	if err := json.Unmarshal(data, &updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

// DeleteApplication deletes an application
func (c *RESTClient) DeleteApplication(name string, cascade bool) error {
	path := fmt.Sprintf("/applications/%s?cascade=%t", name, cascade)
	_, err := c.doRequest("DELETE", path, nil)
	if err != nil {
		return fmt.Errorf("failed to delete application %s: %w", name, err)
	}
	return nil
}

// SyncApplication syncs an application with its source
func (c *RESTClient) SyncApplication(name string, revision string, prune bool, dryRun bool) (*Application, error) {
	path := fmt.Sprintf("/applications/%s/sync", name)

	syncRequest := map[string]interface{}{
		"revision": revision,
		"prune":    prune,
		"dryRun":   dryRun,
	}

	data, err := c.doRequest("POST", path, syncRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to sync application %s: %w", name, err)
	}

	var app Application
	if err := json.Unmarshal(data, &app); err != nil {
		return nil, err
	}

	return &app, nil
}

// RollbackApplication rolls back an application to a previous revision
func (c *RESTClient) RollbackApplication(name string, revision string) (*Application, error) {
	path := fmt.Sprintf("/applications/%s/rollback", name)

	rollbackRequest := map[string]interface{}{
		"id": revision,
	}

	data, err := c.doRequest("POST", path, rollbackRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to rollback application %s: %w", name, err)
	}

	var app Application
	if err := json.Unmarshal(data, &app); err != nil {
		return nil, err
	}

	return &app, nil
}

// GetApplicationResources gets application resources
func (c *RESTClient) GetApplicationResources(name string) ([]ResourceStatus, error) {
	path := fmt.Sprintf("/applications/%s/resource-tree", name)

	data, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get resources for application %s: %w", name, err)
	}

	var tree struct {
		Nodes []ResourceStatus `json:"nodes"`
	}
	if err := json.Unmarshal(data, &tree); err != nil {
		return nil, err
	}

	return tree.Nodes, nil
}

// GetApplicationLogs gets logs for an application resource
func (c *RESTClient) GetApplicationLogs(appName, namespace, podName, container string) (string, error) {
	path := fmt.Sprintf("/applications/%s/pods/%s/logs?namespace=%s&container=%s",
		appName, podName, namespace, container)

	data, err := c.doRequest("GET", path, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get logs: %w", err)
	}

	return string(data), nil
}

// TerminateOperation terminates the current operation
func (c *RESTClient) TerminateOperation(appName string) error {
	path := fmt.Sprintf("/applications/%s/operation", appName)
	_, err := c.doRequest("DELETE", path, nil)
	return err
}
