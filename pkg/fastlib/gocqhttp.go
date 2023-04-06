package fastlib

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/common"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/friend"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/group"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/system"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
	"log"
)

// CqSendMsg send message to group or user.
// Message: Message to send.
// ID: Group ID or User ID.
// IsGroup: true for group, false for user.
func CqSendMsg(Message string, Id int64, IsGroup bool) int32 {
	switch IsGroup {
	case true:
		data := types.SendGroupMsgData{
			GroupID:    Id,
			Message:    Message,
			AutoEscape: false,
		}
		msg, err := common.SendGroupMsg(data)
		if err != nil {
			log.Println("[FastLib] CqSendMsg Error: ", err)
			return 0
		}
		return msg.Data.MessageID
	case false:
		data := types.SendPrivateMsgData{
			UserID:     Id,
			Message:    Message,
			AutoEscape: false,
		}
		msg, err := common.SendPrivateMsg(data)
		if err != nil {
			log.Println("[FastLib] CqSendMsg Error: ", err)
			return 0
		}
		return msg.Data.MessageID
	}
	return 0
}

// CqGetMsg get message from message ID.
func CqGetMsg(MessageID int32) types.GetMsgResp {
	data := types.GetMsgData{
		MessageID: MessageID,
	}
	msg, err := common.GetMsg(data)
	if err != nil {
		log.Println("[FastLib] CqGetMsg Error: ", err)
		return types.GetMsgResp{}
	}
	return msg
}

// CqGetForwardMsg get forward message from message ID.
func CqGetForwardMsg(MessageID string) types.GetForwardMsgResp {
	data := types.GetForwardMsgData{
		MessageID: MessageID,
	}
	msg, err := common.GetForwardMsg(data)
	if err != nil {
		log.Println("[FastLib] CqGetForwardMsg Error: ", err)
		return types.GetForwardMsgResp{}
	}
	return msg
}

// CqGenIdForward generate forward node from message ID.
func CqGenIdForward(MessageID string) types.ForwardIdNode {
	var data = types.ForwardIdData{
		Id: MessageID,
	}
	var forward = types.ForwardIdNode{
		Type: "node",
		Data: data,
	}
	return forward
}

// CqGenCustomForward generate forward node from custom data.
// Name: Name of the sender.
// ID: QQ ID of the sender.
// Content: Content of the message.
func CqGenCustomForward(Name string, Id string, Content string) types.ForwardCustomNode {
	var data = types.ForwardCustomData{
		Name:    Name,
		Uin:     Id,
		Content: Content,
	}
	var forward = types.ForwardCustomNode{
		Type: "node",
		Data: data,
	}
	return forward
}

// CqSendForwardMsg send forward message to group or user.
// Message: Message to send.
// ID: Group ID or User ID.
// IsGroup: true for group, false for user.
func CqSendForwardMsg(Messages interface{}, Id int64, IsGroup bool) int64 {
	switch IsGroup {
	case true:
		data := types.SendGroupForwardMsgData{
			GroupID:  Id,
			Messages: Messages,
		}
		msg, err := common.SendGroupForwardMsg(data)
		if err != nil {
			log.Println("[FastLib] CqSendForwardMsg Error: ", err)
			return 0
		}
		return msg.Data.MessageID
	case false:
		data := types.SendPrivateForwardMsgData{
			UserID:   Id,
			Messages: Messages,
		}
		msg, err := common.SendPrivateForwardMsg(data)
		if err != nil {
			log.Println("[FastLib] CqSendForwardMsg Error: ", err)
			return 0
		}
		return msg.Data.MessageID
	}
	return 0
}

// CqDeleteMsg delete message from message ID.
func CqDeleteMsg(MessageID int32) bool {
	data := types.DeleteMsgData{
		MessageID: MessageID,
	}
	_, err := common.DeleteMsg(data)
	if err != nil {
		log.Println("[FastLib] CqDeleteMsg Error: ", err)
		return false
	}
	return true
}

// CqGetImage get image from cached file name.
func CqGetImage(FileName string) types.GetImageResp {
	data := types.GetImageData{
		File: FileName,
	}
	msg, err := common.GetImage(data)
	if err != nil {
		log.Println("[FastLib] CqGetImage Error: ", err)
		return types.GetImageResp{}
	}
	return msg
}

// CqGetRecord get record cached from file name.
func CqGetRecord(FileName string) types.GetRecordResp {
	data := types.GetRecordData{
		File: FileName,
	}
	msg, err := common.GetRecord(data)
	if err != nil {
		log.Println("[FastLib] CqGetRecord Error: ", err)
		return types.GetRecordResp{}
	}
	return msg
}

