package types

type SetGroupNameData struct {
	GroupID   int64  `json:"group_id"`
	GroupName string `json:"group_name"`
}

type SetGroupPortraitData struct {
	GroupID int64  `json:"group_id"`
	File    string `json:"file"`
	Cache   int    `json:"cache"`
}

type SetGroupAdminData struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	Enable  bool  `json:"enable"`
}

type SetGroupCardData struct {
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	Card    string `json:"card"`
}

type SetGroupSpecialTitleData struct {
	GroupID      int64  `json:"group_id"`
	UserID       int64  `json:"user_id"`
	SpecialTitle string `json:"special_title"`
	Duration     uint32 `json:"duration"`
}
