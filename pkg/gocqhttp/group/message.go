package group

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
)

func GetGroupInfo(data types.GetGroupInfoData) (resp types.GetGroupInfoResp, err error) {
	err = gocqhttp.PostRequest("/get_group_info", data, &resp)
	return
}

func GetGroupList(data types.GetGroupListData) (resp types.GetGroupListResp, err error) {
	err = gocqhttp.PostRequest("/get_group_list", data, &resp)
	return
}

func GetGroupMemberInfo(data types.GetGroupMemberInfoData) (resp types.GetGroupMemberInfoResp, err error) {
	err = gocqhttp.PostRequest("/get_group_member_info", data, &resp)
	return
}

func GetGroupMemberList(data types.GetGroupMemberListData) (resp types.GetGroupMemberListResp, err error) {
	err = gocqhttp.PostRequest("/get_group_member_list", data, &resp)
	return
}

func GetGroupHonorInfo(data types.GetGroupHonorInfoData) (resp types.GetGroupHonorInfoResp, err error) {
	err = gocqhttp.PostRequest("/get_group_honor_info", data, &resp)
	return
}

func GetGroupSystemMsg() (resp types.GetGroupSystemMsgResp, err error) {
	err = gocqhttp.GetRequest("/get_group_system_msg", &resp)
	return
}

func GetEssenceMsgList(data types.GetEssenceMsgListData) (resp types.GetEssenceMsgListResp, err error) {
	err = gocqhttp.PostRequest("/get_essence_msg_list", data, &resp)
	return
}

func GetGroupAtAllRemain(data types.GetGroupAtAllRemainData) (resp types.GetGroupAtAllRemainResp, err error) {
	err = gocqhttp.PostRequest("/get_group_at_all_remain", data, &resp)
	return
}
