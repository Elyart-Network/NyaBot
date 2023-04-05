package gocqhttp

import (
	models2 "github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/request"
)

func DeleteMsg(data models2.DeleteMsgData) (resp models2.Response, err error) {
	err = request.PostRequest("/delete_msg", data, &resp)
	return
}

func GetForwardMsg(data models2.GetForwardMsgData) (resp models2.GetForwardMsgResp, err error) {
	err = request.PostRequest("/get_forward_msg", data, &resp)
	return
}

func GetGroupMsgHistory(data models2.GetGroupMsgHistoryData) (resp models2.GetGroupMsgHistoryResp, err error) {
	err = request.PostRequest("/get_group_msg_history", data, &resp)
	return
}

func GetMsg(data models2.GetMsgData) (resp models2.GetMsgResp, err error) {
	err = request.PostRequest("/get_msg", data, &resp)
	return
}

func MarkMsgAsRead(data models2.MarkMsgAsReadData) (resp models2.Response, err error) {
	err = request.PostRequest("/mark_msg_as_read", data, &resp)
	return
}

func SendGroupForwardMsg(data models2.SendGroupForwardMsgData) (resp models2.SendGroupForwardMsgResp, err error) {
	err = request.PostRequest("/send_group_forward_msg", data, &resp)
	return
}

func SendGroupMsg(data models2.SendGroupMsgData) (resp models2.SendGroupMsgResp, err error) {
	err = request.PostRequest("/send_group_msg", data, &resp)
	return
}

func SendPrivateForwardMsg(data models2.SendPrivateForwardMsgData) (resp models2.SendPrivateForwardMsgResp, err error) {
	err = request.PostRequest("/send_private_forward_msg", data, &resp)
	return
}

func SendPrivateMsg(data models2.SendPrivateMsgData) (resp models2.SendPrivateMsgResp, err error) {
	err = request.PostRequest("/send_private_msg", data, &resp)
	return
}
