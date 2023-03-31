package plugin

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqcall"
	"github.com/gin-gonic/gin"
	"log"
)

func GoCqLoader(ctx *gin.Context, function GoCqInvoker) {
	callback := cqcall.Callback(ctx)
	switch callback.PostType {
	case "message":
		function(callback)
	case "request":
		function(callback)
	case "notice":
		function(callback)
	case "meta_event":
		function(callback)
	default:
		log.Println("Unknown PostType: ", callback.PostType)
	}
	return
}
