package websocket

import (
	"context"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type CqCallback func(ctx context.Context, callback callback.Full)

var up = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(ws *websocket.Conn, call CqCallback) {
	for {
		// Read Message
		_, wsContext, err := ws.ReadMessage()
		if err != nil {
			log.Error("[WebSocket] ws read error: ", err)
			break
		}
		// Encode Message to wsActionData.
		resp, isResp, err := wsResponseEncode(wsContext)
		if err != nil {
			log.Error("[WebSocket] ws response encode error: ", err)
			break
		}
		// Send callback to plugin functions.
		if !isResp {
			data, err := callback.Encode(wsContext, true)
			if err != nil {
				log.Error("[WebSocket] callback encode error: ", err)
				break
			}
			ctx := context.Background()
			call(ctx, data)
		}
		// Send request.
		go func() {
			for data := range requestChan {
				err := ws.WriteJSON(data)
				if err != nil {
					log.Error("[WebSocket] ws write json error: ", err)
					break
				}
			}
		}()
		// Send response.
		if isResp {
			responseChan <- resp
		}
	}
}

func Client(callback CqCallback) {
	dialer := websocket.DefaultDialer
	wsHost := config.Get("gocqhttp.host_url").(string)
	ws, _, err := dialer.Dial(wsHost, nil)
	if err != nil {
		log.Error("[WebSocket] ws dial error: ", err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Error("[WebSocket] ws close error: ", err)
			return
		}
	}(ws)
	wsHandler(ws, callback)
}

func Server(ctx *gin.Context, callback CqCallback) {
	ws, err := up.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Error("[WebSocket] ws upgrade error: ", err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Error("[WebSocket] ws close error: ", err)
			return
		}
	}(ws)
	wsHandler(ws, callback)
}
