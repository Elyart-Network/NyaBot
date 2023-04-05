package models

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
