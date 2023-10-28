package runner

import (
	"fmt"

	"github.com/khulnasoft-lab/gologger"
	"github.com/khulnasoft-lab/vulmap/v3/pkg/catalog/config"
	updateutils "github.com/khulnasoft-lab/utils/update"
)

var banner = fmt.Sprintf(`
                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/   %s
`, config.Version)

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tkhulnasoft-lab.io\n\n")
}

// VulmapToolUpdateCallback updates vulmap binary/tool to latest version
func VulmapToolUpdateCallback() {
	showBanner()
	updateutils.GetUpdateToolCallback("vulmap", config.Version)()
}
