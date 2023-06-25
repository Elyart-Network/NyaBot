package runtime

import (
	"github.com/Elyart-Network/NyaBot/internal/javascript/gocqhttp"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	log "github.com/sirupsen/logrus"
	"os"
)

func JSVM(path string, data CallbackData) {
	registry := new(require.Registry)
	vm := goja.New()
	file, err := os.ReadFile(path)
	_ = registry.Enable(vm)

	// register modules to the table
	gocqhttp.Action(vm)

	// load and run the script
	_, err = vm.RunString(string(file))
	if jserr, ok := err.(*goja.Exception); ok {
		log.Error("[JSVM] Error running javascript script: ", jserr)
	}

	// get metadata from the javascript script
	metadata := vm.Get("metadata").Export().(map[string]any)

	// send callback data to the javascript script
	switch metadata["type"].(string) {
	case "gocqhttp":
		GoCqAdapter(vm, data)
	}
}
