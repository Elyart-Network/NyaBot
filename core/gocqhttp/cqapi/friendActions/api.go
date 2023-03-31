package friendActions

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"
)

func DeleteFriend(data DeleteFriendData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/delete_friend", data, &resp)
	return
}

func DeleteUnidirectionalFriend(data DeleteUnidirectionalFriendData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/delete_unidirectional_friend", data, &resp)
	return
}
