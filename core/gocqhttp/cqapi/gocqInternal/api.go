package gocqInternal

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"
)

func GetCookies(data GetCookiesData) (resp GetCookiesResp, err error) {
	err = cqutil.PostRequest("/get_cookies", data, &resp)
	return
}

func GetCsrfToken() (resp GetCsrfTokenResp, err error) {
	err = cqutil.GetRequest("/get_csrf_token", &resp)
	return
}

func GetCredentials() (resp GetCredentialsResp, err error) {
	err = cqutil.GetRequest("/get_credentials", &resp)
	return
}

func GetVersionInfo() (resp GetVersionInfoResp, err error) {
	err = cqutil.GetRequest("/get_version_info", &resp)
	return
}

func GetStatus() (resp GetStatusResp, err error) {
	err = cqutil.GetRequest("/get_status", &resp)
	return
}

func SetRestart(data SetRestartData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/set_restart", data, &resp)
	return
}

func CleanCache() (resp cqutil.Response, err error) {
	err = cqutil.GetRequest("/clean_cache", &resp)
	return
}

func ReloadEventFilter(data ReloadEventFilterData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/reload_event_filter", data, &resp)
	return
}

func DownloadFile(data DownloadFileData) (resp DownloadFileResp, err error) {
	err = cqutil.PostRequest("/download_file", data, &resp)
	return
}

func CheckUrlSafely(data CheckUrlSafelyData) (resp CheckUrlSafelyResp, err error) {
	err = cqutil.PostRequest("/check_url_safely", data, &resp)
	return
}

func GetWordSlices(data GetWordSlicesData) (resp GetWordSlicesResp, err error) {
	err = cqutil.PostRequest("/get_word_slices", data, &resp)
	return
}

func HandleQuickOperation(data HandleQuickOperationData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/handle_quick_operation", data, &resp)
	return
}
