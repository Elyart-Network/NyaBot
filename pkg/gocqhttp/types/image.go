package types

type GetImageData struct {
	File string `json:"file"`
}

type OcrImageData struct {
	Image string `json:"image"`
}

type GetImageResp struct {
	Response
	Data struct {
		Size     int32  `json:"size"`
		FileName string `json:"filename"`
		Url      string `json:"url"`
	} `json:"data"`
}

type CanSendImageResp struct {
	Response
	Data struct {
		Yes bool `json:"yes"`
	} `json:"data"`
}

type OcrImageResp struct {
	Response
	Data struct {
		Texts    []TextDetectionObject `json:"texts"`
		Language string                `json:"language"`
	} `json:"data"`
}

type TextDetectionObject struct {
	Text        string `json:"text"`
	Confidence  int32  `json:"confidence"`
	Coordinates []any  `json:"coordinates"`
}
