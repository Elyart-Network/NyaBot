package websocket

import (
	"encoding/json"
	"strings"
)

type wsResponseData struct {
	Status  string      `json:"status"`
	RetCode int         `json:"retcode"`
	Msg     string      `json:"msg"`
	Wording string      `json:"wording"`
	Echo    string      `json:"echo"`
	Data    interface{} `json:"data"`
}

var responseChan = make(chan wsResponseData)

func wsResponseEncode(wsContext []byte) (resp wsResponseData, isResp bool, err error) {
	// Encode Message to wsActionData.
	err = json.Unmarshal(wsContext, &resp)
	if echo := resp.Echo; echo != "" && strings.HasPrefix(echo, "nyabot_") {
		return resp, true, err
	}
	return wsResponseData{}, false, err
}

func wsResponse(Response wsResponseData) {
	// Send response to channel.
	responseChan <- Response
}
