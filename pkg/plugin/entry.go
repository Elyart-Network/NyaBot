package plugin

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/websocket"
	"github.com/Elyart-Network/NyaBot/pkg/webhook"
	"github.com/gin-gonic/gin"
	"log"
)

func WhEntry(ctx *gin.Context) {
	data := webhook.Data{}
	err := ctx.BindJSON(&data)
	if err != nil {
		log.Println("callback encode error: ", err)
		return
	}
	WhCallBack(data)
}

func CqEntry(ctx *gin.Context) {
	data, err := callback.Encode(ctx, false)
	if err != nil {
		log.Println("callback encode error: ", err)
		return
	}
	CqCallBack(data)
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
