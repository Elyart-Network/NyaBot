package webhook

type Data struct {
	Source    Source  `json:"source"`
	Target    Target  `json:"target"`
	Message   Message `json:"message"`
	Callback  bool    `json:"callback"`
	TimeStamp int64   `json:"time_stamp"`
}

type Source struct {
	// "GoCqHttp"/"Mirai"/"OneBot"/"Telegram"/"Discord"/"Slack"
	Platform string      `json:"platform"`
	Addition interface{} `json:"addition"`
}

type Target struct {
	GoCqHttp
}

type GoCqHttp struct {
	// "Group"/"Private"
	Type string `json:"type"`
	// Group ID/ Private ID
	ID int64 `json:"id"`
}

type Message struct {
	Type     string      `json:"type"`
	Content  string      `json:"content"`
	Addition interface{} `json:"addition"`
}
