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
		log.Error("[Plugin] Webhook callback decode error: ", err)
		return
	}
	log.Debug("[Entry] Received WebHook Callback. @Data:", data)
	WhCallBack(data)
}

func CqEntry(ctx *gin.Context) {
	data, err := callback.Encode(ctx, false)
	if err != nil {
		log.Error("[Plugin] GoCqHttp callback decode error: ", err)
		return
	}
	log.Debug("[Entry] Received CoolQ Callback. @Data:", data)
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
