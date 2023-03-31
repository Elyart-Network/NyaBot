package comMessage

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"
)

func DeleteMsg(data DeleteMsgData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/delete_msg", data, &resp)
	return
}

func GetForwardMsg(data GetForwardMsgData) (resp GetForwardMsgResp, err error) {
	err = cqutil.PostRequest("/get_forward_msg", data, &resp)
	return
}

func GetGroupMsgHistory(data GetGroupMsgHistoryData) (resp GetGroupMsgHistoryResp, err error) {
	err = cqutil.PostRequest("/get_group_msg_history", data, &resp)
	return
}

func GetMsg(data GetMsgData) (resp GetMsgResp, err error) {
	err = cqutil.PostRequest("/get_msg", data, &resp)
	return
}

func MarkMsgAsRead(data MarkMsgAsReadData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/mark_msg_as_read", data, &resp)
	return
}

func SendGroupForwardMsg(data SendGroupForwardMsgData) (resp SendGroupForwardMsgResp, err error) {
	err = cqutil.PostRequest("/send_group_forward_msg", data, &resp)
	return
}

func SendGroupMsg(data SendGroupMsgData) (resp SendGroupMsgResp, err error) {
	err = cqutil.PostRequest("/send_group_msg", data, &resp)
	return
}

func SendPrivateForwardMsg(data SendPrivateForwardMsgData) (resp SendPrivateForwardMsgResp, err error) {
	err = cqutil.PostRequest("/send_private_forward_msg", data, &resp)
	return
}

func SendPrivateMsg(data SendPrivateMsgData) (resp SendPrivateMsgResp, err error) {
	err = cqutil.PostRequest("/send_private_msg", data, &resp)
	return
}
