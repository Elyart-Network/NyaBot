package friendMessage

type GetStrangerInfoData struct {
	UserID  int64 `json:"user_id"`
	NoCache bool  `json:"no_cache"`
}
