package helm_test

import (
	"github.com/krack8/lighthouse/pkg/helm"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHelm(t *testing.T) {
	var helmClient *helm.HelmClient
	var err error

	t.Run("initialize helm client", func(t *testing.T) {
		helmClient, err = helm.NewHelmClient("default")
		assert.NoError(t, err)
	})

	t.Run("initialize helm repo", func(t *testing.T) {
		err = helmClient.InitializeHelmRepo()
		assert.NoError(t, err)
	})

	t.Run("add repo", func(t *testing.T) {
		err = helmClient.AddRepo("nginx-stable", "https://helm.nginx.com/stable")
		assert.NoError(t, err)
	})

	t.Run("list repo", func(t *testing.T) {
		repos, err := helmClient.ListRepos()
		assert.NoError(t, err)
		assert.NotEmpty(t, repos)
	})

	t.Run("list charts", func(t *testing.T) {
		charts, err := helmClient.ListChartsInRepo("nginx-stable")
		assert.NoError(t, err)
		assert.NotEmpty(t, charts)
	})

	t.Run("list charts", func(t *testing.T) {
		charts, err := helmClient.SearchChart("nginx-stable", "nginx-ingress")
		assert.NoError(t, err)
		assert.NotEmpty(t, charts)
	})

	t.Run("get chart values", func(t *testing.T) {
		chartValues, err := helmClient.GetChartValues("nginx-stable", "nginx-ingress", "2.0.0")
		assert.NoError(t, err)
		assert.NotEmpty(t, chartValues)
	})

	t.Run("install chart", func(t *testing.T) {
		release, err := helmClient.InstallOrUpgradeChart(
			"nginx-stable",  // Repository name
			"nginx-ingress", // Chart name
			"2.0.0",         // Chart version
			"kni",           // Release name
			"default",       // Namespace
			map[string]interface{}{
				"controller": map[string]interface{}{
					"replicaCount": 2,
				},
			}, // Values map
		)
		assert.NoError(t, err)
		assert.NotNil(t, release)
	})

	t.Run("get applied values of current revision", func(t *testing.T) {
		chartValues, err := helmClient.GetCurrentAppliedValues("kni")
		assert.NoError(t, err)
		assert.NotEmpty(t, chartValues)
	})

	t.Run("get applied values of a specific revision", func(t *testing.T) {
		chartValues, err := helmClient.GetAppliedValues("kni", 1)
		assert.NoError(t, err)
		assert.NotEmpty(t, chartValues)
	})

	t.Run("list releases - 1", func(t *testing.T) {
		releases, err := helmClient.ListReleases(false)
		assert.NoError(t, err)
		assert.NotEmpty(t, releases)
	})

	t.Run("get release details", func(t *testing.T) {
		release, err := helmClient.GetReleaseDetails("kni")
		assert.NoError(t, err)
		assert.NotEmpty(t, release)
	})

	t.Run("list revisions", func(t *testing.T) {
		revisions, err := helmClient.ListRevisions("kni")
		assert.NoError(t, err)
		assert.NotEmpty(t, revisions)
	})

	t.Run("get current revision details", func(t *testing.T) {
		revision, err := helmClient.GetCurrentRevisionDetails("kni")
		assert.NoError(t, err)
		assert.NotEmpty(t, revision)
	})

	t.Run("get a specific revision details", func(t *testing.T) {
		revision, err := helmClient.GetRevisionDetails("kni", 1)
		assert.NoError(t, err)
		assert.NotEmpty(t, revision)
	})

	t.Run("uninstall chart", func(t *testing.T) {
		resp, err := helmClient.UninstallChart("kni")
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("list releases - 2", func(t *testing.T) {
		releases, err := helmClient.ListReleases(false)
		assert.NoError(t, err)
		assert.Empty(t, releases)
	})

	t.Run("remove repo", func(t *testing.T) {
		err = helmClient.RemoveRepo("nginx-stable")
		assert.NoError(t, err)
	})

	t.Run("git clone", func(t *testing.T) {
		repoClonePath, err := helm.CloneGitRepo("lighthouse", "https://github.com/krack8/lighthouse.git", "main", "", "")
		assert.NoError(t, err)
		assert.NotNil(t, repoClonePath)

		_, err = os.Stat(*repoClonePath)
		assert.NoError(t, err)

		if repoClonePath != nil {
			err = os.RemoveAll(*repoClonePath)
			assert.NoError(t, err)
		}
	})

	t.Run("add git repo as helm repo", func(t *testing.T) {
		err := helmClient.AddGitRepoAsHelmRepo("lighthouse-nginx", "https://github.com/krack8/lighthouse.git", "helm-functions", "helm/Charts/nginx", "", "")
		assert.NoError(t, err)
	})

	t.Run("create application from git helm repo", func(t *testing.T) {
		err := helmClient.AddGitRepoAsHelmRepo("lighthouse-nginx", "https://github.com/krack8/lighthouse.git", "helm-functions", "helm", "", "")
		assert.NoError(t, err)

		release, err := helmClient.CreateApplication("lighthouse-nginx", "nginx", "lighthouse-nginx", "default", map[string]interface{}{"replicaCount": 3})
		assert.NoError(t, err)
		assert.NotNil(t, release)
	})

	/*t.Run("uninstall application", func(t *testing.T) {
		resp, err := helmClient.UninstallChart("lighthouse-nginx")
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})*/
}
