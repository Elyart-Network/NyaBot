package plugin

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqcall"
	"github.com/gin-gonic/gin"
)

func GoCqEntry(ctx *gin.Context) {
	cqCallback = cqcall.Callback(ctx)
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
