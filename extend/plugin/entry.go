package plugin

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqcall"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqws"
	"github.com/gin-gonic/gin"
	"log"
)

func CqEntry(ctx *gin.Context) {
	callback, err := cqcall.CallbackEncode(ctx, false)
	if err != nil {
		log.Println("callback encode error: ", err)
		return
	}
	CqCallBack(callback)
}
func CqWebSocket(ctx *gin.Context) {
	go cqws.WebSocket(ctx, CqCallBack)
}
func DiscordEntry(ctx *gin.Context)  {}
func TelegramEntry(ctx *gin.Context) {}
func SlackEntry(ctx *gin.Context)    {}
