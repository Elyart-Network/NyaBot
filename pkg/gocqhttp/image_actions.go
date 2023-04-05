package gocqhttp

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/request"
)

func GetImage(data models.GetImageData) (resp models.GetImageResp, err error) {
	err = request.PostRequest("/get_image", data, &resp)
	return
}

func CanSendImage() (resp models.CanSendImageResp, err error) {
	err = request.GetRequest("/can_send_image", &resp)
	return
}

func OcrImage(data models.OcrImageData) (resp models.OcrImageResp, err error) {
	err = request.PostRequest("/ocr_image", data, &resp)
	return
}
