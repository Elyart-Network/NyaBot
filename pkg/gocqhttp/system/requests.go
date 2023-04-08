package system

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
)

func SetFriendAddRequest(data types.SetFriendAddRequestData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_friend_add_request", data, &resp)
	return
}

func SetGroupAddRequest(data types.SetGroupAddRequestData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_group_add_request", data, &resp)
	return
}
