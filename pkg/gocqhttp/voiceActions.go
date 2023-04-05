package gocqhttp

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/request"
)

func GetRecord(data models.GetRecordData) (resp models.GetRecordResp, err error) {
	err = request.PostRequest("/get_record", data, &resp)
	return
}

func CanSendRecord() (resp models.CanSendRecordResp, err error) {
	err = request.GetRequest("/can_send_record", &resp)
	return
}
