package system

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
)

func GetModelShow(data types.GetModelShowData) (resp types.GetModelShowResp, err error) {
	err = gocqhttp.PostRequest("/_get_model_show", data, &resp)
	return
}

func SetModelShow(data types.SetModelShowData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/_set_model_show", data, &resp)
	return
}

func GetLoginInfo() (resp types.GetLoginInfoResp, err error) {
	err = gocqhttp.GetRequest("/get_login_info", &resp)
	return
}

func GetOnlineClients(data types.GetOnlineClientsData) (resp types.GetOnlineClientsResp, err error) {
	err = gocqhttp.PostRequest("/get_online_clients", data, &resp)
	return
}

func QiDianGetAccountInfo() (resp types.Response, err error) {
	err = gocqhttp.GetRequest("/qidian_get_account_info", &resp)
	return
}

func SetQQProfile(data types.SetQQProfileData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_qq_profile", data, &resp)
	return
}
