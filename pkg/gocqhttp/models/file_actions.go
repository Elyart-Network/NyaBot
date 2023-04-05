package models

type UploadGroupFileData struct {
	GroupID int64  `json:"group_id"`
	File    string `json:"file"`
	Name    string `json:"name"`
	Folder  string `json:"folder"`
}

type DeleteGroupFileData struct {
	GroupID int64  `json:"group_id"`
	FileID  string `json:"file_id"`
	BusID   int32  `json:"busid"`
}

type CreateGroupFileFolderData struct {
	GroupID  int64  `json:"group_id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

type DeleteGroupFolderData struct {
	GroupID  int64  `json:"group_id"`
	FolderID string `json:"folder_id"`
}

type GetGroupFileSystemInfoData struct {
	GroupID int64 `json:"group_id"`
}

type GetGroupRootFilesData struct {
	GroupID int64 `json:"group_id"`
}

type GetGroupFilesByFolderData struct {
	GroupID  int64  `json:"group_id"`
	FolderID string `json:"folder_id"`
}

type GetGroupFileUrlData struct {
	GroupID int64  `json:"group_id"`
	FileID  string `json:"file_id"`
	BusID   int32  `json:"busid"`
}

type UploadPrivateFileData struct {
	UserID int64  `json:"user_id"`
	File   string `json:"file"`
	Name   string `json:"name"`
}

type GetGroupFileSystemInfoResp struct {
	Response
	Data struct {
		FileCount  int32 `json:"file_count"`
		LimitCount int32 `json:"limit_count"`
		UsedSpace  int64 `json:"used_space"`
		TotalSpace int64 `json:"total_space"`
	} `json:"data"`
}

type GetGroupRootFilesResp struct {
	Response
	Data struct {
		Files   []FileObject   `json:"files"`
		Folders []FolderObject `json:"folders"`
	} `json:"data"`
}

type GetGroupFilesByFolderResp struct {
	Response
	Data struct {
		Files   []FileObject   `json:"files"`
		Folders []FolderObject `json:"folders"`
	} `json:"data"`
}

type GetGroupFileUrlResp struct {
	Response
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
