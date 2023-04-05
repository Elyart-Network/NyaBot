package models

type DeleteMsgData struct {
	MessageID int32 `json:"message_id"`
}

type GetForwardMsgData struct {
	MessageID string `json:"message_id"`
}

type GetGroupMsgHistoryData struct {
	MessageSeq int64 `json:"message_seq"`
	GroupID    int64 `json:"group_id"`
}

type GetMsgData struct {
	MessageID int32 `json:"message_id"`
}

type MarkMsgAsReadData struct {
	MessageID int32 `json:"message_id"`
}

type SendGroupForwardMsgData struct {
	GroupID  int64       `json:"group_id"`
	Messages interface{} `json:"messages"`
}

type SendGroupMsgData struct {
	GroupID    int64  `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

type SendPrivateForwardMsgData struct {
	UserID   int64       `json:"user_id"`
	Messages interface{} `json:"messages"`
}

type SendPrivateMsgData struct {
	UserID     int64  `json:"user_id"`
	Message    string `json:"message"`
	GroupID    int64  `json:"group_id"`
	AutoEscape bool   `json:"auto_escape"`
}

type GetForwardMsgResp struct {
	Response
	Data struct {
		Messages []ForwardMessagesObject `json:"messages"`
	} `json:"data"`
}

type GetGroupMsgHistoryResp struct {
	Response
	Data struct {
		Messages []interface{} `json:"messages"`
	} `json:"data"`
}

type GetMsgResp struct {
	Response
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
	Response
	Data struct {
		MessageID int64  `json:"message_id"`
		ForwardID string `json:"forward_id"`
	} `json:"data"`
}

type SendGroupMsgResp struct {
	Response
	Data struct {
		MessageID int32 `json:"message_id"`
	} `json:"data"`
}

type SendPrivateForwardMsgResp struct {
	Response
	Data struct {
		MessageID int64  `json:"message_id"`
		ForwardID string `json:"forward_id"`
	} `json:"data"`
}

type SendPrivateMsgResp struct {
	Response
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
