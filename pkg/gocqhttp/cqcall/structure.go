package cqcall

type CallbackFull struct {
	CallbackBase
	CallbackMessage
	CallbackRequest
	CallbackNotice
	CallbackMetaEvent
	CallbackUnique
}

type CallbackBase struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
}

type CallbackMessage struct {
	MessageType string       `json:"message_type"`
	SubType     string       `json:"sub_type"`
	MessageID   int32        `json:"message_id"`
	UserID      int64        `json:"user_id"`
	Message     string       `json:"message"`
	RawMessage  string       `json:"raw_message"`
	Font        int32        `json:"font"`
	Sender      SenderObject `json:"sender"`
}

type CallbackRequest struct {
	RequestType string `json:"request_type"`
}

type CallbackNotice struct {
	NoticeType string `json:"notice_type"`
}

type CallbackMetaEvent struct {
	MetaEventType string `json:"meta_event_type"`
}

type CallbackUnique struct {
	GroupID    int64           `json:"group_id"`
	OperatorID int64           `json:"operator_id"`
	TargetID   int64           `json:"target_id"`
	TempSource int             `json:"temp_source"`
	Anonymous  AnonymousObject `json:"anonymous"`
	UploadFile FileObject      `json:"file"`
	Duration   int64           `json:"duration"`
	SenderID   int64           `json:"sender_id"`
	HonorType  string          `json:"honor_type"`
	Title      string          `json:"title"`
	NewCard    string          `json:"card_new"`
	OldCard    string          `json:"card_old"`
	Client     []DeviceObject  `json:"client"`
	Online     bool            `json:"online"`
	Comment    string          `json:"comment"`
	Flag       string          `json:"flag"`
	Status     StatusObject    `json:"status"`
	Interval   int64           `json:"interval"`
}

type SenderObject struct {
	UserID   int64  `json:"user_id"`
	GroupID  int64  `json:"group_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
	Card     string `json:"card"`
	Area     string `json:"area"`
	Level    string `json:"level"`
	Role     string `json:"role"`
	Title    string `json:"title"`
}

type AnonymousObject struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

type FileObject struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	URL   string `json:"url"`
	BusID int64  `json:"busid"`
}

type DeviceObject struct {
	AppID      int64  `json:"app_id"`
	DeviceName string `json:"device_name"`
	DeviceKind string `json:"device_kind"`
}

type StatusObject struct {
	AppInitialized bool                   `json:"app_initialized"`
	AppEnabled     bool                   `json:"app_enabled"`
	PluginsGood    bool                   `json:"plugins_good"`
	AppGood        bool                   `json:"app_good"`
	Online         bool                   `json:"online"`
	Stat           StatusStatisticsObject `json:"stat"`
}

type StatusStatisticsObject struct {
	PacketReceived  uint64 `json:"PacketReceived"`
	PacketSent      uint64 `json:"PacketSent"`
	PacketLost      uint64 `json:"PacketLost"`
	MessageReceived uint64 `json:"MessageReceived"`
	MessageSent     uint64 `json:"MessageSent"`
	DisconnectTimes uint32 `json:"DisconnectTimes"`
	LostTimes       uint32 `json:"LostTimes"`
	LastMessageTime int64  `json:"LastMessageTime"`
}
