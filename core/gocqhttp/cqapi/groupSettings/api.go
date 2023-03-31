package groupSettings

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"
)

func SetGroupName(data SetGroupNameData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_group_name", data, &resp)
	return
}

func SetGroupPortrait(data SetGroupPortraitData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_group_portrait", data, &resp)
	return
}

func SetGroupAdmin(data SetGroupAdminData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_group_admin", data, &resp)
	return
}

func SetGroupCard(data SetGroupCardData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_group_card", data, &resp)
	return
}

func SetGroupSpecialTitle(data SetGroupSpecialTitleData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_group_special_title", data, &resp)
	return
}
