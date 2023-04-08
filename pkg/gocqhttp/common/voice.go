package common

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
)

func GetRecord(data types.GetRecordData) (resp types.GetRecordResp, err error) {
	err = gocqhttp.PostRequest("/get_record", data, &resp)
	return
}

func CanSendRecord() (resp types.CanSendRecordResp, err error) {
	err = gocqhttp.GetRequest("/can_send_record", &resp)
	return
}
