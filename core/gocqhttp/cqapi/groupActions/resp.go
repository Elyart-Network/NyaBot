package groupActions

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"
)

type GetGroupNoticeResp struct {
	cqutil.Response
	Data struct {
		SenderID    int64         `json:"sender_id"`
		PublishTime int64         `json:"publish_time"`
		Message     MessageObject `json:"message"`
	} `json:"data"`
}

type MessageObject struct {
	Text   string        `json:"text"`
	Images []ImageObject `json:"images"`
}

type ImageObject struct {
	Height string `json:"height"`
	Width  string `json:"width"`
	Id     string `json:"id"`
}
