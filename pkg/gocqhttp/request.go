package gocqhttp

import (
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/internal/config"
	"github.com/Elyart-Network/NyaBot/internal/utils"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/websocket"
	"math/rand"
	"time"
)

func GetRequest(Endpoint string, RespStruct interface{}) (err error) {
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
	return
}

func PostRequest(Endpoint string, Params interface{}, RespStruct interface{}) (err error) {
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
	return
}
