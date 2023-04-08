package group

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
)

func SetGroupBan(data types.SetGroupBanData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_group_ban", data, &resp)
	return
}

func SetGroupWholeBan(data types.SetGroupWholeBanData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_group_whole_ban", data, &resp)
	return
}

func SetGroupAnonymousBan(data types.SetGroupAnonymousBanData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_group_anonymous_ban", data, &resp)
	return
}

func SetEssenceMsg(data types.SetEssenceMsgData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_essence_msg", data, &resp)
	return
}

func DeleteEssenceMsg(data types.DeleteEssenceMsgData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/delete_essence_msg", data, &resp)
	return
}

func SendGroupSign(data types.SendGroupSignData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/send_group_sign", data, &resp)
	return
}

func SetGroupAnonymous(data types.SetGroupAnonymousData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_group_anonymous", data, &resp)
	return
}

func SendGroupNotice(data types.SendGroupNoticeData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/send_group_notice", data, &resp)
	return
}

func GetGroupNotice(data types.GetGroupNoticeData) (resp types.GetGroupNoticeResp, err error) {
	err = gocqhttp.PostRequest("/get_group_notice", data, &resp)
	return
}

func SetGroupKick(data types.SetGroupKickData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_group_kick", data, &resp)
	return
}

func SetGroupLeave(data types.SetGroupLeaveData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_group_leave", data, &resp)
	return
}
