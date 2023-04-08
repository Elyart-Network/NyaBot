package types

type GetModelShowData struct {
	Model string `json:"model"`
}

type SetModelShowData struct {
	Model     string `json:"model"`
	ModelShow string `json:"model_show"`
}

type GetOnlineClientsData struct {
	NoCache bool `json:"no_cache"`
}

type SetQQProfileData struct {
	Nickname     string `json:"nickname"`
	Company      string `json:"company"`
	Email        string `json:"email"`
	College      string `json:"college"`
	PersonalNote string `json:"personal_note"`
}

type GetModelShowResp struct {
	Response
	Data struct {
		Variants []VariantsObject `json:"variants"`
	} `json:"data"`
}

type GetLoginInfoResp struct {
	Response
	Data struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
	} `json:"data"`
}

type GetOnlineClientsResp struct {
	Response
	Data struct {
		Clients []DeviceObject `json:"clients"`
	} `json:"data"`
}

type VariantsObject struct {
	ModelShow string `json:"model_show"`
	NeedPay   bool   `json:"need_pay"`
}

type DeviceObject struct {
	AppID      int64  `json:"app_id"`
	DeviceName string `json:"device_name"`
	DeviceKind string `json:"device_kind"`
}
