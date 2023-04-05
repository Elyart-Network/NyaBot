package models

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
