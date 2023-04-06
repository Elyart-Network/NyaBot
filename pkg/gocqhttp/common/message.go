package common

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
)

func DeleteMsg(data types.DeleteMsgData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/delete_msg", data, &resp)
	return
}

func GetForwardMsg(data types.GetForwardMsgData) (resp types.GetForwardMsgResp, err error) {
	err = gocqhttp.PostRequest("/get_forward_msg", data, &resp)
	return
}

func GetGroupMsgHistory(data types.GetGroupMsgHistoryData) (resp types.GetGroupMsgHistoryResp, err error) {
	err = gocqhttp.PostRequest("/get_group_msg_history", data, &resp)
	return
}

func GetMsg(data types.GetMsgData) (resp types.GetMsgResp, err error) {
	err = gocqhttp.PostRequest("/get_msg", data, &resp)
	return
}

func MarkMsgAsRead(data types.MarkMsgAsReadData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/mark_msg_as_read", data, &resp)
	return
}

func SendGroupForwardMsg(data types.SendGroupForwardMsgData) (resp types.SendGroupForwardMsgResp, err error) {
	err = gocqhttp.PostRequest("/send_group_forward_msg", data, &resp)
	return
}

func SendGroupMsg(data types.SendGroupMsgData) (resp types.SendGroupMsgResp, err error) {
	err = gocqhttp.PostRequest("/send_group_msg", data, &resp)
	return
}

func SendPrivateForwardMsg(data types.SendPrivateForwardMsgData) (resp types.SendPrivateForwardMsgResp, err error) {
	err = gocqhttp.PostRequest("/send_private_forward_msg", data, &resp)
	return
}

func SendPrivateMsg(data types.SendPrivateMsgData) (resp types.SendPrivateMsgResp, err error) {
	err = gocqhttp.PostRequest("/send_private_msg", data, &resp)
	return
}
