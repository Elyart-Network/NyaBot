package gocqhttp

import (
	models2 "github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/request"
)

func SetGroupName(data models2.SetGroupNameData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_group_name", data, &resp)
	return
}

func SetGroupPortrait(data models2.SetGroupPortraitData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_group_portrait", data, &resp)
	return
}

func SetGroupAdmin(data models2.SetGroupAdminData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_group_admin", data, &resp)
	return
}

func SetGroupCard(data models2.SetGroupCardData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_group_card", data, &resp)
	return
}

func SetGroupSpecialTitle(data models2.SetGroupSpecialTitleData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_group_special_title", data, &resp)
	return
}
