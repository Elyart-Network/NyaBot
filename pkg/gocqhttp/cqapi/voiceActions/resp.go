package voiceActions

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqutil"
)

type GetRecordResp struct {
	cqutil.Response
	Data struct {
		File string `json:"file"`
	} `json:"data"`
}

type CanSendRecordResp struct {
	cqutil.Response
	Data struct {
		Yes bool `json:"yes"`
	} `json:"data"`
}
