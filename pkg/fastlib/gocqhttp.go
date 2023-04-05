package fastlib

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	models2 "github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"log"
)

// CqSendMsg send message to group or user.
// Message: Message to send.
// ID: Group ID or User ID.
// IsGroup: true for group, false for user.
func CqSendMsg(Message string, Id int64, IsGroup bool) int32 {
	switch IsGroup {
	case true:
		data := models2.SendGroupMsgData{
			GroupID:    Id,
			Message:    Message,
			AutoEscape: false,
		}
		msg, err := gocqhttp.SendGroupMsg(data)
		if err != nil {
			log.Println("[FastLib] CqSendMsg Error: ", err)
			return 0
		}
		return msg.Data.MessageID
	case false:
		data := models2.SendPrivateMsgData{
			UserID:     Id,
			Message:    Message,
			AutoEscape: false,
		}
		msg, err := gocqhttp.SendPrivateMsg(data)
		if err != nil {
			log.Println("[FastLib] CqSendMsg Error: ", err)
			return 0
		}
		return msg.Data.MessageID
	}
	return 0
}

// CqGetMsg get message from message ID.
func CqGetMsg(MessageID int32) models2.GetMsgResp {
	data := models2.GetMsgData{
		MessageID: MessageID,
	}
	msg, err := gocqhttp.GetMsg(data)
	if err != nil {
		log.Println("[FastLib] CqGetMsg Error: ", err)
		return models2.GetMsgResp{}
	}
	return msg
}

// CqGetForwardMsg get forward message from message ID.
func CqGetForwardMsg(MessageID string) models2.GetForwardMsgResp {
	data := models2.GetForwardMsgData{
		MessageID: MessageID,
	}
	msg, err := gocqhttp.GetForwardMsg(data)
	if err != nil {
		log.Println("[FastLib] CqGetForwardMsg Error: ", err)
		return models2.GetForwardMsgResp{}
	}
	return msg
}

// CqGenIdForward generate forward node from message ID.
func CqGenIdForward(MessageID string) models2.ForwardIdNode {
	var data = models2.ForwardIdData{
		Id: MessageID,
	}
	var forward = models2.ForwardIdNode{
		Type: "node",
		Data: data,
	}
	return forward
}

