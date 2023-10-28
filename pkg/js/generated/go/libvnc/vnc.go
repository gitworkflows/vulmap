package vnc

import (
	lib_vnc "github.com/khulnasoft-lab/vulmap/v3/pkg/js/libs/vnc"

	"github.com/dop251/goja"
	"github.com/khulnasoft-lab/vulmap/v3/pkg/js/gojs"
)

var (
	module = gojs.NewGojaModule("vulmap/vnc")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions

			// Var and consts

			// Types (value type)
			"IsVNCResponse": func() lib_vnc.IsVNCResponse { return lib_vnc.IsVNCResponse{} },
			"VNCClient":     func() lib_vnc.VNCClient { return lib_vnc.VNCClient{} },

			// Types (pointer type)
			"NewIsVNCResponse": func() *lib_vnc.IsVNCResponse { return &lib_vnc.IsVNCResponse{} },
			"NewVNCClient":     func() *lib_vnc.VNCClient { return &lib_vnc.VNCClient{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
