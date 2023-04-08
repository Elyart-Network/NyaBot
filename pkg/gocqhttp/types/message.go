package types

// Common Message
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

// Friend Message
type GetStrangerInfoData struct {
	UserID  int64 `json:"user_id"`
	NoCache bool  `json:"no_cache"`
}

type GetStrangerInfoResp struct {
	Response
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
	Response
	Data []struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
		Remark   string `json:"remark"`
	} `json:"data"`
}

type GetUnidirectionalFriendListResp struct {
	Response
	Data []struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
		Source   string `json:"source"`
	} `json:"data"`
}

// Group Message
type GetGroupInfoData struct {
	GroupID int64 `json:"group_id"`
	NoCache bool  `json:"no_cache"`
}

type GetGroupListData struct {
	NoCache bool `json:"no_cache"`
}

type GetGroupMemberInfoData struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	NoCache bool  `json:"no_cache"`
}

type GetGroupMemberListData struct {
	GroupID int64 `json:"group_id"`
	NoCache bool  `json:"no_cache"`
}

type GetGroupHonorInfoData struct {
	GroupID int64  `json:"group_id"`
	Type    string `json:"type"`
}

type GetEssenceMsgListData struct {
	GroupID int64 `json:"group_id"`
}

type GetGroupAtAllRemainData struct {
	GroupID int64 `json:"group_id"`
}

type GetGroupInfoResp struct {
	Response
	Data GroupInfoObject `json:"data"`
}

type GetGroupListResp struct {
	Response
	Data []GroupInfoObject `json:"data"`
}

type GetGroupMemberInfoResp struct {
	Response
	Data GroupMemberInfoObject `json:"data"`
}

type GetGroupMemberListResp struct {
	Response
	Data []GroupMemberInfoObject `json:"data"`
}

type GetGroupHonorInfoResp struct {
	Response
	Data struct {
		GroupID          int64                  `json:"group_id"`
		CurrentTalkative CurrentTalkativeObject `json:"current_talkative"`
		TalkativeList    []HonorListObject      `json:"talkative_list"`
		PerformerList    []HonorListObject      `json:"performer_list"`
		LegendList       []HonorListObject      `json:"legend_list"`
		StrongNewbieList []HonorListObject      `json:"strong_newbie_list"`
		EmotionList      []HonorListObject      `json:"emotion_list"`
	} `json:"data"`
}

type GetGroupSystemMsgResp struct {
	Response
	Data struct {
		InvitedRequests []InvitedRequestObject `json:"invited_requests"`
		JoinRequests    []JoinRequestObject    `json:"join_requests"`
	} `json:"data"`
}

type GetEssenceMsgListResp struct {
	Response
	Data []struct {
		SenderID     int64  `json:"sender_id"`
		SenderNick   string `json:"sender_nick"`
		SenderTime   int64  `json:"sender_time"`
		OperatorID   int64  `json:"operator_id"`
		OperatorNick string `json:"operator_nick"`
		OperatorTime int64  `json:"operator_time"`
		MessageID    int32  `json:"message_id"`
	} `json:"data"`
}

type GetGroupAtAllRemainResp struct {
	Response
	Data struct {
		CanAtAll                 bool  `json:"can_at_all"`
		RemainAtAllCountForGroup int16 `json:"remain_at_all_count_for_group"`
		RemainAtAllCountForUin   int16 `json:"remain_at_all_count_for_uin"`
	} `json:"data"`
}

type GroupInfoObject struct {
	GroupID         int64  `json:"group_id"`
	GroupName       string `json:"group_name"`
	GroupMemo       string `json:"group_memo"`
	GroupCreateTime uint32 `json:"group_create_time"`
	GroupLevel      uint32 `json:"group_level"`
	MemberCount     int32  `json:"member_count"`
	MaxMemberCount  int32  `json:"max_member_count"`
}

type GroupMemberInfoObject struct {
	GroupID         int64  `json:"group_id"`
	UserID          int64  `json:"user_id"`
	Nickname        string `json:"nickname"`
	Card            string `json:"card"`
	Sex             string `json:"sex"`
	Age             int32  `json:"age"`
	Area            string `json:"area"`
	JoinTime        int32  `json:"join_time"`
	LastSentTime    int32  `json:"last_sent_time"`
	Level           string `json:"level"`
	Role            string `json:"role"`
	Unfriendly      bool   `json:"unfriendly"`
	Title           string `json:"title"`
	TitleExpireTime int64  `json:"title_expire_time"`
	CardChangeable  bool   `json:"card_changeable"`
	ShutUpTimestamp int64  `json:"shut_up_timestamp"`
}

type CurrentTalkativeObject struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	DayCount int32  `json:"day_count"`
}

type HonorListObject struct {
	UserID      int64  `json:"user_id"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

type InvitedRequestObject struct {
	RequestID   int64  `json:"request_id"`
	InvitorUin  int64  `json:"invitor_uin"`
	InvitorNick string `json:"invitor_nick"`
	GroupID     int64  `json:"group_id"`
	GroupName   string `json:"group_name"`
	Checked     bool   `json:"checked"`
	Actor       int64  `json:"actor"`
}

type JoinRequestObject struct {
	RequestID     int64  `json:"request_id"`
	RequesterUin  int64  `json:"requester_uin"`
	RequesterNick string `json:"requester_nick"`
	Message       string `json:"message"`
	GroupID       int64  `json:"group_id"`
	GroupName     string `json:"group_name"`
	Checked       bool   `json:"checked"`
	Actor         int64  `json:"actor"`
}
