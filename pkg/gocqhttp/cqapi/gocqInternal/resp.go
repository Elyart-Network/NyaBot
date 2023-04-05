package gocqInternal

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqutil"
)

type GetCookiesResp struct {
	cqutil.Response
	Data struct {
		Cookies string `json:"cookies"`
	} `json:"data"`
}

type GetCsrfTokenResp struct {
	cqutil.Response
	Data struct {
		Token int `json:"token"`
	} `json:"data"`
}

type GetCredentialsResp struct {
	cqutil.Response
	Data struct {
		Cookies   string `json:"cookies"`
		CsrfToken int32  `json:"csrf_token"`
	} `json:"data"`
}

type GetVersionInfoResp struct {
	cqutil.Response
	Data struct {
		AppName                  string `json:"app_name"`
		AppVersion               string `json:"app_version"`
		AppFullName              string `json:"app_full_name"`
		ProtocolVersion          string `json:"protocol_version"`
		CoolqEdition             string `json:"coolq_edition"`
		CoolqDirectory           string `json:"coolq_directory"`
		GoCqHttp                 bool   `json:"go-cqhttp"`
		PluginVersion            string `json:"plugin_version"`
		PluginBuildNumber        int    `json:"plugin_build_number"`
		PluginBuildConfiguration string `json:"plugin_build_configuration"`
		RuntimeVersion           string `json:"runtime_version"`
		RuntimeOs                string `json:"runtime_os"`
		Version                  string `json:"version"`
		Protocol                 int    `json:"protocol"`
	} `json:"data"`
}

type GetStatusResp struct {
	cqutil.Response
	Data struct {
		AppInitialized bool             `json:"app_initialized"`
		AppEnabled     bool             `json:"app_enabled"`
		PluginsGood    bool             `json:"plugins_good"`
		AppGood        bool             `json:"app_good"`
		Online         bool             `json:"online"`
		Good           bool             `json:"good"`
		Stat           StatisticsObject `json:"stat"`
	} `json:"data"`
}

type StatisticsObject struct {
	PacketReceived  uint64 `json:"PacketReceived"`
	PacketSent      uint64 `json:"PacketSent"`
	PacketLost      uint32 `json:"PacketLost"`
	MessageReceived uint64 `json:"MessageReceived"`
	MessageSent     uint64 `json:"MessageSent"`
	DisconnectTimes uint32 `json:"DisconnectTimes"`
	LostTimes       uint32 `json:"LostTimes"`
	LastMessageTime int64  `json:"LastMessageTime"`
}

type DownloadFileResp struct {
	cqutil.Response
	Data struct {
		File string `json:"file"`
	} `json:"data"`
}

type CheckUrlSafelyResp struct {
	cqutil.Response
	Data struct {
		Level int `json:"level"`
	} `json:"data"`
}

type GetWordSlicesResp struct {
	cqutil.Response
	Data struct {
		Slices []string `json:"slices"`
	} `json:"data"`
}
