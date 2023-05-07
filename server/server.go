package server

import (
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/logger"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func Start() {
	FileLogger := config.Get("server.file_logger").(bool)
	ServerPort := config.Get("server.listen_port").(string)
	DebugMode := config.Get("server.debug_mode").(bool)

	if FileLogger {
		gin.DisableConsoleColor()
		file, err := os.Create("app.log")
		if err != nil {
			log.Panicln(err)
		}
		gin.DefaultWriter = io.MultiWriter(file)
	}

	if DebugMode {
		gin.SetMode(gin.DebugMode)
		logger.SetDebug(true)
	} else {
		gin.SetMode(gin.ReleaseMode)
		logger.SetDebug(false)
	}

	engine := gin.Default()
	engine.Use(gin.Recovery())
	Entry(engine)

	err := http.ListenAndServe(":"+ServerPort, engine)
	if err != nil {
		log.Panicln(err)
	}
}
