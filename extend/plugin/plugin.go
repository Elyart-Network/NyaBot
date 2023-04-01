package plugin

import (
	"log"
)

var plugins = make(map[string]interface{})

func CqRegister(plugin CqPlugin) {
	plugins[plugin.(CommonInfo).Info().Name] = plugin
	log.Println("[Nya-Plugin] Plugin", plugin.(CommonInfo).Info().Name, "registered.")
}
