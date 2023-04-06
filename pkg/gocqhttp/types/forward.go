package types

type ForwardIdNode struct {
	Type string        `json:"type"`
	Data ForwardIdData `json:"data"`
}

type ForwardIdData struct {
	Id string `json:"id"`
}

type ForwardCustomNode struct {
	Type string            `json:"type"`
	Data ForwardCustomData `json:"data"`
}

type ForwardCustomData struct {
	Name    string `json:"name"`
	Uin     string `json:"uin"`
	Content string `json:"content"`
	Seq     string `json:"seq"`
	Time    string `json:"time"`
}
