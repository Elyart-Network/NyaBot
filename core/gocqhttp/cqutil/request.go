package cqutil

import (
	"bytes"
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqws"
	"github.com/Elyart-Network/NyaBot/extend/config"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func GetRequest(Endpoint string, RespStruct interface{}) (err error) {
	// Delay
	time.Sleep(time.Duration(rand.Intn(config.Get("gocqhttp.delay").(int))) * time.Millisecond)
	// Websocket Forward
	CqHttpHost := config.Get("gocqhttp.host_url").(string)
	cqws.WsRequestForward(CqHttpHost, Endpoint, nil, &RespStruct)
	// Websocket Reverse
	cqws.WsRequestReverse(Endpoint, nil, &RespStruct)
	// HTTP Reverse
	go func() {
		Request, err := http.Get(CqHttpHost + Endpoint)
		if err != nil {
			return
		}
		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				return
			}
		}(Request.Body)
		Context, err := io.ReadAll(Request.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(Context, &RespStruct)
		if err != nil {
			return
		}
	}()
	return
}

func PostRequest(Endpoint string, Params interface{}, RespStruct interface{}) (err error) {
	// Delay
	time.Sleep(time.Duration(rand.Intn(config.Get("gocqhttp.delay").(int))) * time.Millisecond)
	// Websocket Forward
	CqHttpHost := config.Get("gocqhttp.host_url").(string)
	cqws.WsRequestForward(CqHttpHost, Endpoint, Params, &RespStruct)
	// Websocket Reverse
	cqws.WsRequestReverse(Endpoint, Params, &RespStruct)
	// HTTP Reverse
	go func() {
		byteSlice, err := json.Marshal(Params)
		if err != nil {
			return
		}
		Request, err := http.Post(CqHttpHost+Endpoint, "application/json", bytes.NewBuffer(byteSlice))
		if err != nil {
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(Request.Body)
		Context, err := io.ReadAll(Request.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(Context, &RespStruct)
		if err != nil {
			return
		}
	}()
	return
}
