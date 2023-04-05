package comMessage

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqutil"
)

type GetForwardMsgResp struct {
	cqutil.Response
	Data struct {
		Messages []ForwardMessagesObject `json:"messages"`
	} `json:"data"`
}

type GetGroupMsgHistoryResp struct {
	cqutil.Response
	Data struct {
		Messages []interface{} `json:"messages"`
	} `json:"data"`
}

type GetMsgResp struct {
	cqutil.Response
	Data struct {
		Group       bool         `json:"group"`
		GroupID     int64        `json:"group_id"`
		MessageID   int32        `json:"message_id"`
		RealID      int32        `json:"real_id"`
		MessageType string       `json:"message_type"`
		Sender      SenderObject `json:"sender"`
		Time        int          `json:"time"`
		Message     interface{}  `json:"message"`
		RawMessage  interface{}  `json:"raw_message"`
	} `json:"data"`
}

type SendGroupForwardMsgResp struct {
	cqutil.Response
	Data struct {
		MessageID int64  `json:"message_id"`
		ForwardID string `json:"forward_id"`
	} `json:"data"`
}

type SendGroupMsgResp struct {
	cqutil.Response
	Data struct {
		MessageID int32 `json:"message_id"`
	} `json:"data"`
}

type SendPrivateForwardMsgResp struct {
	cqutil.Response
	Data struct {
		MessageID int64  `json:"message_id"`
		ForwardID string `json:"forward_id"`
	} `json:"data"`
}

type SendPrivateMsgResp struct {
	cqutil.Response
	Data struct {
		MessageID int32 `json:"message_id"`
	} `json:"data"`
}

type ForwardMessagesObject struct {
	Content string       `json:"content"`
	Sender  SenderObject `json:"sender"`
	Time    int          `json:"time"`
}

type SenderObject struct {
	Nickname string `json:"nickname"`
	UserID   int64  `json:"user_id"`
}
