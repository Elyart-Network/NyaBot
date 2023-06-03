package logger

import (
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	FileLogger := config.Get("server.file_logger").(bool)
	DebugMode := config.Get("server.debug_mode").(bool)

	if FileLogger {
		gin.DisableConsoleColor()
		file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Panicln("Error opening log file: ", err)
		}
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(file)
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		log.SetLevel(log.WarnLevel)
	}

	if DebugMode {
		gin.SetMode(gin.DebugMode)
		log.SetLevel(log.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func Gin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		log.WithFields(log.Fields{
			"METHOD":    reqMethod,
			"URI":       reqUri,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIP,
		}).Info("HTTP REQUEST")

		ctx.Next()
	}
}
