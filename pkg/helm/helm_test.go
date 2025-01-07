package helm_test

import (
	"github.com/krack8/lighthouse/pkg/helm"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

	t.Run("install chart", func(t *testing.T) {
		release, err := helmClient.InstallChart(
			"nginx-stable",  // Repository name
			"nginx-ingress", // Chart name
			"2.0.0",         // Chart version
			"kni",           // Release name
			"default",       // Namespace
			nil,             // Values map
		)
		assert.NoError(t, err)
		assert.NotNil(t, release)
	})

	t.Run("list releases", func(t *testing.T) {
		releases, err := helmClient.ListReleases(false)
		assert.NoError(t, err)
		assert.NotEmpty(t, releases)
	})

	time.Sleep(5 * time.Second)

	t.Run("uninstall chart", func(t *testing.T) {
		resp, err := helmClient.UninstallChart("kni")
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("list releases", func(t *testing.T) {
		releases, err := helmClient.ListReleases(false)
		assert.NoError(t, err)
		assert.Empty(t, releases)
	})

	t.Run("remove repo", func(t *testing.T) {
		err = helmClient.RemoveRepo("nginx-stable")
		assert.NoError(t, err)
	})

}
