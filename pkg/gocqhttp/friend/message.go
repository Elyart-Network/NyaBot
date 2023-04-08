package friend

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
)

func GetStrangerInfo(data types.GetStrangerInfoData) (resp types.GetStrangerInfoResp, err error) {
	err = gocqhttp.PostRequest("/get_stranger_info", data, &resp)
	return
}

func GetFriendList() (resp types.GetFriendListResp, err error) {
	err = gocqhttp.GetRequest("/get_friend_list", &resp)
	return
}

func GetUnidirectionalFriendList() (resp types.GetUnidirectionalFriendListResp, err error) {
	err = gocqhttp.GetRequest("/get_unidirectional_friend_list", &resp)
	return
}
