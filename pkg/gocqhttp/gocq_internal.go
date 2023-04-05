package gocqhttp

import (
	models2 "github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/request"
)

func GetCookies(data models2.GetCookiesData) (resp models2.GetCookiesResp, err error) {
	err = request.PostRequest("/get_cookies", data, &resp)
	return
}

func GetCsrfToken() (resp models2.GetCsrfTokenResp, err error) {
	err = request.GetRequest("/get_csrf_token", &resp)
	return
}

func GetCredentials() (resp models2.GetCredentialsResp, err error) {
	err = request.GetRequest("/get_credentials", &resp)
	return
}

func GetVersionInfo() (resp models2.GetVersionInfoResp, err error) {
	err = request.GetRequest("/get_version_info", &resp)
	return
}

func GetStatus() (resp models2.GetStatusResp, err error) {
	err = request.GetRequest("/get_status", &resp)
	return
}

func SetRestart(data models2.SetRestartData) (resp models2.Response, err error) {
	err = request.PostRequest("/set_restart", data, &resp)
	return
}

func CleanCache() (resp models2.Response, err error) {
	err = request.GetRequest("/clean_cache", &resp)
	return
}

func ReloadEventFilter(data models2.ReloadEventFilterData) (resp models2.Response, err error) {
	err = request.PostRequest("/reload_event_filter", data, &resp)
	return
}

func DownloadFile(data models2.DownloadFileData) (resp models2.DownloadFileResp, err error) {
	err = request.PostRequest("/download_file", data, &resp)
	return
}

func CheckUrlSafely(data models2.CheckUrlSafelyData) (resp models2.CheckUrlSafelyResp, err error) {
	err = request.PostRequest("/check_url_safely", data, &resp)
	return
}

func GetWordSlices(data models2.GetWordSlicesData) (resp models2.GetWordSlicesResp, err error) {
	err = request.PostRequest("/get_word_slices", data, &resp)
	return
}

func HandleQuickOperation(data models2.HandleQuickOperationData) (resp models2.Response, err error) {
	err = request.PostRequest("/handle_quick_operation", data, &resp)
	return
}
