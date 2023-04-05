package comMessage

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
