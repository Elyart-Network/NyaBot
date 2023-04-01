package plugin

import (
	"log"
)

var plugins = make(map[InfoStruct]interface{})

func CqPlugRegister(plugin GoCqPlugin) {
	plugins[plugin.Info()] = plugin
	log.Println("[Nya-GoCq] Plugin", plugin.Info().Name, "registered.")
}
