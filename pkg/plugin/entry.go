package plugin

import (
	"github.com/Elyart-Network/NyaBot/internal/logger"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/websocket"
	"github.com/Elyart-Network/NyaBot/pkg/webhook"
	"github.com/gin-gonic/gin"
)

func WhEntry(ctx *gin.Context) {
	data := webhook.Data{}
	err := ctx.BindJSON(&data)
	if err != nil {
		logger.Warningf("Plugin", "Webhook callback decode error", err)
		return
	}
	WhCallBack(data)
}

func CqEntry(ctx *gin.Context) {
	data, err := callback.Encode(ctx, false)
	if err != nil {
		logger.Warningf("Plugin", "GoCqHttp callback decode error", err)
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
