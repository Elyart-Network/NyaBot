package botAccount

import "github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"

type GetModelShowResp struct {
	cqutil.Response
	Data struct {
		Variants []VariantsObject `json:"variants"`
	} `json:"data"`
}

type GetLoginInfoResp struct {
	cqutil.Response
	Data struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
	} `json:"data"`
}

type GetOnlineClientsResp struct {
	cqutil.Response
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
