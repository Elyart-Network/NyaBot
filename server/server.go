package server

import (
	"github.com/Elyart-Network/NyaBot/config"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

var g errgroup.Group

func Start() {
	ServerPort := config.Get("server.listen_port").(string)
	RpcPort := config.Get("server.rpc_port").(string)

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
