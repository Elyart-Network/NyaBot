package groupMessage

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
