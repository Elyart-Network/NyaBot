package fileActions

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
