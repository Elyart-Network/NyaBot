package common

import (
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/utils"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type HttpReqFunc struct{}

func (h *HttpReqFunc) Get(url string, params string) string {
	resp, err := utils.GetRequest(url, params)
	if err != nil {
		log.Warningf("[Lua] Error while sending GET request: %v", err)
		return ""
	}
	return string(resp)
}

func (h *HttpReqFunc) GetJson(url string, params string) map[string]any {
	var data map[string]any
	resp, err := utils.GetRequest(url, params)
	if err != nil {
		log.Warningf("[Lua] Error while sending GET request: %v", err)
		return nil
	}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Warningf("[Lua] Error while parsing JSON: %v", err)
		return nil
	}
	return data
}

func (h *HttpReqFunc) Post(url string, params any) string {
	resp, err := utils.PostRequest(url, params)
	if err != nil {
		log.Warningf("[Lua] Error while sending POST request: %v", err)
		return ""
	}
	return string(resp)
}

func (h *HttpReqFunc) PostJson(url string, params any) map[string]any {
	var data map[string]any
	resp, err := utils.PostRequest(url, params)
	if err != nil {
		log.Warningf("[Lua] Error while sending POST request: %v", err)
		return nil
	}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Warningf("[Lua] Error while parsing JSON: %v", err)
		return nil
	}
	return data
}

func HttpReq(L *lua.LState) {
	var HttpReqFunc = &HttpReqFunc{}
	L.SetGlobal("HttpReq", luar.New(L, HttpReqFunc))
}
