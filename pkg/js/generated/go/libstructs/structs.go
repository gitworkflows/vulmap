package structs

import (
	lib_structs "github.com/khulnasoft-lab/vulmap/v3/pkg/js/libs/structs"

	"github.com/dop251/goja"
	"github.com/khulnasoft-lab/vulmap/v3/pkg/js/gojs"
)

var (
	module = gojs.NewGojaModule("vulmap/structs")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"Pack":            lib_structs.Pack,
			"StructsCalcSize": lib_structs.StructsCalcSize,
			"Unpack":          lib_structs.Unpack,

			// Var and consts

			// Types (value type)

			// Types (pointer type)
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
