package gocqhttp

import (
	models2 "github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/request"
)

func GetModelShow(data models2.GetModelShowData) (resp models2.GetModelShowResp, err error) {
	err = request.PostRequest("/_get_model_show", data, &resp)
	return
}

func SetModelShow(data models2.SetModelShowData) (resp models2.Response, err error) {
	err = request.PostRequest("/_set_model_show", data, &resp)
	return
}

func GetLoginInfo() (resp models2.GetLoginInfoResp, err error) {
	err = request.GetRequest("/get_login_info", &resp)
	return
}

func GetOnlineClients(data models2.GetOnlineClientsData) (resp models2.GetOnlineClientsResp, err error) {
	err = request.PostRequest("/get_online_clients", data, &resp)
	return
}

func QiDianGetAccountInfo() (resp models2.Response, err error) {
	err = request.GetRequest("/qidian_get_account_info", &resp)
	return
}

func SetQQProfile(data models2.SetQQProfileData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_qq_profile", data, &resp)
	return
}
