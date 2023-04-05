package gocqhttp

import (
	models2 "github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/request"
)

func SetFriendAddRequest(data models2.SetFriendAddRequestData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_friend_add_request", data, &resp)
	return
}

func SetGroupAddRequest(data models2.SetGroupAddRequestData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_group_add_request", data, &resp)
	return
}
