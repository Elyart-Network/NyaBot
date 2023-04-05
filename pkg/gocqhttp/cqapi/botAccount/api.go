package botAccount

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqutil"
)

func GetModelShow(data GetModelShowData) (resp GetModelShowResp, err error) {
	err = cqutil.PostRequest("/_get_model_show", data, &resp)
	return
}

func SetModelShow(data SetModelShowData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/_set_model_show", data, &resp)
	return
}

func GetLoginInfo() (resp GetLoginInfoResp, err error) {
	err = cqutil.GetRequest("/get_login_info", &resp)
	return
}

func GetOnlineClients(data GetOnlineClientsData) (resp GetOnlineClientsResp, err error) {
	err = cqutil.PostRequest("/get_online_clients", data, &resp)
	return
}

func QiDianGetAccountInfo() (resp cqutil.Response, err error) {
	err = cqutil.GetRequest("/qidian_get_account_info", &resp)
	return
}

func SetQQProfile(data SetQQProfileData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_qq_profile", data, &resp)
	return
}
