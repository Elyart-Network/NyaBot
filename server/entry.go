package server

import (
	config2 "github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/pkg/plugin"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func Entry(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status": "Not Found",
		})
	})
	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	if config2.EncodeMagic(os.Getenv("MAGIC")) == "000000000000000000000000c764bcb2dc755ba7a60dc20dec2dc7f18f68c4b56d84950eca9f6e7516d6ee0d2571b0c278a9a861bfa3235c" {
		server.POST("/webhook", plugin.WhEntry)
	}
	if config2.Get("gocqhttp.enable").(bool) {
		if config2.Get("gocqhttp.enable_ws").(bool) {
			switch strings.HasPrefix(config2.Get("gocqhttp.host_url").(string), "ws") {
			case true:
				plugin.CqWebSocketForward()
			case false:
				server.GET("/api/gocqhttp", plugin.CqWebSocketReverse)
			}
		}
		server.POST("/api/gocqhttp", plugin.CqEntry)
	}
}
