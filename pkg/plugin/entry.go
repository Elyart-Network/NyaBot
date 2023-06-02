package plugin

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/websocket"
	"github.com/Elyart-Network/NyaBot/pkg/webhook"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func WhEntry(ctx *gin.Context) {
	data := webhook.Data{}
	err := ctx.BindJSON(&data)
	if err != nil {
		log.Errorf("[Plugin] Webhook callback decode error: %v", err)
		return
	}
	WhCallBack(data)
}

func CqEntry(ctx *gin.Context) {
	data, err := callback.Encode(ctx, false)
	if err != nil {
		log.Errorf("[Plugin] GoCqHttp callback decode error: %v", err)
		return
	}
	CqCallBack(data)
}
func CqWebSocketForward() {
	go websocket.Client(CqCallBack)
}
func CqWebSocketReverse(ctx *gin.Context) {
	go websocket.Server(ctx.Writer, ctx.Request, CqCallBack)
}
func DiscordEntry(ctx *gin.Context)  {}
func TelegramEntry(ctx *gin.Context) {}
func SlackEntry(ctx *gin.Context)    {}
