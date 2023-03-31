package groupActions

import "github.com/Elyart-Network/NyaBot/core/gocqhttp/cqcall"

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
	GroupID   int64                  `json:"group_id"`
	Anonymous cqcall.AnonymousObject `json:"anonymous"`
	Flag      string                 `json:"flag"`
	Duration  uint32                 `json:"duration"`
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
