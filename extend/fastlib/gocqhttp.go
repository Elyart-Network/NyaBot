package fastlib

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqapi/comMessage"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqapi/friendMessage"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqapi/groupActions"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqapi/groupMessage"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqapi/groupSettings"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqapi/imageActions"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqapi/processRequests"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqapi/voiceActions"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"
	"log"
)

// CqSendMsg send message to group or user.
// Message: Message to send.
// ID: Group ID or User ID.
// IsGroup: true for group, false for user.
func CqSendMsg(Message string, Id int64, IsGroup bool) int32 {
	switch IsGroup {
	case true:
		data := comMessage.SendGroupMsgData{
			GroupID:    Id,
			Message:    Message,
			AutoEscape: false,
		}
		msg, err := comMessage.SendGroupMsg(data)
		if err != nil {
			log.Println("[FastLib] CqSendMsg Error: ", err)
			return 0
		}
		return msg.Data.MessageID
	case false:
		data := comMessage.SendPrivateMsgData{
			UserID:     Id,
			Message:    Message,
			AutoEscape: false,
		}
		msg, err := comMessage.SendPrivateMsg(data)
		if err != nil {
			log.Println("[FastLib] CqSendMsg Error: ", err)
			return 0
		}
		return msg.Data.MessageID
	}
	return 0
}

// CqGetMsg get message from message ID.
func CqGetMsg(MessageID int32) comMessage.GetMsgResp {
	data := comMessage.GetMsgData{
		MessageID: MessageID,
	}
	msg, err := comMessage.GetMsg(data)
	if err != nil {
		log.Println("[FastLib] CqGetMsg Error: ", err)
		return comMessage.GetMsgResp{}
	}
	return msg
}

// CqGetForwardMsg get forward message from message ID.
func CqGetForwardMsg(MessageID string) comMessage.GetForwardMsgResp {
	data := comMessage.GetForwardMsgData{
		MessageID: MessageID,
	}
	msg, err := comMessage.GetForwardMsg(data)
	if err != nil {
		log.Println("[FastLib] CqGetForwardMsg Error: ", err)
		return comMessage.GetForwardMsgResp{}
	}
	return msg
}

// CqGenIdForward generate forward node from message ID.
func CqGenIdForward(MessageID string) cqutil.ForwardIdNode {
	var data = cqutil.ForwardIdData{
		Id: MessageID,
	}
	var forward = cqutil.ForwardIdNode{
		Type: "node",
		Data: data,
	}
	return forward
}

