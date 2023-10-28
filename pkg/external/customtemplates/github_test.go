package customtemplates

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/khulnasoft-lab/gologger"
	"github.com/khulnasoft-lab/vulmap/v3/pkg/catalog/config"
	"github.com/khulnasoft-lab/vulmap/v3/pkg/testutils"
	"github.com/stretchr/testify/require"
)

func TestDownloadCustomTemplatesFromGitHub(t *testing.T) {
	gologger.DefaultLogger.SetWriter(&testutils.NoopWriter{})

	templatesDirectory, err := os.MkdirTemp("", "template-custom-*")
	require.Nil(t, err, "could not create temp directory")
	defer os.RemoveAll(templatesDirectory)

	config.DefaultConfig.SetTemplatesDir(templatesDirectory)

	options := testutils.DefaultOptions
	options.GitHubTemplateRepo = []string{"khulnasoft-lab/vulmap-templates", "ehsandeep/vulmap-templates"}
	options.GitHubToken = os.Getenv("GITHUB_TOKEN")

	ctm, err := NewCustomTemplatesManager(options)
	require.Nil(t, err, "could not create custom templates manager")

	ctm.Download(context.Background())

	require.DirExists(t, filepath.Join(templatesDirectory, "github", "vulmap-templates-khulnasoft-lab"), "cloned directory does not exists")
	require.DirExists(t, filepath.Join(templatesDirectory, "github", "vulmap-templates-ehsandeep"), "cloned directory does not exists")
}
