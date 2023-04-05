package fileActions

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqutil"
)

func UploadGroupFile(data UploadGroupFileData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/upload_group_file", data, &resp)
	return
}

func DeleteGroupFile(data DeleteGroupFileData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/delete_group_file", data, &resp)
	return
}

func DeleteGroupFolder(data DeleteGroupFolderData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/delete_group_folder", data, &resp)
	return
}

func GetGroupFileSystemInfo(data GetGroupFileSystemInfoData) (resp GetGroupFileSystemInfoResp, err error) {
	err = cqutil.PostRequest("/get_group_file_system_info", data, &resp)
	return
}

func GetGroupRootFiles(data GetGroupRootFilesData) (resp GetGroupRootFilesResp, err error) {
	err = cqutil.PostRequest("/get_group_root_files", data, &resp)
	return
}

func GetGroupFilesByFolder(data GetGroupFilesByFolderData) (resp GetGroupFilesByFolderResp, err error) {
	err = cqutil.PostRequest("/get_group_files_by_folder", data, &resp)
	return
}

func GetGroupFileUrl(data GetGroupFileUrlData) (resp GetGroupFileUrlResp, err error) {
	err = cqutil.PostRequest("/get_group_file_url", data, &resp)
	return
}

func UploadPrivateFile(data UploadPrivateFileData) (resp cqutil.Response, err error) {
	err = cqutil.PostRequest("/upload_private_file", data, &resp)
	return
}
