package friendMessage

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"
)

type GetStrangerInfoResp struct {
	cqutil.Response
	Data struct {
		UserID    int64  `json:"user_id"`
		Nickname  string `json:"nickname"`
		Sex       string `json:"sex"`
		Age       int32  `json:"age"`
		Qid       string `json:"qid"`
		Level     int32  `json:"level"`
		LoginDays int32  `json:"login_days"`
	} `json:"data"`
}

type GetFriendListResp struct {
	cqutil.Response
	Data []struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
		Remark   string `json:"remark"`
	} `json:"data"`
}

type GetUnidirectionalFriendListResp struct {
	cqutil.Response
	Data []struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
		Source   string `json:"source"`
	} `json:"data"`
}
