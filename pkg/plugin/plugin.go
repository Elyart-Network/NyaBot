package plugin

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/webhook"
	"log"
)

var plugins = make(map[string]interface{})

func WhRegister(plugin WhPlugin) {
	plugins[plugin.(CommonInfo).Info().Name] = plugin
	log.Println("[Nya-Plugin] Plugin", plugin.(CommonInfo).Info().Name, "registered.")
}

func CqRegister(plugin CqPlugin) {
	plugins[plugin.(CommonInfo).Info().Name] = plugin
	log.Println("[Nya-Plugin] Plugin", plugin.(CommonInfo).Info().Name, "registered.")
}

func WhCallBack(callback webhook.Data) {
	// Send callback to plugin functions.
	for _, value := range plugins {
		value := value
		go func() {
			if value.(CommonInfo).Info().Type == "Webhook" {
				value.(WhPlugin).Receive(callback)
			}
		}()
	}
}

func CqCallBack(callback callback.Full) {
	// Send callback to plugin functions.
	for _, value := range plugins {
		value := value
		go func() {
			if value.(CommonInfo).Info().Type == "GoCqHttp" {
				switch callback.PostType {
				case "message":
					value.(CqPlugin).Message(callback)
				case "request":
					value.(CqPlugin).Request(callback)
				case "notice":
					value.(CqPlugin).Notice(callback)
				case "meta_event":
					value.(CqPlugin).MetaEvent(callback)
				}
			}
		}()
	}
}
