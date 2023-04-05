package cqws

import (
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/extend/config"
	"log"
	"strconv"
	"strings"
	"time"
)

type wsRequestData struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}

var requestChan = make(chan wsRequestData)

func WsSendRequest(Endpoint string, Params interface{}, RespStruct interface{}) {
	if !config.Get("gocqhttp.enable").(bool) || !config.Get("gocqhttp.enable_ws").(bool) {
		return
	}
	// Prepare request data.
	Endpoint = strings.TrimPrefix(Endpoint, "/")
	timeUnixNano := strconv.FormatInt(time.Now().UnixNano(), 10)
	echo := "nyabot_" + Endpoint + "_" + timeUnixNano
	requestData := wsRequestData{
		Action: Endpoint,
		Params: Params,
		Echo:   echo,
	}
	// Send request.
	requestChan <- requestData
	// Wait for response.
	for data := range responseChan {
		if data.Echo == echo {
			// Decode response data from wsResponseData to RespStruct.
			jsonData, err := json.Marshal(data)
			err = json.Unmarshal(jsonData, &RespStruct)
			if err != nil {
				log.Println("ws response decode error: ", err)
				return
			}
			return
		}
	}
}
