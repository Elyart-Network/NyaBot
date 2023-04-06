package common

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
)

func UploadGroupFile(data types.UploadGroupFileData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/upload_group_file", data, &resp)
	return
}

func DeleteGroupFile(data types.DeleteGroupFileData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/delete_group_file", data, &resp)
	return
}

func DeleteGroupFolder(data types.DeleteGroupFolderData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/delete_group_folder", data, &resp)
	return
}

func GetGroupFileSystemInfo(data types.GetGroupFileSystemInfoData) (resp types.GetGroupFileSystemInfoResp, err error) {
	err = gocqhttp.PostRequest("/get_group_file_system_info", data, &resp)
	return
}

func GetGroupRootFiles(data types.GetGroupRootFilesData) (resp types.GetGroupRootFilesResp, err error) {
	err = gocqhttp.PostRequest("/get_group_root_files", data, &resp)
	return
}

func GetGroupFilesByFolder(data types.GetGroupFilesByFolderData) (resp types.GetGroupFilesByFolderResp, err error) {
	err = gocqhttp.PostRequest("/get_group_files_by_folder", data, &resp)
	return
}

func GetGroupFileUrl(data types.GetGroupFileUrlData) (resp types.GetGroupFileUrlResp, err error) {
	err = gocqhttp.PostRequest("/get_group_file_url", data, &resp)
	return
}

func UploadPrivateFile(data types.UploadPrivateFileData) (resp types.Response, err error) {
	err = gocqhttp.PostRequest("/upload_private_file", data, &resp)
	return
}
