package processRequests

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"
)

func SetFriendAddRequest(data SetFriendAddRequestData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_friend_add_request", data, &resp)
	return
}

func SetGroupAddRequest(data SetGroupAddRequestData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_group_add_request", data, &resp)
	return
}
