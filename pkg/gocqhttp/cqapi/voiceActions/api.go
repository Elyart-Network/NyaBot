package voiceActions

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqutil"
)

func GetRecord(data GetRecordData) (resp GetRecordResp, err error) {
	err = cqutil.PostRequest("/get_record", data, &resp)
	return
}

func CanSendRecord() (resp CanSendRecordResp, err error) {
	err = cqutil.GetRequest("/can_send_record", &resp)
	return
}
