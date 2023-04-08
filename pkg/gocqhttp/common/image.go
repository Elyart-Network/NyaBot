package common

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
)

func GetImage(data types.GetImageData) (resp types.GetImageResp, err error) {
	err = gocqhttp.PostRequest("/get_image", data, &resp)
	return
}

func CanSendImage() (resp types.CanSendImageResp, err error) {
	err = gocqhttp.GetRequest("/can_send_image", &resp)
	return
}

func OcrImage(data types.OcrImageData) (resp types.OcrImageResp, err error) {
	err = gocqhttp.PostRequest("/ocr_image", data, &resp)
	return
}
