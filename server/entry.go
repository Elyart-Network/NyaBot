package server

import (
	"github.com/Elyart-Network/NyaBot/internal/config"
	"github.com/Elyart-Network/NyaBot/pkg/plugin"
	"github.com/gin-gonic/gin"
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
	if config.Get("gocqhttp.enable").(bool) {
		if config.Get("gocqhttp.enable_ws").(bool) {
			switch strings.HasPrefix(config.Get("gocqhttp.host_url").(string), "ws") {
			case true:
				plugin.CqWebSocketForward()
			case false:
				server.GET("/api/gocqhttp", plugin.CqWebSocketReverse)
			}
		}
		server.POST("/api/gocqhttp", plugin.CqEntry)
	}
	if config.Get("mirai.enable").(bool) {

	}
}
