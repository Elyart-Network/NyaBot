package cqutil

import (
	"bytes"
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/internal/config"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqws"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func GetRequest(Endpoint string, RespStruct interface{}) (err error) {
	// Delay
	time.Sleep(time.Duration(rand.Intn(config.Get("gocqhttp.delay").(int))) * time.Millisecond)
	// Websocket Reverse
	cqws.WsSendRequest(Endpoint, nil, &RespStruct)
	// HTTP Reverse
	CqHttpHost := config.Get("gocqhttp.host_url").(string)
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
	// Websocket Reverse
	cqws.WsSendRequest(Endpoint, Params, &RespStruct)
	// HTTP Reverse
	CqHttpHost := config.Get("gocqhttp.host_url").(string)
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
