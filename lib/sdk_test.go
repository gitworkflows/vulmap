package vulmap_test

import (
	"testing"

	vulmap "github.com/khulnasoft-lab/vulmap/v3/lib"
	"github.com/stretchr/testify/require"
)

func TestSimpleVulmap(t *testing.T) {
	ne, err := vulmap.NewVulmapEngine(
		vulmap.WithTemplateFilters(vulmap.TemplateFilters{ProtocolTypes: "dns"}),
		vulmap.EnableStatsWithOpts(vulmap.StatsOptions{JSON: true}),
	)
	require.Nil(t, err)
	ne.LoadTargets([]string{"scanme.sh"}, false) // probe non http/https target is set to false here
	// when callback is nil it vulmap will print JSON output to stdout
	err = ne.ExecuteWithCallback(nil)
	require.Nil(t, err)
	defer ne.Close()
}

func TestSimpleVulmapRemote(t *testing.T) {
	ne, err := vulmap.NewVulmapEngine(
		vulmap.WithTemplatesOrWorkflows(
			vulmap.TemplateSources{
				RemoteTemplates: []string{"https://templates.vulmap.sh/public/nameserver-fingerprint.yaml"},
			},
		),
	)
	require.Nil(t, err)
	ne.LoadTargets([]string{"scanme.sh"}, false) // probe non http/https target is set to false here
	err = ne.LoadAllTemplates()
	require.Nil(t, err, "could not load templates")
	// when callback is nil it vulmap will print JSON output to stdout
	err = ne.ExecuteWithCallback(nil)
	require.Nil(t, err)
	defer ne.Close()
}

func TestThreadSafeVulmap(t *testing.T) {
	// create vulmap engine with options
	ne, err := vulmap.NewThreadSafeVulmapEngine()
	require.Nil(t, err)

	// scan 1 = run dns templates on scanme.sh
	t.Run("scanme.sh", func(t *testing.T) {
		err = ne.ExecuteVulmapWithOpts([]string{"scanme.sh"}, vulmap.WithTemplateFilters(vulmap.TemplateFilters{ProtocolTypes: "dns"}))
		require.Nil(t, err)
	})

	// scan 2 = run dns templates on honey.scanme.sh
	t.Run("honey.scanme.sh", func(t *testing.T) {
		err = ne.ExecuteVulmapWithOpts([]string{"honey.scanme.sh"}, vulmap.WithTemplateFilters(vulmap.TemplateFilters{ProtocolTypes: "dns"}))
		require.Nil(t, err)
	})

	// wait for all scans to finish
	defer ne.Close()
}
