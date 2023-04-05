package friendMessage

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqutil"
)

func GetStrangerInfo(data GetStrangerInfoData) (resp GetStrangerInfoResp, err error) {
	err = cqutil.PostRequest("/get_stranger_info", data, &resp)
	return
}

func GetFriendList() (resp GetFriendListResp, err error) {
	err = cqutil.GetRequest("/get_friend_list", &resp)
	return
}

func GetUnidirectionalFriendList() (resp GetUnidirectionalFriendListResp, err error) {
	err = cqutil.GetRequest("/get_unidirectional_friend_list", &resp)
	return
}
