package helm

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
	"k8s.io/client-go/kubernetes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type HelmClient struct {
	settings   *cli.EnvSettings
	config     *action.Configuration
	kubeClient *kubernetes.Clientset
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

// InstallOrUpgradeChart installs a Helm chart with the specified repo, chart, and version
func (h *HelmClient) InstallOrUpgradeChart(repoName, chartName, chartVersion, releaseName, namespace string, values map[string]interface{}) (*release.Release, error) {
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

// ListChartsInRepo list all charts and versions of a repo
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

// GetChartValues fetches the values of a chart
func (h *HelmClient) GetChartValues(repoName, chartName, chartVersion string) (map[string]interface{}, error) {
	// Load the repository configuration
	repoFile := h.settings.RepositoryConfig
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

	// Initialize the repository
	chartRepo, err := repo.NewChartRepository(foundEntry, getter.All(h.settings))
	if err != nil {
		return nil, fmt.Errorf("failed to create chart repository: %w", err)
	}

	// Download the index file to locate the chart
	tempDir, err := os.MkdirTemp("", "helm-repo")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	var indexFilePath string
	if indexFilePath, err = chartRepo.DownloadIndexFile(); err != nil {
		return nil, fmt.Errorf("failed to download index file: %w", err)
	}

	indexFile, err := repo.LoadIndexFile(indexFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load index file: %w", err)
	}

	// Find the chart version in the index
	chartVersions, exists := indexFile.Entries[chartName]
	if !exists {
		return nil, fmt.Errorf("chart %q not found in repository %q", chartName, repoName)
	}

	var chartURL string
	for _, version := range chartVersions {
		if version.Version == chartVersion {
			chartURL = version.URLs[0]
			break
		}
	}

	if chartURL == "" {
		return nil, fmt.Errorf("chart %q with version %q not found in repository %q", chartName, chartVersion, repoName)
	}

	// Download the chart archive
	chartPath := filepath.Join(tempDir, "chart.tgz")
	chartURL, err = repo.ResolveReferenceURL(foundEntry.URL, chartURL)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve chart url: %w", err)
	}

	// Use Helm's getter to download the chart
	err = h.DownloadChart(chartURL, chartPath)
	if err != nil {
		return nil, fmt.Errorf("failed to download chart: %w", err)
	}

	// Defer cleanup to delete the file after use
	defer func() {
		if err := os.Remove(chartPath); err != nil {
			fmt.Printf("Failed to delete downloaded chart: %v\n", err)
		}
	}()

	// Load the chart
	chart, err := loader.Load(chartPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load chart: %w", err)
	}

	// Return the chart's default values
	return chart.Values, nil
}

// DownloadChart downloads a Helm chart from a given URL and saves it to a specified path
func (h *HelmClient) DownloadChart(chartURL, chartPath string) error {
	// Get the getter providers
	getters := getter.All(h.settings)

	// Resolve the appropriate getter for the URL (e.g., http/https/file)
	provider, err := getters.ByScheme("http")
	if err != nil {
		return fmt.Errorf("failed to get provider: %w", err)
	}

	// Download the chart
	data, err := provider.Get(chartURL)
	if err != nil {
		return fmt.Errorf("failed to download chart: %w", err)
	}

	// Save the downloaded chart to the local path
	err = os.WriteFile(chartPath, data.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write chart file: %w", err)
	}

	return nil
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

// GetCurrentRevisionDetails returns the details of current revision
func (h *HelmClient) GetCurrentRevisionDetails(releaseName string) (*release.Release, error) {
	// Create a new Get action for fetching release details
	getAction := action.NewGet(h.config)

	// Fetch the current (latest) release details
	releaseDetails, err := getAction.Run(releaseName)
	if err != nil {
		return nil, fmt.Errorf("failed to get current release details for %s: %w", releaseName, err)
	}

	// Return the details of the current (latest) revision
	return releaseDetails, nil
}

// GetRevisionDetails returns the details of a revision
func (h *HelmClient) GetRevisionDetails(releaseName string, revision int) (*release.Release, error) {
	// Create a new History action for fetching release history
	historyAction := action.NewHistory(h.config)
	historyAction.Max = 0 // Retrieve all revisions (unlimited)

	// Fetch all the revisions for the given release
	revisions, err := historyAction.Run(releaseName)
	if err != nil {
		return nil, fmt.Errorf("failed to get history for release %s: %w", releaseName, err)
	}

	// Iterate through the revisions to find the requested revision
	var revisionDetails *release.Release
	for _, r := range revisions {
		if r.Version == revision {
			revisionDetails = r
			break
		}
	}

	// If the specified revision is not found, return an error
	if revisionDetails == nil {
		return nil, fmt.Errorf("revision %d not found for release %s", revision, releaseName)
	}

	// Return the details of the specified revision
	return revisionDetails, nil
}

// GetCurrentAppliedValues fetches the applied values of current revision
func (h *HelmClient) GetCurrentAppliedValues(releaseName string) (map[string]interface{}, error) {
	// Create a new Get action for fetching the release details
	getAction := action.NewGet(h.config)

	// Fetch the current release details
	releaseDetails, err := getAction.Run(releaseName)
	if err != nil {
		return nil, fmt.Errorf("failed to get release details for %s: %w", releaseName, err)
	}

	// Extract the applied values from the release details
	appliedValues := releaseDetails.Config

	return appliedValues, nil
}

// GetAppliedValues fetches the applied values of a specific revision
func (h *HelmClient) GetAppliedValues(releaseName string, revision int) (map[string]interface{}, error) {
	// Create a new Get action for fetching release details
	getAction := action.NewGet(h.config)

	// Fetch the release details for the specific revision
	releaseDetails, err := getAction.Run(releaseName)
	if err != nil {
		return nil, fmt.Errorf("failed to get release details for %s: %w", releaseName, err)
	}

	// Check if the requested revision matches
	if revision > 0 && releaseDetails.Version != revision {
		return nil, fmt.Errorf("release %s is not at revision %d", releaseName, revision)
	}

	// Extract the applied values
	appliedValues := releaseDetails.Config

	return appliedValues, nil
}

// RollbackRelease rolls back a release to a specified revision
func (h *HelmClient) RollbackRelease(releaseName string, revision int) error {
	client := action.NewRollback(h.config)
	client.Version = revision // Specify the revision to roll back to
	return client.Run(releaseName)
}

func (h *HelmClient) AddGitRepo(name, gitURL, username, password string) error {
	// Clone the Git repository
	tempDir, err := os.MkdirTemp("", "helm-git-repo")
	if err != nil {
		return fmt.Errorf("failed to create temporary directory: %w", err)
	}
	defer os.RemoveAll(tempDir) // Cleanup after use

	gitCloneCmd := []string{"clone", gitURL, tempDir}
	if username != "" && password != "" {
		gitURL = fmt.Sprintf("https://%s:%s@%s", username, password, gitURL[8:])
		gitCloneCmd = []string{"clone", gitURL, tempDir}
	}

	cmd := exec.Command("git", gitCloneCmd...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to clone Git repository: %w\nOutput: %s", err, string(output))
	}

	// Find the Chart.yaml to verify it's a valid chart repo
	chartPath := filepath.Join(tempDir, "Chart.yaml")
	if _, err := os.Stat(chartPath); os.IsNotExist(err) {
		return fmt.Errorf("no Chart.yaml found in the repository")
	}

	// Add the repo to the Helm repository config
	repoFile := h.settings.RepositoryConfig
	repositories, err := repo.LoadFile(repoFile)
	if err != nil {
		return fmt.Errorf("failed to load repository file: %w", err)
	}

	entry := &repo.Entry{
		Name: name,
		URL:  tempDir, // Local path acts as the repo
	}

	repositories.Update(entry)
	if err := repositories.WriteFile(repoFile, 0644); err != nil {
		return fmt.Errorf("failed to write repository file: %w", err)
	}

	return nil
}

func CloneGitRepo(name, gitURL, revision, username, password string) (*string, error) {
	tempDir, err := os.MkdirTemp("", "hgr-"+name)
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary directory: %w", err)
	}

	// Clone the Git repository
	if username != "" && password != "" {
		gitURL = fmt.Sprintf("https://%s:%s@%s", username, password, gitURL[8:])
	}

	// Clone the Git repository
	cloneCmd := exec.Command("git", "clone", "--branch", revision, "--single-branch", gitURL, tempDir)

	output, err := cloneCmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to clone Git repository: %w\nOutput: %s", err, string(output))
	}

	return &tempDir, nil
}

