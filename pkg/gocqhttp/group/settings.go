package group

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
)

func SetGroupName(data types.SetGroupNameData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_group_name", data, &resp)
	return
}

func SetGroupPortrait(data types.SetGroupPortraitData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_group_portrait", data, &resp)
	return
}

func SetGroupAdmin(data types.SetGroupAdminData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_group_admin", data, &resp)
	return
}

func SetGroupCard(data types.SetGroupCardData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_group_card", data, &resp)
	return
}

func SetGroupSpecialTitle(data types.SetGroupSpecialTitleData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_group_special_title", data, &resp)
	return
}