// CqFriendReq process friend request.
// Flag: Request flag, get from request event.
// Approve: true for approve, false for reject.
// Remark: Remark for friend.
func CqFriendReq(Flag string, Approve bool, Remark string) bool {
	data := types.SetFriendAddRequestData{
		Flag:    Flag,
		Approve: Approve,
		Remark:  Remark,
	}
	_, err := system.SetFriendAddRequest(data)
	if err != nil {
		log.Println("[FastLib] CqFriendAdd Error: ", err)
		return false
	}
	return true
}

// CqGroupReq process group request.
// Flag: Request flag, get from request event.
// Type: Request type, get from request event.
// Approve: true for approve, false for reject.
// RejectReason: Reject reason.
func CqGroupReq(Flag string, Type string, Approve bool, RejectReason string) bool {
	data := types.SetGroupAddRequestData{
		Flag:    Flag,
		SubType: Type,
		Approve: Approve,
		Reason:  RejectReason,
	}
	_, err := system.SetGroupAddRequest(data)
	if err != nil {
		log.Println("[FastLib] CqGroupAdd Error: ", err)
		return false
	}
	return true
}

// CqGetStrangerInfo get stranger info.
func CqGetStrangerInfo(UserID int64) types.GetStrangerInfoResp {
	data := types.GetStrangerInfoData{
		UserID:  UserID,
		NoCache: false,
	}
	msg, err := friend.GetStrangerInfo(data)
	if err != nil {
		log.Println("[FastLib] CqGetStrangerInfo Error: ", err)
		return types.GetStrangerInfoResp{}
	}
	return msg
}

// CqGetGroupMemberInfo get group member info.
func CqGetGroupMemberInfo(GroupID int64, UserID int64) types.GetGroupMemberInfoResp {
	data := types.GetGroupMemberInfoData{
		GroupID: GroupID,
		UserID:  UserID,
		NoCache: false,
	}
	msg, err := group.GetGroupMemberInfo(data)
	if err != nil {
		log.Println("[FastLib] CqGetGroupMembers Error: ", err)
		return types.GetGroupMemberInfoResp{}
	}
	return msg
}

// CqGetGroupMembers get group member list.
func CqGetGroupMembers(GroupID int64) types.GetGroupMemberListResp {
	data := types.GetGroupMemberListData{
		GroupID: GroupID,
		NoCache: false,
	}
	msg, err := group.GetGroupMemberList(data)
	if err != nil {
		log.Println("[FastLib] CqGetGroupMembers Error: ", err)
		return types.GetGroupMemberListResp{}
	}
	return msg
}

// CqGetFriends get friend list.
func CqGetFriends() types.GetFriendListResp {
	msg, err := friend.GetFriendList()
	if err != nil {
		log.Println("[FastLib] CqGetFriends Error: ", err)
		return types.GetFriendListResp{}
	}
	return msg
}

// CqGetGroups get group list.
func CqGetGroups() types.GetGroupListResp {
	data := types.GetGroupListData{
		NoCache: false,
	}
	msg, err := group.GetGroupList(data)
	if err != nil {
		log.Println("[FastLib] CqGetGroups Error: ", err)
		return types.GetGroupListResp{}
	}
	return msg
}

// CqGetGroupInfo get group info.
func CqGetGroupInfo(GroupID int64) types.GetGroupInfoResp {
	data := types.GetGroupInfoData{
		GroupID: GroupID,
		NoCache: false,
	}
	msg, err := group.GetGroupInfo(data)
	if err != nil {
		log.Println("[FastLib] CqGetGroupInfo Error: ", err)
		return types.GetGroupInfoResp{}
	}
	return msg
}

