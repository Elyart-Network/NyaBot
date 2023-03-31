package groupActions

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"
)

func SetGroupBan(data SetGroupBanData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_group_ban", data, &resp)
	return
}

func SetGroupWholeBan(data SetGroupWholeBanData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_group_whole_ban", data, &resp)
	return
}

func SetGroupAnonymousBan(data SetGroupAnonymousBanData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_group_anonymous_ban", data, &resp)
	return
}

func SetEssenceMsg(data SetEssenceMsgData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_essence_msg", data, &resp)
	return
}

func DeleteEssenceMsg(data DeleteEssenceMsgData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/delete_essence_msg", data, &resp)
	return
}

func SendGroupSign(data SendGroupSignData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/send_group_sign", data, &resp)
	return
}

func SetGroupAnonymous(data SetGroupAnonymousData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_group_anonymous", data, &resp)
	return
}

func SendGroupNotice(data SendGroupNoticeData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/send_group_notice", data, &resp)
	return
}

func GetGroupNotice(data GetGroupNoticeData) (resp GetGroupNoticeResp, err error) {
	err = cqutil.PostRequest("/get_group_notice", data, &resp)
	return
}

func SetGroupKick(data SetGroupKickData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_group_kick", data, &resp)
	return
}

func SetGroupLeave(data SetGroupLeaveData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_group_leave", data, &resp)
	return
}
