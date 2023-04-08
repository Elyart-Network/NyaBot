package types

import "github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"

// Friend Actions
type DeleteFriendData struct {
	UserID int64 `json:"user_id"`
}

type DeleteUnidirectionalFriendData struct {
	UserID int64 `json:"user_id"`
}

// Group Actions
type SetGroupBanData struct {
	GroupID  int64  `json:"group_id"`
	UserID   int64  `json:"user_id"`
	Duration uint32 `json:"duration"`
}

type SetGroupWholeBanData struct {
	GroupID int64 `json:"group_id"`
	Enable  bool  `json:"enable"`
}

type SetGroupAnonymousBanData struct {
	GroupID   int64                    `json:"group_id"`
	Anonymous callback.AnonymousObject `json:"anonymous"`
	Flag      string                   `json:"flag"`
	Duration  uint32                   `json:"duration"`
}

type SetEssenceMsgData struct {
	MessageID int32 `json:"message_id"`
}

type DeleteEssenceMsgData struct {
	MessageID int32 `json:"message_id"`
}

type SendGroupSignData struct {
	GroupID int64 `json:"group_id"`
}

type SetGroupAnonymousData struct {
	GroupID int64 `json:"group_id"`
	Enable  bool  `json:"enable"`
}

type SendGroupNoticeData struct {
	GroupID int64  `json:"group_id"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

type GetGroupNoticeData struct {
	GroupID int64 `json:"group_id"`
}

type SetGroupKickData struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	Reject  bool  `json:"reject_add_request"`
}

type SetGroupLeaveData struct {
	GroupID   int64 `json:"group_id"`
	IsDismiss bool  `json:"is_dismiss"`
}

type GetGroupNoticeResp struct {
	Response
	Data struct {
		SenderID    int64         `json:"sender_id"`
		PublishTime int64         `json:"publish_time"`
		Message     MessageObject `json:"message"`
	} `json:"data"`
}

type MessageObject struct {
	Text   string        `json:"text"`
	Images []ImageObject `json:"images"`
}

type ImageObject struct {
	Height string `json:"height"`
	Width  string `json:"width"`
	Id     string `json:"id"`
}
