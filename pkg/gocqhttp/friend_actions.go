package gocqhttp

import (
	models2 "github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/request"
)

func DeleteFriend(data models2.DeleteFriendData) (resp models2.Response, err error) {
	err = request.PostRequest("/delete_friend", data, &resp)
	return
}

func DeleteUnidirectionalFriend(data models2.DeleteUnidirectionalFriendData) (resp models2.Response, err error) {
	err = request.PostRequest("/delete_unidirectional_friend", data, &resp)
	return
}