// CqSetGroupInfo set group info.
// Type 1: GroupID, GroupName set group name.
// Type 2: GroupID, Avatar(path) set group avatar.
// Type 3: GroupID, UserId, Card set group card.
// Type 4: GroupID, UserId, SpecialTitle set group member special title.
func CqSetGroupInfo(GroupID int64, GroupName string, Avatar string, UserId int64, Card string, SpecialTitle string, Type int) bool {
	switch Type {
	case 1:
		data := types.SetGroupNameData{
			GroupID:   GroupID,
			GroupName: GroupName,
		}
		_, err := group.SetGroupName(data)
		if err != nil {
			log.Println("[FastLib] CqSetGroupInfo Error: ", err)
			return false
		}
		return true
	case 2:
		data := types.SetGroupPortraitData{
			GroupID: GroupID,
			File:    Avatar,
			Cache:   1,
		}
		_, err := group.SetGroupPortrait(data)
		if err != nil {
			log.Println("[FastLib] CqSetGroupInfo Error: ", err)
			return false
		}
		return true
	case 3:
		data := types.SetGroupCardData{
			GroupID: GroupID,
			UserID:  UserId,
			Card:    Card,
		}
		_, err := group.SetGroupCard(data)
		if err != nil {
			log.Println("[FastLib] CqSetGroupInfo Error: ", err)
			return false
		}
		return true
	case 4:
		data := types.SetGroupSpecialTitleData{
			GroupID:      GroupID,
			UserID:       UserId,
			SpecialTitle: SpecialTitle,
		}
		_, err := group.SetGroupSpecialTitle(data)
		if err != nil {
			log.Println("[FastLib] CqSetGroupInfo Error: ", err)
			return false
		}
		return true
	}
	return false
}

// CqSetGroupAdmin set group admin.
func CqSetGroupAdmin(GroupID int64, UserID int64, Enable bool) bool {
	data := types.SetGroupAdminData{
		GroupID: GroupID,
		UserID:  UserID,
		Enable:  Enable,
	}
	_, err := group.SetGroupAdmin(data)
	if err != nil {
		log.Println("[FastLib] CqSetGroupAdmin Error: ", err)
		return false
	}
	return true
}

// CqGroupBan ban group member.
// BanAll true: ban all member.
// DeBan true: unban.
// Duration only work when BanAll is false.
func CqGroupBan(GroupID int64, UserID int64, Duration uint32, DeBan bool) bool {
	if DeBan {
		Duration = 0
	}
	data := types.SetGroupBanData{
		GroupID:  GroupID,
		UserID:   UserID,
		Duration: Duration,
	}
	_, err := group.SetGroupBan(data)
	if err != nil {
		log.Println("[FastLib] CqGroupBan Error: ", err)
		return false
	}
	return true
}

// CqGroupMute mute group.
func CqGroupMute(GroupID int64, UnMute bool) bool {
	data := types.SetGroupWholeBanData{
		GroupID: GroupID,
		Enable:  UnMute,
	}
	_, err := group.SetGroupWholeBan(data)
	if err != nil {
		log.Println("[FastLib] CqGroupBan Error: ", err)
		return false
	}
	return true
}

// CqGroupEssenceMsg set group essence message.
func CqGroupEssenceMsg(MessageID int32, Remove bool) bool {
	switch Remove {
	case true:
		data := types.DeleteEssenceMsgData{
			MessageID: MessageID,
		}
		_, err := group.DeleteEssenceMsg(data)
		if err != nil {
			log.Println("[FastLib] CqGroupEssenceMsg Error: ", err)
			return false
		}
		return true
	case false:
		data := types.SetEssenceMsgData{
			MessageID: MessageID,
		}
		_, err := group.SetEssenceMsg(data)
		if err != nil {
			log.Println("[FastLib] CqGroupEssenceMsg Error: ", err)
			return false
		}
		return true
	}
	return false
}

func CqGroupSendNotice(GroupID int64, Content string, Image string) bool {
	data := types.SendGroupNoticeData{
		GroupID: GroupID,
		Content: Content,
		Image:   Image,
	}
	_, err := group.SendGroupNotice(data)
	if err != nil {
		log.Println("[FastLib] CqGroupSendNotice Error: ", err)
		return false
	}
	return true
}

// CqGroupKick kick group member.
func CqGroupKick(GroupID int64, UserID int64, RejectAddRequest bool) bool {
	data := types.SetGroupKickData{
		GroupID: GroupID,
		UserID:  UserID,
		Reject:  RejectAddRequest,
	}
	_, err := group.SetGroupKick(data)
	if err != nil {
		log.Println("[FastLib] CqGroupKick Error: ", err)
		return false
	}
	return true
}

func CqLeaveGroup(GroupID int64) bool {
	data := types.SetGroupLeaveData{
		GroupID:   GroupID,
		IsDismiss: false,
	}
	_, err := group.SetGroupLeave(data)
	if err != nil {
		log.Println("[FastLib] CqLeaveGroup Error: ", err)
		return false
	}
	return true
}

func CqDismissGroup(GroupID int64) bool {
	data := types.SetGroupLeaveData{
		GroupID:   GroupID,
		IsDismiss: true,
	}
	_, err := group.SetGroupLeave(data)
	if err != nil {
		log.Println("[FastLib] CqDismissGroup Error: ", err)
		return false
	}
	return true
}
