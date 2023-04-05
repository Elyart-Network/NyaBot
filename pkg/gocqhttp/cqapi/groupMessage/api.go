package groupMessage

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqutil"
)

func GetGroupInfo(data GetGroupInfoData) (resp GetGroupInfoResp, err error) {
	err = cqutil.PostRequest("/get_group_info", data, &resp)
	return
}

func GetGroupList(data GetGroupListData) (resp GetGroupListResp, err error) {
	err = cqutil.PostRequest("/get_group_list", data, &resp)
	return
}

func GetGroupMemberInfo(data GetGroupMemberInfoData) (resp GetGroupMemberInfoResp, err error) {
	err = cqutil.PostRequest("/get_group_member_info", data, &resp)
	return
}

func GetGroupMemberList(data GetGroupMemberListData) (resp GetGroupMemberListResp, err error) {
	err = cqutil.PostRequest("/get_group_member_list", data, &resp)
	return
}

func GetGroupHonorInfo(data GetGroupHonorInfoData) (resp GetGroupHonorInfoResp, err error) {
	err = cqutil.PostRequest("/get_group_honor_info", data, &resp)
	return
}

func GetGroupSystemMsg() (resp GetGroupSystemMsgResp, err error) {
	err = cqutil.GetRequest("/get_group_system_msg", &resp)
	return
}

func GetEssenceMsgList(data GetEssenceMsgListData) (resp GetEssenceMsgListResp, err error) {
	err = cqutil.PostRequest("/get_essence_msg_list", data, &resp)
	return
}

func GetGroupAtAllRemain(data GetGroupAtAllRemainData) (resp GetGroupAtAllRemainResp, err error) {
	err = cqutil.PostRequest("/get_group_at_all_remain", data, &resp)
	return
}
