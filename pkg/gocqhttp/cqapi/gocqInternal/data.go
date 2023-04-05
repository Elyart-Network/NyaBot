package gocqInternal

type GetCookiesData struct {
	Domain string `json:"domain"`
}

type GetCredentialsData struct {
	Domain string `json:"domain"`
}

type SetRestartData struct {
	Delay int `json:"delay"`
}

type ReloadEventFilterData struct {
	File string `json:"file"`
}

type DownloadFileData struct {
	Url         string            `json:"url"`
	ThreadCount int32             `json:"thread_count"`
	Headers     map[string]string `json:"headers"`
}

type CheckUrlSafelyData struct {
	Url string `json:"url"`
}

type GetWordSlicesData struct {
	Content string `json:"content"`
}

type HandleQuickOperationData struct {
	Context   interface{} `json:"context"`
	Operation interface{} `json:"operation"`
}