// CqGenCustomForward generate forward node from custom data.
// Name: Name of the sender.
// ID: QQ ID of the sender.
// Content: Content of the message.
func CqGenCustomForward(Name string, Id string, Content string) models2.ForwardCustomNode {
	var data = models2.ForwardCustomData{
		Name:    Name,
		Uin:     Id,
		Content: Content,
	}
	var forward = models2.ForwardCustomNode{
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
		data := models2.SendGroupForwardMsgData{
			GroupID:  Id,
			Messages: Messages,
		}
		msg, err := gocqhttp.SendGroupForwardMsg(data)
		if err != nil {
			log.Println("[FastLib] CqSendForwardMsg Error: ", err)
			return 0
		}
		return msg.Data.MessageID
	case false:
		data := models2.SendPrivateForwardMsgData{
			UserID:   Id,
			Messages: Messages,
		}
		msg, err := gocqhttp.SendPrivateForwardMsg(data)
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
	data := models2.DeleteMsgData{
		MessageID: MessageID,
	}
	_, err := gocqhttp.DeleteMsg(data)
	if err != nil {
		log.Println("[FastLib] CqDeleteMsg Error: ", err)
		return false
	}
	return true
}

// CqGetImage get image from cached file name.
func CqGetImage(FileName string) models2.GetImageResp {
	data := models2.GetImageData{
		File: FileName,
	}
	msg, err := gocqhttp.GetImage(data)
	if err != nil {
		log.Println("[FastLib] CqGetImage Error: ", err)
		return models2.GetImageResp{}
	}
	return msg
}

// CqGetRecord get record cached from file name.
func CqGetRecord(FileName string) models2.GetRecordResp {
	data := models2.GetRecordData{
		File: FileName,
	}
	msg, err := gocqhttp.GetRecord(data)
	if err != nil {
		log.Println("[FastLib] CqGetRecord Error: ", err)
		return models2.GetRecordResp{}
	}
	return msg
}

// CqFriendReq process friend request.
// Flag: Request flag, get from request event.
// Approve: true for approve, false for reject.
// Remark: Remark for friend.
func CqFriendReq(Flag string, Approve bool, Remark string) bool {
	data := models2.SetFriendAddRequestData{
		Flag:    Flag,
		Approve: Approve,
		Remark:  Remark,
	}
	_, err := gocqhttp.SetFriendAddRequest(data)
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
	data := models2.SetGroupAddRequestData{
		Flag:    Flag,
		SubType: Type,
		Approve: Approve,
		Reason:  RejectReason,
	}
	_, err := gocqhttp.SetGroupAddRequest(data)
	if err != nil {
		log.Println("[FastLib] CqGroupAdd Error: ", err)
		return false
	}
	return true
}

// CqGetStrangerInfo get stranger info.
func CqGetStrangerInfo(UserID int64) models2.GetStrangerInfoResp {
	data := models2.GetStrangerInfoData{
		UserID:  UserID,
		NoCache: false,
	}
	msg, err := gocqhttp.GetStrangerInfo(data)
	if err != nil {
		log.Println("[FastLib] CqGetStrangerInfo Error: ", err)
		return models2.GetStrangerInfoResp{}
	}
	return msg
}

// CqGetGroupMemberInfo get group member info.
func CqGetGroupMemberInfo(GroupID int64, UserID int64) models2.GetGroupMemberInfoResp {
	data := models2.GetGroupMemberInfoData{
		GroupID: GroupID,
		UserID:  UserID,
		NoCache: false,
	}
	msg, err := gocqhttp.GetGroupMemberInfo(data)
	if err != nil {
		log.Println("[FastLib] CqGetGroupMembers Error: ", err)
		return models2.GetGroupMemberInfoResp{}
	}
	return msg
}

// CqGetGroupMembers get group member list.
func CqGetGroupMembers(GroupID int64) models2.GetGroupMemberListResp {
	data := models2.GetGroupMemberListData{
		GroupID: GroupID,
		NoCache: false,
	}
	msg, err := gocqhttp.GetGroupMemberList(data)
	if err != nil {
		log.Println("[FastLib] CqGetGroupMembers Error: ", err)
		return models2.GetGroupMemberListResp{}
	}
	return msg
}

// CqGetFriends get friend list.
func CqGetFriends() models2.GetFriendListResp {
	msg, err := gocqhttp.GetFriendList()
	if err != nil {
		log.Println("[FastLib] CqGetFriends Error: ", err)
		return models2.GetFriendListResp{}
	}
	return msg
}

// CqGetGroups get group list.
func CqGetGroups() models2.GetGroupListResp {
	data := models2.GetGroupListData{
		NoCache: false,
	}
	msg, err := gocqhttp.GetGroupList(data)
	if err != nil {
		log.Println("[FastLib] CqGetGroups Error: ", err)
		return models2.GetGroupListResp{}
	}
	return msg
}

// CqGetGroupInfo get group info.
func CqGetGroupInfo(GroupID int64) models2.GetGroupInfoResp {
	data := models2.GetGroupInfoData{
		GroupID: GroupID,
		NoCache: false,
	}
	msg, err := gocqhttp.GetGroupInfo(data)
	if err != nil {
		log.Println("[FastLib] CqGetGroupInfo Error: ", err)
		return models2.GetGroupInfoResp{}
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
		data := models2.SetGroupNameData{
			GroupID:   GroupID,
			GroupName: GroupName,
		}
		_, err := gocqhttp.SetGroupName(data)
		if err != nil {
			log.Println("[FastLib] CqSetGroupInfo Error: ", err)
			return false
		}
		return true
	case 2:
		data := models2.SetGroupPortraitData{
			GroupID: GroupID,
			File:    Avatar,
			Cache:   1,
		}
		_, err := gocqhttp.SetGroupPortrait(data)
		if err != nil {
			log.Println("[FastLib] CqSetGroupInfo Error: ", err)
			return false
		}
		return true
	case 3:
		data := models2.SetGroupCardData{
			GroupID: GroupID,
			UserID:  UserId,
			Card:    Card,
		}
		_, err := gocqhttp.SetGroupCard(data)
		if err != nil {
			log.Println("[FastLib] CqSetGroupInfo Error: ", err)
			return false
		}
		return true
	case 4:
		data := models2.SetGroupSpecialTitleData{
			GroupID:      GroupID,
			UserID:       UserId,
			SpecialTitle: SpecialTitle,
		}
		_, err := gocqhttp.SetGroupSpecialTitle(data)
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
	data := models2.SetGroupAdminData{
		GroupID: GroupID,
		UserID:  UserID,
		Enable:  Enable,
	}
	_, err := gocqhttp.SetGroupAdmin(data)
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
	data := models2.SetGroupBanData{
		GroupID:  GroupID,
		UserID:   UserID,
		Duration: Duration,
	}
	_, err := gocqhttp.SetGroupBan(data)
	if err != nil {
		log.Println("[FastLib] CqGroupBan Error: ", err)
		return false
	}
	return true
}

// CqGroupMute mute group.
func CqGroupMute(GroupID int64, UnMute bool) bool {
	data := models2.SetGroupWholeBanData{
		GroupID: GroupID,
		Enable:  UnMute,
	}
	_, err := gocqhttp.SetGroupWholeBan(data)
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
		data := models2.DeleteEssenceMsgData{
			MessageID: MessageID,
		}
		_, err := gocqhttp.DeleteEssenceMsg(data)
		if err != nil {
			log.Println("[FastLib] CqGroupEssenceMsg Error: ", err)
			return false
		}
		return true
	case false:
		data := models2.SetEssenceMsgData{
			MessageID: MessageID,
		}
		_, err := gocqhttp.SetEssenceMsg(data)
		if err != nil {
			log.Println("[FastLib] CqGroupEssenceMsg Error: ", err)
			return false
		}
		return true
	}
	return false
}

func CqGroupSendNotice(GroupID int64, Content string, Image string) bool {
	data := models2.SendGroupNoticeData{
		GroupID: GroupID,
		Content: Content,
		Image:   Image,
	}
	_, err := gocqhttp.SendGroupNotice(data)
	if err != nil {
		log.Println("[FastLib] CqGroupSendNotice Error: ", err)
		return false
	}
	return true
}

// CqGroupKick kick group member.
func CqGroupKick(GroupID int64, UserID int64, RejectAddRequest bool) bool {
	data := models2.SetGroupKickData{
		GroupID: GroupID,
		UserID:  UserID,
		Reject:  RejectAddRequest,
	}
	_, err := gocqhttp.SetGroupKick(data)
	if err != nil {
		log.Println("[FastLib] CqGroupKick Error: ", err)
		return false
	}
	return true
}

func CqLeaveGroup(GroupID int64) bool {
	data := models2.SetGroupLeaveData{
		GroupID:   GroupID,
		IsDismiss: false,
	}
	_, err := gocqhttp.SetGroupLeave(data)
	if err != nil {
		log.Println("[FastLib] CqLeaveGroup Error: ", err)
		return false
	}
	return true
}

func CqDismissGroup(GroupID int64) bool {
	data := models2.SetGroupLeaveData{
		GroupID:   GroupID,
		IsDismiss: true,
	}
	_, err := gocqhttp.SetGroupLeave(data)
	if err != nil {
		log.Println("[FastLib] CqDismissGroup Error: ", err)
		return false
	}
	return true
}