func (h *HelmClient) AddGitRepoAsHelmRepo(name, gitURL, revision, subPath, username, password string) error {
	tempDir, err := CloneGitRepo(name, gitURL, revision, username, password)
	if err != nil {
		return fmt.Errorf("failed to create temporary directory: %w", err)
	}

	// clean up
	defer os.RemoveAll(*tempDir)

	// Verify the subpath exists
	chartRootPath := filepath.Join(*tempDir, subPath)
	if _, err := os.Stat(chartRootPath); os.IsNotExist(err) {
		return fmt.Errorf("subpath %q does not exist in the Git repository", subPath)
	}

	// Find the Chart.yaml to verify it's a valid chart repo
	/*chartPath := filepath.Join(chartRootPath, "Chart.yaml")
	if _, err := os.Stat(chartPath); os.IsNotExist(err) {
		return fmt.Errorf("no Chart.yaml found in the chart path")
	}*/

	// Add the subpath directory as a Helm repository
	repoFile := h.settings.RepositoryConfig
	repositories, err := repo.LoadFile(repoFile)
	if err != nil {
		return fmt.Errorf("failed to load repositories file: %w", err)
	}

	entry := &repo.Entry{
		Name: name,
		URL:  chartRootPath,
	}

	repositories.Update(entry)
	if err := repositories.WriteFile(repoFile, 0644); err != nil {
		return fmt.Errorf("failed to write repositories file: %w", err)
	}

	return nil
}

func (h *HelmClient) CreateApplication(repoName, chartPath, releaseName string, namespace string, values map[string]interface{}) (*release.Release, error) {
	client := action.NewInstall(h.config)
	client.ReleaseName = releaseName
	client.Namespace = namespace

	// Load the chart from the specific path
	fullChartPath := filepath.Join(h.settings.RepositoryConfig, repoName, chartPath)
	chart, err := loader.Load(fullChartPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load chart from path %q: %w", fullChartPath, err)
	}

	// Install or upgrade the chart
	return client.Run(chart, values)
}

func (h *HelmClient) SyncApplication(repoName, chartPath, releaseName string, namespace string, values map[string]interface{}) (*release.Release, error) {
	client := action.NewUpgrade(h.config)
	client.Namespace = namespace

	// Pull the latest changes from the Git repository
	tempDir := filepath.Join(h.settings.RepositoryConfig, repoName)
	cmd := exec.Command("git", "-C", tempDir, "pull")
	if output, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("failed to sync Git repository: %w\nOutput: %s", err, string(output))
	}

	// Load the chart
	fullChartPath := filepath.Join(tempDir, chartPath)
	chart, err := loader.Load(fullChartPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load chart from path %q: %w", fullChartPath, err)
	}

	// Upgrade the release with the latest chart and values
	return client.Run(releaseName, chart, values)
}

func debugLog(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
