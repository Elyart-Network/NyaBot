package gocqhttp

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/request"
)

func GetGroupInfo(data models.GetGroupInfoData) (resp models.GetGroupInfoResp, err error) {
	err = request.PostRequest("/get_group_info", data, &resp)
	return
}

func GetGroupList(data models.GetGroupListData) (resp models.GetGroupListResp, err error) {
	err = request.PostRequest("/get_group_list", data, &resp)
	return
}

func GetGroupMemberInfo(data models.GetGroupMemberInfoData) (resp models.GetGroupMemberInfoResp, err error) {
	err = request.PostRequest("/get_group_member_info", data, &resp)
	return
}

func GetGroupMemberList(data models.GetGroupMemberListData) (resp models.GetGroupMemberListResp, err error) {
	err = request.PostRequest("/get_group_member_list", data, &resp)
	return
}

func GetGroupHonorInfo(data models.GetGroupHonorInfoData) (resp models.GetGroupHonorInfoResp, err error) {
	err = request.PostRequest("/get_group_honor_info", data, &resp)
	return
}

func GetGroupSystemMsg() (resp models.GetGroupSystemMsgResp, err error) {
	err = request.GetRequest("/get_group_system_msg", &resp)
	return
}

func GetEssenceMsgList(data models.GetEssenceMsgListData) (resp models.GetEssenceMsgListResp, err error) {
	err = request.PostRequest("/get_essence_msg_list", data, &resp)
	return
}

func GetGroupAtAllRemain(data models.GetGroupAtAllRemainData) (resp models.GetGroupAtAllRemainResp, err error) {
	err = request.PostRequest("/get_group_at_all_remain", data, &resp)
	return
}
