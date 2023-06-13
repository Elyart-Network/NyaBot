package gocqhttp

import (
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/websocket"
	"github.com/Elyart-Network/NyaBot/utils"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func GetRequest(Endpoint string, RespStruct any) (err error) {
	// Delay
	time.Sleep(time.Duration(rand.Intn(config.Get("gocqhttp.delay").(int))) * time.Millisecond)
	// Websocket Reverse
	if websocket.WsSendRequest(Endpoint, nil, &RespStruct) {
		return
	}
	// HTTP Reverse
	CqHttpHost := config.Get("gocqhttp.host_url").(string)
	request, err := utils.GetRequest(CqHttpHost+Endpoint, "")
	if err != nil {
		return err
	}
	err = json.Unmarshal(request, &RespStruct)
	if err != nil {
		return err
	}
	log.Debug("[GoCqHttp] HTTP GET REQUEST sent! @Endpoint:", Endpoint, " @Response:", RespStruct)
	return
}

func PostRequest(Endpoint string, Params any, RespStruct any) (err error) {
	// Delay
	time.Sleep(time.Duration(rand.Intn(config.Get("gocqhttp.delay").(int))) * time.Millisecond)
	// Websocket Reverse
	if websocket.WsSendRequest(Endpoint, Params, &RespStruct) {
		return
	}
	// HTTP Reverse
	CqHttpHost := config.Get("gocqhttp.host_url").(string)
	request, err := utils.PostRequest(CqHttpHost+Endpoint, Params)
	if err != nil {
		return err
	}
	err = json.Unmarshal(request, &RespStruct)
	if err != nil {
		return err
	}
	log.Debug("[GoCqHttp] HTTP POST REQUEST sent! @Endpoint:", Endpoint, " @Params:", Params, " @Response:", RespStruct)
	return
}
