package gocqhttp

import (
	models2 "github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/request"
)

func UploadGroupFile(data models2.UploadGroupFileData) (resp models2.Response, err error) {
	err = request.PostRequest("/upload_group_file", data, &resp)
	return
}

func DeleteGroupFile(data models2.DeleteGroupFileData) (resp models2.Response, err error) {
	err = request.PostRequest("/delete_group_file", data, &resp)
	return
}

func DeleteGroupFolder(data models2.DeleteGroupFolderData) (resp models2.Response, err error) {
	err = request.PostRequest("/delete_group_folder", data, &resp)
	return
}

func GetGroupFileSystemInfo(data models2.GetGroupFileSystemInfoData) (resp models2.GetGroupFileSystemInfoResp, err error) {
	err = request.PostRequest("/get_group_file_system_info", data, &resp)
	return
}

func GetGroupRootFiles(data models2.GetGroupRootFilesData) (resp models2.GetGroupRootFilesResp, err error) {
	err = request.PostRequest("/get_group_root_files", data, &resp)
	return
}

func GetGroupFilesByFolder(data models2.GetGroupFilesByFolderData) (resp models2.GetGroupFilesByFolderResp, err error) {
	err = request.PostRequest("/get_group_files_by_folder", data, &resp)
	return
}

func GetGroupFileUrl(data models2.GetGroupFileUrlData) (resp models2.GetGroupFileUrlResp, err error) {
	err = request.PostRequest("/get_group_file_url", data, &resp)
	return
}

func UploadPrivateFile(data models2.UploadPrivateFileData) (resp models2.Response, err error) {
	err = request.PostRequest("/upload_private_file", data, &resp)
	return
}
