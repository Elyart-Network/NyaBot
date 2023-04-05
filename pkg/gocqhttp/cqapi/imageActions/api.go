package imageActions

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqutil"
)

func GetImage(data GetImageData) (resp GetImageResp, err error) {
	err = cqutil.PostRequest("/get_image", data, &resp)
	return
}

func CanSendImage() (resp CanSendImageResp, err error) {
	err = cqutil.GetRequest("/can_send_image", &resp)
	return
}

func OcrImage(data OcrImageData) (resp OcrImageResp, err error) {
	err = cqutil.PostRequest("/ocr_image", data, &resp)
	return
}
