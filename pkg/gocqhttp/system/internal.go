package system

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
)

func GetCookies(data types.GetCookiesData) (resp types.GetCookiesResp, err error) {
	err = gocqhttp.PostRequest("/get_cookies", data, &resp)
	return
}

func GetCsrfToken() (resp types.GetCsrfTokenResp, err error) {
	err = gocqhttp.GetRequest("/get_csrf_token", &resp)
	return
}

func GetCredentials() (resp types.GetCredentialsResp, err error) {
	err = gocqhttp.GetRequest("/get_credentials", &resp)
	return
}

func GetVersionInfo() (resp types.GetVersionInfoResp, err error) {
	err = gocqhttp.GetRequest("/get_version_info", &resp)
	return
}

func GetStatus() (resp types.GetStatusResp, err error) {
	err = gocqhttp.GetRequest("/get_status", &resp)
	return
}

func SetRestart(data types.SetRestartData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/set_restart", data, &resp)
	return
}

func CleanCache() (resp types.Response, err error) {
	err = gocqhttp.GetRequest("/clean_cache", &resp)
	return
}

func ReloadEventFilter(data types.ReloadEventFilterData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/reload_event_filter", data, &resp)
	return
}

func DownloadFile(data types.DownloadFileData) (resp types.DownloadFileResp, err error) {
	err = gocqhttp.PostRequest("/download_file", data, &resp)
	return
}

func CheckUrlSafely(data types.CheckUrlSafelyData) (resp types.CheckUrlSafelyResp, err error) {
	err = gocqhttp.PostRequest("/check_url_safely", data, &resp)
	return
}

func GetWordSlices(data types.GetWordSlicesData) (resp types.GetWordSlicesResp, err error) {
	err = gocqhttp.PostRequest("/get_word_slices", data, &resp)
	return
}

func HandleQuickOperation(data types.HandleQuickOperationData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/handle_quick_operation", data, &resp)
	return
}
