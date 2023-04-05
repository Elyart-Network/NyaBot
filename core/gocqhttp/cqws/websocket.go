package cqws

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqcall"
	"github.com/Elyart-Network/NyaBot/extend/config"
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

func wsHandler(ws *websocket.Conn, callback CqCallback) {
	for {
		// Read Message
		_, wsContext, err := ws.ReadMessage()
		if err != nil {
			log.Println("[WebSocket] ws read error: ", err)
			break
		}
		// Encode Message to wsActionData.
		resp, isResp, err := wsResponseEncode(wsContext)
		if err != nil {
			log.Println("[WebSocket] ws response encode error: ", err)
			break
		}
		// Send callback to plugin functions.
		if !isResp {
			data, err := cqcall.CallbackEncode(wsContext, true)
			if err != nil {
				log.Println("[WebSocket] callback encode error: ", err)
				break
			}
			callback(data)
		}
		// Send request.
		go func() {
			for data := range requestChan {
				err := ws.WriteJSON(data)
				if err != nil {
					log.Println("[WebSocket] ws write json error: ", err)
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

func WebSocketClient(callback CqCallback) {
	dialer := websocket.DefaultDialer
	wsHost := config.Get("gocqhttp.host_url").(string)
	ws, _, err := dialer.Dial(wsHost, nil)
	if err != nil {
		log.Println("[WebSocket] ws dial error: ", err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Println("[WebSocket] ws close error: ", err)
			return
		}
	}(ws)
	wsHandler(ws, callback)
}

func WebSocketServer(ctx *gin.Context, callback CqCallback) {
	ws, err := up.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("[WebSocket] ws upgrade error: ", err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Println("[WebSocket] ws close error: ", err)
			return
		}
	}(ws)
	wsHandler(ws, callback)
}
