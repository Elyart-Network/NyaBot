package friendActions

type DeleteFriendData struct {
	UserID int64 `json:"user_id"`
}

type DeleteUnidirectionalFriendData struct {
	UserID int64 `json:"user_id"`
}