// CqGenCustomForward generate forward node from custom data.
// Name: Name of the sender.
// ID: QQ ID of the sender.
// Content: Content of the message.
func CqGenCustomForward(Name string, Id string, Content string) cqutil.ForwardCustomNode {
	var data = cqutil.ForwardCustomData{
		Name:    Name,
		Uin:     Id,
		Content: Content,
	}
	var forward = cqutil.ForwardCustomNode{
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
		data := comMessage.SendGroupForwardMsgData{
			GroupID:  Id,
			Messages: Messages,
		}
		msg, err := comMessage.SendGroupForwardMsg(data)
		if err != nil {
			log.Println("[FastLib] CqSendForwardMsg Error: ", err)
			return 0
		}
		return msg.Data.MessageID
	case false:
		data := comMessage.SendPrivateForwardMsgData{
			UserID:   Id,
			Messages: Messages,
		}
		msg, err := comMessage.SendPrivateForwardMsg(data)
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
	data := comMessage.DeleteMsgData{
		MessageID: MessageID,
	}
	_, err := comMessage.DeleteMsg(data)
	if err != nil {
		log.Println("[FastLib] CqDeleteMsg Error: ", err)
		return false
	}
	return true
}

// CqGetImage get image from cached file name.
func CqGetImage(FileName string) imageActions.GetImageResp {
	data := imageActions.GetImageData{
		File: FileName,
	}
	msg, err := imageActions.GetImage(data)
	if err != nil {
		log.Println("[FastLib] CqGetImage Error: ", err)
		return imageActions.GetImageResp{}
	}
	return msg
}

// CqGetRecord get record cached from file name.
func CqGetRecord(FileName string) voiceActions.GetRecordResp {
	data := voiceActions.GetRecordData{
		File: FileName,
	}
	msg, err := voiceActions.GetRecord(data)
	if err != nil {
		log.Println("[FastLib] CqGetRecord Error: ", err)
		return voiceActions.GetRecordResp{}
	}
	return msg
}

// CqFriendReq process friend request.
// Flag: Request flag, get from request event.
// Approve: true for approve, false for reject.
// Remark: Remark for friend.
func CqFriendReq(Flag string, Approve bool, Remark string) bool {
	data := processRequests.SetFriendAddRequestData{
		Flag:    Flag,
		Approve: Approve,
		Remark:  Remark,
	}
	_, err := processRequests.SetFriendAddRequest(data)
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
	data := processRequests.SetGroupAddRequestData{
		Flag:    Flag,
		SubType: Type,
		Approve: Approve,
		Reason:  RejectReason,
	}
	_, err := processRequests.SetGroupAddRequest(data)
	if err != nil {
		log.Println("[FastLib] CqGroupAdd Error: ", err)
		return false
	}
	return true
}

// CqGetStrangerInfo get stranger info.
func CqGetStrangerInfo(UserID int64) friendMessage.GetStrangerInfoResp {
	data := friendMessage.GetStrangerInfoData{
		UserID:  UserID,
		NoCache: false,
	}
	msg, err := friendMessage.GetStrangerInfo(data)
	if err != nil {
		log.Println("[FastLib] CqGetStrangerInfo Error: ", err)
		return friendMessage.GetStrangerInfoResp{}
	}
	return msg
}

// CqGetGroupMemberInfo get group member info.
func CqGetGroupMemberInfo(GroupID int64, UserID int64) groupMessage.GetGroupMemberInfoResp {
	data := groupMessage.GetGroupMemberInfoData{
		GroupID: GroupID,
		UserID:  UserID,
		NoCache: false,
	}
	msg, err := groupMessage.GetGroupMemberInfo(data)
	if err != nil {
		log.Println("[FastLib] CqGetGroupMembers Error: ", err)
		return groupMessage.GetGroupMemberInfoResp{}
	}
	return msg
}

// CqGetGroupMembers get group member list.
func CqGetGroupMembers(GroupID int64) groupMessage.GetGroupMemberListResp {
	data := groupMessage.GetGroupMemberListData{
		GroupID: GroupID,
		NoCache: false,
	}
	msg, err := groupMessage.GetGroupMemberList(data)
	if err != nil {
		log.Println("[FastLib] CqGetGroupMembers Error: ", err)
		return groupMessage.GetGroupMemberListResp{}
	}
	return msg
}

// CqGetFriends get friend list.
func CqGetFriends() friendMessage.GetFriendListResp {
	msg, err := friendMessage.GetFriendList()
	if err != nil {
		log.Println("[FastLib] CqGetFriends Error: ", err)
		return friendMessage.GetFriendListResp{}
	}
	return msg
}

// CqGetGroups get group list.
func CqGetGroups() groupMessage.GetGroupListResp {
	data := groupMessage.GetGroupListData{
		NoCache: false,
	}
	msg, err := groupMessage.GetGroupList(data)
	if err != nil {
		log.Println("[FastLib] CqGetGroups Error: ", err)
		return groupMessage.GetGroupListResp{}
	}
	return msg
}

// CqGetGroupInfo get group info.
func CqGetGroupInfo(GroupID int64) groupMessage.GetGroupInfoResp {
	data := groupMessage.GetGroupInfoData{
		GroupID: GroupID,
		NoCache: false,
	}
	msg, err := groupMessage.GetGroupInfo(data)
	if err != nil {
		log.Println("[FastLib] CqGetGroupInfo Error: ", err)
		return groupMessage.GetGroupInfoResp{}
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
		data := groupSettings.SetGroupNameData{
			GroupID:   GroupID,
			GroupName: GroupName,
		}
		_, err := groupSettings.SetGroupName(data)
		if err != nil {
			log.Println("[FastLib] CqSetGroupInfo Error: ", err)
			return false
		}
		return true
	case 2:
		data := groupSettings.SetGroupPortraitData{
			GroupID: GroupID,
			File:    Avatar,
			Cache:   1,
		}
		_, err := groupSettings.SetGroupPortrait(data)
		if err != nil {
			log.Println("[FastLib] CqSetGroupInfo Error: ", err)
			return false
		}
		return true
	case 3:
		data := groupSettings.SetGroupCardData{
			GroupID: GroupID,
			UserID:  UserId,
			Card:    Card,
		}
		_, err := groupSettings.SetGroupCard(data)
		if err != nil {
			log.Println("[FastLib] CqSetGroupInfo Error: ", err)
			return false
		}
		return true
	case 4:
		data := groupSettings.SetGroupSpecialTitleData{
			GroupID:      GroupID,
			UserID:       UserId,
			SpecialTitle: SpecialTitle,
		}
		_, err := groupSettings.SetGroupSpecialTitle(data)
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
	data := groupSettings.SetGroupAdminData{
		GroupID: GroupID,
		UserID:  UserID,
		Enable:  Enable,
	}
	_, err := groupSettings.SetGroupAdmin(data)
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
	data := groupActions.SetGroupBanData{
		GroupID:  GroupID,
		UserID:   UserID,
		Duration: Duration,
	}
	_, err := groupActions.SetGroupBan(data)
	if err != nil {
		log.Println("[FastLib] CqGroupBan Error: ", err)
		return false
	}
	return true
}

// CqGroupMute mute group.
func CqGroupMute(GroupID int64, UnMute bool) bool {
	data := groupActions.SetGroupWholeBanData{
		GroupID: GroupID,
		Enable:  UnMute,
	}
	_, err := groupActions.SetGroupWholeBan(data)
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
		data := groupActions.DeleteEssenceMsgData{
			MessageID: MessageID,
		}
		_, err := groupActions.DeleteEssenceMsg(data)
		if err != nil {
			log.Println("[FastLib] CqGroupEssenceMsg Error: ", err)
			return false
		}
		return true
	case false:
		data := groupActions.SetEssenceMsgData{
			MessageID: MessageID,
		}
		_, err := groupActions.SetEssenceMsg(data)
		if err != nil {
			log.Println("[FastLib] CqGroupEssenceMsg Error: ", err)
			return false
		}
		return true
	}
	return false
}

func CqGroupSendNotice(GroupID int64, Content string, Image string) bool {
	data := groupActions.SendGroupNoticeData{
		GroupID: GroupID,
		Content: Content,
		Image:   Image,
	}
	_, err := groupActions.SendGroupNotice(data)
	if err != nil {
		log.Println("[FastLib] CqGroupSendNotice Error: ", err)
		return false
	}
	return true
}

// CqGroupKick kick group member.
func CqGroupKick(GroupID int64, UserID int64, RejectAddRequest bool) bool {
	data := groupActions.SetGroupKickData{
		GroupID: GroupID,
		UserID:  UserID,
		Reject:  RejectAddRequest,
	}
	_, err := groupActions.SetGroupKick(data)
	if err != nil {
		log.Println("[FastLib] CqGroupKick Error: ", err)
		return false
	}
	return true
}

func CqLeaveGroup(GroupID int64) bool {
	data := groupActions.SetGroupLeaveData{
		GroupID:   GroupID,
		IsDismiss: false,
	}
	_, err := groupActions.SetGroupLeave(data)
	if err != nil {
		log.Println("[FastLib] CqLeaveGroup Error: ", err)
		return false
	}
	return true
}

func CqDismissGroup(GroupID int64) bool {
	data := groupActions.SetGroupLeaveData{
		GroupID:   GroupID,
		IsDismiss: true,
	}
	_, err := groupActions.SetGroupLeave(data)
	if err != nil {
		log.Println("[FastLib] CqDismissGroup Error: ", err)
		return false
	}
	return true
}
