package helm

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
)

type HelmClient struct {
	settings *cli.EnvSettings
	config   *action.Configuration
}

func NewHelmClient(namespace string) (*HelmClient, error) {
	settings := cli.New()
	config := new(action.Configuration)
	if err := config.Init(settings.RESTClientGetter(), namespace, "memory", debugLog); err != nil {
		return nil, err
	}

	return &HelmClient{
		settings: settings,
		config:   config,
	}, nil
}

// ListReleases lists all installed releases in the namespace
func (h *HelmClient) ListReleases() ([]*release.Release, error) {
	list := action.NewList(h.config)
	list.AllNamespaces = false
	return list.Run()
}

// InstallOrUpgradeChart installs or upgrades a Helm chart
func (h *HelmClient) InstallOrUpgradeChart(chartPath, releaseName string, values map[string]interface{}) (*release.Release, error) {
	client := action.NewInstall(h.config)
	client.ReleaseName = releaseName
	client.Namespace = h.settings.Namespace()

	chart, err := loader.Load(chartPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load chart: %w", err)
	}

	return client.Run(chart, values)
}

// UninstallChart uninstalls a Helm release
func (h *HelmClient) UninstallChart(releaseName string) (*release.UninstallReleaseResponse, error) {
	client := action.NewUninstall(h.config)
	return client.Run(releaseName)
}

// AddRepo adds a new Helm repository
func (h *HelmClient) AddRepo(name, url string) error {
	repoFile := h.settings.RepositoryConfig
	repositories, err := repo.LoadFile(repoFile)
	if err != nil {
		return fmt.Errorf("failed to load repositories file: %w", err)
	}

	entry := &repo.Entry{
		Name: name,
		URL:  url,
	}

	repo, err := repo.NewChartRepository(entry, getter.All(h.settings))
	if err != nil {
		return fmt.Errorf("failed to create chart repository: %w", err)
	}

	if _, err := repo.DownloadIndexFile(); err != nil {
		return fmt.Errorf("failed to download index file: %w", err)
	}

	repositories.Update(entry)
	return repositories.WriteFile(repoFile, 0644)
}

// RemoveRepo removes a Helm repository
func (h *HelmClient) RemoveRepo(name string) error {
	repoFile := h.settings.RepositoryConfig
	repositories, err := repo.LoadFile(repoFile)
	if err != nil {
		return fmt.Errorf("failed to load repositories file: %w", err)
	}

	repositories.Remove(name)
	return repositories.WriteFile(repoFile, 0644)
}

// SearchChart searches for a chart in a Helm repository by name
func (h *HelmClient) SearchChart(repoName, chartName string) ([]*repo.ChartVersion, error) {
	repoFile := h.settings.RepositoryConfig
	repositories, err := repo.LoadFile(repoFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load repository file: %w", err)
	}

	var foundEntry *repo.Entry
	for _, entry := range repositories.Repositories {
		if entry.Name == repoName {
			foundEntry = entry
			break
		}
	}

	if foundEntry == nil {
		return nil, fmt.Errorf("repository %q not found", repoName)
	}

	// Create a chart repository object
	chartRepo, err := repo.NewChartRepository(foundEntry, getter.All(h.settings))
	if err != nil {
		return nil, fmt.Errorf("failed to create chart repository: %w", err)
	}

	var indexFilePath string
	if indexFilePath, err = chartRepo.DownloadIndexFile(); err != nil {
		return nil, fmt.Errorf("failed to download index file: %w", err)
	}

	// Load the downloaded index file
	indexFile, err := repo.LoadIndexFile(indexFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load index file: %w", err)
	}

	// Search for the chart in the index file
	var results []*repo.ChartVersion
	for name, versions := range indexFile.Entries {
		if name == chartName {
			results = versions
			break
		}
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("chart %q not found in repository %q", chartName, repoName)
	}

	return results, nil
}

// ListAllReleases lists all Helm releases in all namespaces
func (h *HelmClient) ListAllReleases() ([]*release.Release, error) {
	client := action.NewList(h.config)
	client.AllNamespaces = true // Enable listing releases across all namespaces

	releases, err := client.Run()
	if err != nil {
		return nil, err
	}
	return releases, nil
}

// GetReleaseDetails fetches details of a specific release
func (h *HelmClient) GetReleaseDetails(releaseName string) (*release.Release, error) {
	client := action.NewGet(h.config)
	return client.Run(releaseName)
}

// ListRevisions lists all revisions of a given Helm release
func (h *HelmClient) ListRevisions(releaseName string) ([]*release.Release, error) {
	client := action.NewHistory(h.config)
	client.Max = 0 // Set to 0 for no limit on the number of revisions returned

	revisions, err := client.Run(releaseName)
	if err != nil {
		return nil, fmt.Errorf("failed to get revisions for release %s: %w", releaseName, err)
	}
	return revisions, nil
}

// RollbackRelease rolls back a release to a specified revision
func (h *HelmClient) RollbackRelease(releaseName string, revision int) error {
	client := action.NewRollback(h.config)
	client.Version = revision // Specify the revision to roll back to
	return client.Run(releaseName)
}

func debugLog(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
