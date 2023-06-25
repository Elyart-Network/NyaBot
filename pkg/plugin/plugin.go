package plugin

import (
	"context"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/webhook"
	log "github.com/sirupsen/logrus"
)

var plugins = make(map[string]any)

func WhRegister(plugin WhPlugin) {
	plugins[plugin.(CommonInfo).Info().Name] = plugin
	log.Debug("[Nya-Plugin] Plugin ", plugin.(CommonInfo).Info().Name, " registered.")
}

func CqRegister(plugin CqPlugin) {
	plugins[plugin.(CommonInfo).Info().Name] = plugin
	log.Debug("[Nya-Plugin] Plugin ", plugin.(CommonInfo).Info().Name, " registered.")
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

func CqCallBack(ctx context.Context, callback callback.Full) {
	// Send callback to plugin functions.
	for _, value := range plugins {
		value := value
		go func() {
			if value.(CommonInfo).Info().Type == "GoCqHttp" {
				switch callback.PostType {
				case "message":
					log.Debug("[Plugin]", "(GoCqHttp)", "{"+value.(CommonInfo).Info().Name+"}", " Message Event Received.")
					value.(CqPlugin).Message(ctx, callback)
				case "request":
					log.Debug("[Plugin]", "(GoCqHttp)", "{"+value.(CommonInfo).Info().Name+"}", " Request Event Received.")
					value.(CqPlugin).Request(ctx, callback)
				case "notice":
					log.Debug("[Plugin]", "(GoCqHttp)", "{"+value.(CommonInfo).Info().Name+"}", " Notice Event Received.")
					value.(CqPlugin).Notice(ctx, callback)
				case "meta_event":
					log.Debug("[Plugin]", "(GoCqHttp)", "{"+value.(CommonInfo).Info().Name+"}", " Meta Event Received.")
					value.(CqPlugin).MetaEvent(ctx, callback)
				}
			}
		}()
	}
}
