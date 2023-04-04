package fileActions

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"
)

type GetGroupFileSystemInfoResp struct {
	cqutil.Response
	Data struct {
		FileCount  int32 `json:"file_count"`
		LimitCount int32 `json:"limit_count"`
		UsedSpace  int64 `json:"used_space"`
		TotalSpace int64 `json:"total_space"`
	} `json:"data"`
}

type GetGroupRootFilesResp struct {
	cqutil.Response
	Data struct {
		Files   []FileObject   `json:"files"`
		Folders []FolderObject `json:"folders"`
	} `json:"data"`
}

type GetGroupFilesByFolderResp struct {
	cqutil.Response
	Data struct {
		Files   []FileObject   `json:"files"`
		Folders []FolderObject `json:"folders"`
	} `json:"data"`
}

type GetGroupFileUrlResp struct {
	cqutil.Response
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}

type FileObject struct {
	GroupID       int32  `json:"group_id"`
	FileID        string `json:"file_id"`
	FileName      string `json:"file_name"`
	BusID         int32  `json:"busid"`
	FileSize      int64  `json:"file_size"`
	UploadTime    int64  `json:"upload_time"`
	DeadTime      int64  `json:"dead_time"`
	ModifyTime    int64  `json:"modify_time"`
	DownloadTimes int32  `json:"download_times"`
	Uploader      int64  `json:"uploader"`
	UploaderName  string `json:"uploader_name"`
}

type FolderObject struct {
	GroupID        int32  `json:"group_id"`
	FolderID       string `json:"folder_id"`
	FolderName     string `json:"folder_name"`
	CreateTime     int64  `json:"create_time"`
	Creator        int64  `json:"creator"`
	CreatorName    string `json:"creator_name"`
	TotalFileCount int32  `json:"total_file_count"`
}
