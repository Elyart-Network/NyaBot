package types

type GetRecordData struct {
	File      string `json:"file"`
	OutFormat string `json:"out_format"`
}

type GetRecordResp struct {
	Response
	Data struct {
		File string `json:"file"`
	} `json:"data"`
}

type CanSendRecordResp struct {
	Response
	Data struct {
		Yes bool `json:"yes"`
	} `json:"data"`
}
