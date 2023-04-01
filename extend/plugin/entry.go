package plugin

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqcall"
	"github.com/gin-gonic/gin"
)

func GoCqEntry(ctx *gin.Context) {
	cqCallback = cqcall.Callback(ctx)
	cqPlugins := make(map[string]GoCqPlugin)
	// Filter plugins.
	for key, value := range plugins {
		if key.Type == "cq" {
			cqPlugins[key.Name] = value.(GoCqPlugin)
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
