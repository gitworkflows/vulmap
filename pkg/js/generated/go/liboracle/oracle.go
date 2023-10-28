package oracle

import (
	lib_oracle "github.com/khulnasoft-lab/vulmap/v3/pkg/js/libs/oracle"

	"github.com/dop251/goja"
	"github.com/khulnasoft-lab/vulmap/v3/pkg/js/gojs"
)

var (
	module = gojs.NewGojaModule("vulmap/oracle")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions

			// Var and consts

			// Types (value type)
			"IsOracleResponse": func() lib_oracle.IsOracleResponse { return lib_oracle.IsOracleResponse{} },
			"OracleClient":     func() lib_oracle.OracleClient { return lib_oracle.OracleClient{} },

			// Types (pointer type)
			"NewIsOracleResponse": func() *lib_oracle.IsOracleResponse { return &lib_oracle.IsOracleResponse{} },
			"NewOracleClient":     func() *lib_oracle.OracleClient { return &lib_oracle.OracleClient{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
