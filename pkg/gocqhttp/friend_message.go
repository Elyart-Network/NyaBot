package gocqhttp

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/request"
)

func GetStrangerInfo(data models.GetStrangerInfoData) (resp models.GetStrangerInfoResp, err error) {
	err = request.PostRequest("/get_stranger_info", data, &resp)
	return
}

func GetFriendList() (resp models.GetFriendListResp, err error) {
	err = request.GetRequest("/get_friend_list", &resp)
	return
}

func GetUnidirectionalFriendList() (resp models.GetUnidirectionalFriendListResp, err error) {
	err = request.GetRequest("/get_unidirectional_friend_list", &resp)
	return
}
