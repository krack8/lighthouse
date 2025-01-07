package helm

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
	"os"
	"strings"
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

func (h *HelmClient) InitializeHelmRepo() error {
	repoFile := h.settings.RepositoryConfig
	_, err := repo.LoadFile(repoFile)
	if err != nil {
		// If the file does not exist, initialize Helm by creating the repositories file
		if os.IsNotExist(err) || strings.HasPrefix(err.Error(), "couldn't load repositories file") {
			// Create a new repositories file with a default entry
			repositories := repo.NewFile()

			// Add the default "stable" repository (or another default repository)
			defaultRepo := &repo.Entry{
				Name: "stable",
				URL:  "https://charts.helm.sh/stable",
			}
			repositories.Add(defaultRepo)

			// Write the new repositories file to the default location
			if err := repositories.WriteFile(repoFile, 0644); err != nil {
				return fmt.Errorf("failed to write repositories file: %w", err)
			}

			// Download the index file for the default repository
			chartRepo, err := repo.NewChartRepository(defaultRepo, getter.All(h.settings))
			if err != nil {
				return fmt.Errorf("failed to create chart repository: %w", err)
			}

			if _, err := chartRepo.DownloadIndexFile(); err != nil {
				return fmt.Errorf("failed to download index file for default repository: %w", err)
			}

			fmt.Println("Helm initialized with the default stable repository.")

			return nil
		}
	}
	return err
}

// ListReleases lists all installed releases in the namespace
func (h *HelmClient) ListReleases(allNamespaces bool) ([]*release.Release, error) {
	list := action.NewList(h.config)
	list.AllNamespaces = allNamespaces
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

// InstallChart installs a Helm chart with the specified repo, chart, and version
func (h *HelmClient) InstallChart(repoName, chartName, chartVersion, releaseName, namespace string, values map[string]interface{}) (*release.Release, error) {
	// Load the repository configuration
	repoFile := h.settings.RepositoryConfig
	repositories, err := repo.LoadFile(repoFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load repositories file: %w", err)
	}

	// Find the repository entry
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

	// Download the index file
	indexFilePath, err := chartRepo.DownloadIndexFile()
	if err != nil {
		return nil, fmt.Errorf("failed to download index file: %w", err)
	}

	// Load the index file to find the chart
	indexFile, err := repo.LoadIndexFile(indexFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load index file: %w", err)
	}

	chartVersions, ok := indexFile.Entries[chartName]
	if !ok {
		return nil, fmt.Errorf("chart %q not found in repository %q", chartName, repoName)
	}

	// Find the specific version
	var selectedChart *repo.ChartVersion
	for _, version := range chartVersions {
		if version.Version == chartVersion {
			selectedChart = version
			break
		}
	}

	if selectedChart == nil {
		return nil, fmt.Errorf("version %q of chart %q not found in repository %q", chartVersion, chartName, repoName)
	}

	// Prepare the install action
	client := action.NewInstall(h.config)
	client.ReleaseName = releaseName
	client.Namespace = namespace
	client.Version = chartVersion

	// Construct the chart URL and load it
	chartURL := selectedChart.URLs[0] // URL for the chart tarball
	chartPath, err := client.ChartPathOptions.LocateChart(chartURL, h.settings)
	if err != nil {
		return nil, fmt.Errorf("failed to locate chart: %w", err)
	}

	chart, err := loader.Load(chartPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load chart: %w", err)
	}

	// Run the install action
	release, err := client.Run(chart, values)
	if err != nil {
		return nil, fmt.Errorf("failed to install chart: %w", err)
	}

	return release, nil
}

// ListRepos lists all configured Helm repositories
func (h *HelmClient) ListRepos() ([]*repo.Entry, error) {
	repoFile := h.settings.RepositoryConfig

	// Load the repositories file
	repositories, err := repo.LoadFile(repoFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load repositories file: %w", err)
	}

	return repositories.Repositories, nil
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

func (h *HelmClient) ListChartsInRepo(repoName string) (map[string]repo.ChartVersions, error) {
	repoFile := h.settings.RepositoryConfig

	// Load the repositories file
	repositories, err := repo.LoadFile(repoFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load repositories file: %w", err)
	}

	// Find the specified repository
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

	// Download the index file
	indexFilePath, err := chartRepo.DownloadIndexFile()
	if err != nil {
		return nil, fmt.Errorf("failed to download index file: %w", err)
	}

	// Load the downloaded index file
	indexFile, err := repo.LoadIndexFile(indexFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load index file: %w", err)
	}

	// Return the charts (keyed by chart name, each containing its versions)
	return indexFile.Entries, nil
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
