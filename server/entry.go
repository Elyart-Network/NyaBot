package server

import (
	"github.com/Elyart-Network/NyaBot/extend/config"
	"github.com/Elyart-Network/NyaBot/extend/plugin"
	"github.com/gin-gonic/gin"
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
		if config.Get("gocqhttp.enable_ws").(bool) && !config.Get("gocqhttp.ws_forward").(bool) {
			server.GET("/api/gocqhttp", plugin.CqWebSocket)
		}
		server.POST("/api/gocqhttp", plugin.CqEntry)
	}
}
