package gocqhttp

import (
	models2 "github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/request"
)

func SetGroupBan(data models2.SetGroupBanData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_group_ban", data, &resp)
	return
}

func SetGroupWholeBan(data models2.SetGroupWholeBanData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_group_whole_ban", data, &resp)
	return
}

func SetGroupAnonymousBan(data models2.SetGroupAnonymousBanData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_group_anonymous_ban", data, &resp)
	return
}

func SetEssenceMsg(data models2.SetEssenceMsgData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_essence_msg", data, &resp)
	return
}

func DeleteEssenceMsg(data models2.DeleteEssenceMsgData) (resp models2.Response, err error) {
	err = request.PostRequest("/delete_essence_msg", data, &resp)
	return
}

func SendGroupSign(data models2.SendGroupSignData) (resp models2.Response, err error) {
	err = request.PostRequest("/send_group_sign", data, &resp)
	return
}

func SetGroupAnonymous(data models2.SetGroupAnonymousData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_group_anonymous", data, &resp)
	return
}

func SendGroupNotice(data models2.SendGroupNoticeData) (resp models2.Response, err error) {
	err = request.PostRequest("/send_group_notice", data, &resp)
	return
}

func GetGroupNotice(data models2.GetGroupNoticeData) (resp models2.GetGroupNoticeResp, err error) {
	err = request.PostRequest("/get_group_notice", data, &resp)
	return
}

func SetGroupKick(data models2.SetGroupKickData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_group_kick", data, &resp)
	return
}

func SetGroupLeave(data models2.SetGroupLeaveData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_group_leave", data, &resp)
	return
}
