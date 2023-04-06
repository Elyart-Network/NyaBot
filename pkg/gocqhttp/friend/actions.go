package friend

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
)

func DeleteFriend(data types.DeleteFriendData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/delete_friend", data, &resp)
	return
}

func DeleteUnidirectionalFriend(data types.DeleteUnidirectionalFriendData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/delete_unidirectional_friend", data, &resp)
	return
}
