package imageActions

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"
)

type GetImageResp struct {
	cqutil.Response
	Data struct {
		Size     int32  `json:"size"`
		FileName string `json:"filename"`
		Url      string `json:"url"`
	} `json:"data"`
}

type CanSendImageResp struct {
	cqutil.Response
	Data struct {
		Yes bool `json:"yes"`
	} `json:"data"`
}

type OcrImageResp struct {
	cqutil.Response
	Data struct {
		Texts    []TextDetectionObject `json:"texts"`
		Language string                `json:"language"`
	} `json:"data"`
}

type TextDetectionObject struct {
	Text        string        `json:"text"`
	Confidence  int32         `json:"confidence"`
	Coordinates []interface{} `json:"coordinates"`
}
