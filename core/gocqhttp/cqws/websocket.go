package cqws

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqcall"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type CqCallback func(callback cqcall.CallbackFull)

var up = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocket(ctx *gin.Context, callback CqCallback) {
	ws, err := up.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(ws)
	for {
		// Read Message
		_, wsContext, err := ws.ReadMessage()
		if err != nil {
			log.Println("ws read error: ", err)
			break
		}
		// Encode Message to wsActionData.
		resp, isResp, err := wsResponseEncode(wsContext)
		if err != nil {
			log.Println("ws response encode error: ", err)
			break
		}
		// Send callback to plugin functions.
		if !isResp {
			data, err := cqcall.CallbackEncode(wsContext, true)
			if err != nil {
				log.Println("callback encode error: ", err)
				break
			}
			callback(data)
		}
		// Send request.
		go func() {
			for data := range requestChan {
				err := ws.WriteJSON(data)
				if err != nil {
					log.Println("ws write json error: ", err)
					break
				}
			}
		}()
		// Send response.
		if isResp {
			wsResponse(resp)
		}
	}
}
