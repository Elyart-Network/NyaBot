package server

import (
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var g errgroup.Group

func Start() {
	FileLogger := config.Get("server.file_logger").(bool)
	ServerPort := config.Get("server.listen_port").(string)
	RpcPort := config.Get("server.rpc_port").(string)
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

	gs := &http.Server{
		Addr:         ":" + ServerPort,
		Handler:      GinServer(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	rpc := &http.Server{
		Addr:         ":" + RpcPort,
		Handler:      RPCServer(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return gs.ListenAndServe()
	})

	g.Go(func() error {
		if RpcPort != "" {
			return rpc.ListenAndServe()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Panicln(err)
	}
}
