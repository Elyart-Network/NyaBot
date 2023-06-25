package logger

import (
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	FileLogger := config.Get("server.file_logger").(bool)
	LogLevel := config.Get("server.log_level").(string)

	// Set log format
	log.SetFormatter(&log.JSONFormatter{})
	gin.SetMode(gin.ReleaseMode)
	switch LogLevel {
	case "debug":
		gin.SetMode(gin.DebugMode)
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		gin.SetMode(gin.ReleaseMode)
		log.SetLevel(log.InfoLevel)
	}

	// Set log output
	if FileLogger {
		gin.DisableConsoleColor()
		file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Panicln("Error opening log file: ", err)
		}
		log.SetOutput(file)
	} else {
		log.SetOutput(os.Stdout)
	}
}
