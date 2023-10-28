package installer

import (
	"testing"

	"github.com/khulnasoft-lab/vulmap/v3/pkg/catalog/config"
	"github.com/khulnasoft-lab/utils/generic"
	"github.com/stretchr/testify/require"
)

func TestVersionCheck(t *testing.T) {
	err := VulmapVersionCheck()
	require.Nil(t, err)
	cfg := config.DefaultConfig
	if generic.EqualsAny("", cfg.LatestVulmapIgnoreHash, cfg.LatestVulmapVersion, cfg.LatestVulmapTemplatesVersion) {
		// all above values cannot be empty
		t.Errorf("something went wrong got empty response vulmap-version=%v templates-version=%v ignore-hash=%v", cfg.LatestVulmapVersion, cfg.LatestVulmapTemplatesVersion, cfg.LatestVulmapIgnoreHash)
	}
}
