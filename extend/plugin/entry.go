package plugin

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqcall"
	"github.com/gin-gonic/gin"
)

func CqEntry(ctx *gin.Context) {
	// Get callback from request.
	cqCallback = cqcall.Callback(ctx)
	cqPlugins := make(map[string]CqPlugin)
	// Filter plugins.
	for _, value := range plugins {
		if value.(CommonInfo).Info().Type == "GoCqHttp" {
			cqPlugins[value.(CommonInfo).Info().Name] = value.(CqPlugin)
		}
	}
	// Send callback to plugin functions.
	for _, value := range cqPlugins {
		switch cqCallback.PostType {
		case "message":
			value.Message(cqCallback)
		case "request":
			value.Request(cqCallback)
		case "notice":
			value.Notice(cqCallback)
		case "meta_event":
			value.MetaEvent(cqCallback)
		}
	}
}
func DiscordEntry(ctx *gin.Context)  {}
func TelegramEntry(ctx *gin.Context) {}
func SlackEntry(ctx *gin.Context)    {}
