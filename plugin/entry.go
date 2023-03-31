package plugin

import (
	"github.com/Elyart-Network/NyaBot/extend/plugin"
	"github.com/Elyart-Network/NyaBot/plugin/FlarumGroup"
	"github.com/gin-gonic/gin"
)

func GoCqEntry(ctx *gin.Context) {
	go plugin.GoCqLoader(ctx, FlarumGroup.Entry)
}

func DiscordEntry(ctx *gin.Context)  {}
func TelegramEntry(ctx *gin.Context) {}
func SlackEntry(ctx *gin.Context)    {}
