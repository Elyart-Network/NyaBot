package plugin

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqcall"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqws"
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
func CqWebSocketForward() {
	go cqws.WebSocketClient(CqCallBack)
}
func CqWebSocketReverse(ctx *gin.Context) {
	go cqws.WebSocketServer(ctx, CqCallBack)
}
func DiscordEntry(ctx *gin.Context)  {}
func TelegramEntry(ctx *gin.Context) {}
func SlackEntry(ctx *gin.Context)    {}
