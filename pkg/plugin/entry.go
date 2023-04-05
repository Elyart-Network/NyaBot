package plugin

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/websocket"
	"github.com/gin-gonic/gin"
	"log"
)

func CqEntry(ctx *gin.Context) {
	callback, err := callback.Encode(ctx, false)
	if err != nil {
		log.Println("callback encode error: ", err)
		return
	}
	CqCallBack(callback)
}
func CqWebSocketForward() {
	go websocket.Client(CqCallBack)
}
func CqWebSocketReverse(ctx *gin.Context) {
	go websocket.Server(ctx, CqCallBack)
}
func DiscordEntry(ctx *gin.Context)  {}
func TelegramEntry(ctx *gin.Context) {}
func SlackEntry(ctx *gin.Context)    {}
