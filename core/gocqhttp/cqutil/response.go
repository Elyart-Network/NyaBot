package cqutil

type Response struct {
	Status  string `json:"status"`
	RetCode int    `json:"retcode"`
	Msg     string `json:"msg"`
	Wording string `json:"wording"`
	Echo    string `json:"echo"`
}
