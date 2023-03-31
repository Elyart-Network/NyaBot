package cqutil

import (
	"bytes"
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/extend/config"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func GetRequest(Endpoint string, RespStruct interface{}) (err error) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	CqHttpHost := config.Get("gocqhttp.http_host").(string)
	Request, err := http.Get(CqHttpHost + Endpoint)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(Request.Body)
	Context, err := io.ReadAll(Request.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(Context, &RespStruct)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func PostRequest(Endpoint string, Params interface{}, RespStruct interface{}) (err error) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	CqHttpHost := config.Get("gocqhttp.http_host").(string)
	byteSlice, err := json.Marshal(Params)
	if err != nil {
		log.Println(err)
		return
	}
	Request, err := http.Post(CqHttpHost+Endpoint, "application/json", bytes.NewBuffer(byteSlice))
	if err != nil {
		log.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(Request.Body)
	Context, err := io.ReadAll(Request.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(Context, &RespStruct)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
