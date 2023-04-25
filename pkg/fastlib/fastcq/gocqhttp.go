package fastcq

import (
	"errors"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/common"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/friend"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/group"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/system"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
	"log"
)

// SendMsg send message to group or user.
// Message: Message to send.
// ID: Group ID or User ID.
// IsGroup: true for group, false for user.
func SendMsg(Message string, Id int64, IsGroup bool) (int32, error) {
	switch IsGroup {
	case true:
		data := types.SendGroupMsgData{
			GroupID:    Id,
			Message:    Message,
			AutoEscape: false,
		}
		msg, err := common.SendGroupMsg(data)
		if err != nil {
			return 0, err
		}
		return msg.Data.MessageID, nil
	case false:
		data := types.SendPrivateMsgData{
			UserID:     Id,
			Message:    Message,
			AutoEscape: false,
		}
		msg, err := common.SendPrivateMsg(data)
		if err != nil {
			return 0, err
		}
		return msg.Data.MessageID, nil
	}
	return 0, errors.New("unknown error")
}

// GetMsg get message from message ID.
func GetMsg(MessageID int32) (types.GetMsgResp, error) {
	data := types.GetMsgData{
		MessageID: MessageID,
	}
	msg, err := common.GetMsg(data)
	if err != nil {
		return types.GetMsgResp{}, err
	}
	return msg, nil
}

// GetForwardMsg get forward message from message ID.
func GetForwardMsg(MessageID string) (types.GetForwardMsgResp, error) {
	data := types.GetForwardMsgData{
		MessageID: MessageID,
	}
	msg, err := common.GetForwardMsg(data)
	if err != nil {
		log.Println("[FastLib] CqGetForwardMsg Error: ", err)
		return types.GetForwardMsgResp{}, err
	}
	return msg, nil
}

// GenIdForward generate forward node from message ID.
func GenIdForward(MessageID string) types.ForwardIdNode {
	var data = types.ForwardIdData{
		Id: MessageID,
	}
	var forward = types.ForwardIdNode{
		Type: "node",
		Data: data,
	}
	return forward
}

// GenCustomForward generate forward node from custom data.
// Name: Name of the sender.
// ID: QQ ID of the sender.
// Content: Content of the message.
func GenCustomForward(Name string, Id string, Content string) types.ForwardCustomNode {
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

// SendForwardMsg send forward message to group or user.
// Message: Message to send.
// ID: Group ID or User ID.
// IsGroup: true for group, false for user.
func SendForwardMsg(Messages interface{}, Id int64, IsGroup bool) (int64, error) {
	switch IsGroup {
	case true:
		data := types.SendGroupForwardMsgData{
			GroupID:  Id,
			Messages: Messages,
		}
		msg, err := common.SendGroupForwardMsg(data)
		if err != nil {
			return 0, err
		}
		return msg.Data.MessageID, nil
	case false:
		data := types.SendPrivateForwardMsgData{
			UserID:   Id,
			Messages: Messages,
		}
		msg, err := common.SendPrivateForwardMsg(data)
		if err != nil {
			return 0, err
		}
		return msg.Data.MessageID, nil
	}
	return 0, errors.New("unknown error")
}

// DeleteMsg delete message from message ID.
func DeleteMsg(MessageID int32) error {
	data := types.DeleteMsgData{
		MessageID: MessageID,
	}
	_, err := common.DeleteMsg(data)
	if err != nil {
		return err
	}
	return nil
}

// GetImage get image from cached file name.
func GetImage(FileName string) (types.GetImageResp, error) {
	data := types.GetImageData{
		File: FileName,
	}
	msg, err := common.GetImage(data)
	if err != nil {
		return types.GetImageResp{}, err
	}
	return msg, nil
}

// GetRecord get record cached from file name.
func GetRecord(FileName string) (types.GetRecordResp, error) {
	data := types.GetRecordData{
		File: FileName,
	}
	msg, err := common.GetRecord(data)
	if err != nil {
		return types.GetRecordResp{}, err
	}
	return msg, nil
}

// FriendReq process friend request.
// Flag: Request flag, get from request event.
// Approve: true for approve, false for reject.
// Remark: Remark for friend.
func FriendReq(Flag string, Approve bool, Remark string) error {
	data := types.SetFriendAddRequestData{
		Flag:    Flag,
		Approve: Approve,
		Remark:  Remark,
	}
	_, err := system.SetFriendAddRequest(data)
	if err != nil {
		return err
	}
	return nil
}

// GroupReq process group request.
// Flag: Request flag, get from request event.
// Type: Request type, get from request event.
// Approve: true for approve, false for reject.
// RejectReason: Reject reason.
func GroupReq(Flag string, Type string, Approve bool, RejectReason string) error {
	data := types.SetGroupAddRequestData{
		Flag:    Flag,
		SubType: Type,
		Approve: Approve,
		Reason:  RejectReason,
	}
	_, err := system.SetGroupAddRequest(data)
	if err != nil {
		return err
	}
	return nil
}

// GetStrangerInfo get stranger info.
func GetStrangerInfo(UserID int64) (types.GetStrangerInfoResp, error) {
	data := types.GetStrangerInfoData{
		UserID:  UserID,
		NoCache: false,
	}
	msg, err := friend.GetStrangerInfo(data)
	if err != nil {
		return types.GetStrangerInfoResp{}, err
	}
	return msg, nil
}

// GetGroupMemberInfo get group member info.
func GetGroupMemberInfo(GroupID int64, UserID int64) (types.GetGroupMemberInfoResp, error) {
	data := types.GetGroupMemberInfoData{
		GroupID: GroupID,
		UserID:  UserID,
		NoCache: false,
	}
	msg, err := group.GetGroupMemberInfo(data)
	if err != nil {
		return types.GetGroupMemberInfoResp{}, err
	}
	return msg, nil
}

// GetGroupMembers get group member list.
func GetGroupMembers(GroupID int64) (types.GetGroupMemberListResp, error) {
	data := types.GetGroupMemberListData{
		GroupID: GroupID,
		NoCache: false,
	}
	msg, err := group.GetGroupMemberList(data)
	if err != nil {
		return types.GetGroupMemberListResp{}, err
	}
	return msg, nil
}

// GetFriends get friend list.
func GetFriends() (types.GetFriendListResp, error) {
	msg, err := friend.GetFriendList()
	if err != nil {
		return types.GetFriendListResp{}, err
	}
	return msg, nil
}

// GetGroups get group list.
func GetGroups() (types.GetGroupListResp, error) {
	data := types.GetGroupListData{
		NoCache: false,
	}
	msg, err := group.GetGroupList(data)
	if err != nil {
		return types.GetGroupListResp{}, err
	}
	return msg, nil
}

// GetGroupInfo get group info.
func GetGroupInfo(GroupID int64) (types.GetGroupInfoResp, error) {
	data := types.GetGroupInfoData{
		GroupID: GroupID,
		NoCache: false,
	}
	msg, err := group.GetGroupInfo(data)
	if err != nil {
		return types.GetGroupInfoResp{}, err
	}
	return msg, nil
}

// SetGroupInfo set group info.
// Type 1: GroupID, GroupName set group name.
// Type 2: GroupID, Avatar(path) set group avatar.
// Type 3: GroupID, UserId, Card set group card.
// Type 4: GroupID, UserId, SpecialTitle set group member special title.
func SetGroupInfo(GroupID int64, GroupName string, Avatar string, UserId int64, Card string, SpecialTitle string, Type int) error {
	switch Type {
	case 1:
		data := types.SetGroupNameData{
			GroupID:   GroupID,
			GroupName: GroupName,
		}
		_, err := group.SetGroupName(data)
		if err != nil {
			return err
		}
		return nil
	case 2:
		data := types.SetGroupPortraitData{
			GroupID: GroupID,
			File:    Avatar,
			Cache:   1,
		}
		_, err := group.SetGroupPortrait(data)
		if err != nil {
			return err
		}
		return nil
	case 3:
		data := types.SetGroupCardData{
			GroupID: GroupID,
			UserID:  UserId,
			Card:    Card,
		}
		_, err := group.SetGroupCard(data)
		if err != nil {
			return err
		}
		return nil
	case 4:
		data := types.SetGroupSpecialTitleData{
			GroupID:      GroupID,
			UserID:       UserId,
			SpecialTitle: SpecialTitle,
		}
		_, err := group.SetGroupSpecialTitle(data)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("type error")
}

// SetGroupAdmin set group admin.
func SetGroupAdmin(GroupID int64, UserID int64, Enable bool) error {
	data := types.SetGroupAdminData{
		GroupID: GroupID,
		UserID:  UserID,
		Enable:  Enable,
	}
	_, err := group.SetGroupAdmin(data)
	if err != nil {
		return err
	}
	return nil
}

// GroupBan ban group member.
// BanAll true: ban all member.
// DeBan true: unban.
// Duration only work when BanAll is false.
func GroupBan(GroupID int64, UserID int64, Duration uint32, DeBan bool) error {
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
		return err
	}
	return nil
}

// GroupMute mute group.
func GroupMute(GroupID int64, UnMute bool) error {
	data := types.SetGroupWholeBanData{
		GroupID: GroupID,
		Enable:  UnMute,
	}
	_, err := group.SetGroupWholeBan(data)
	if err != nil {
		return err
	}
	return nil
}

// GroupEssenceMsg set group essence message.
func GroupEssenceMsg(MessageID int32, Remove bool) error {
	switch Remove {
	case true:
		data := types.DeleteEssenceMsgData{
			MessageID: MessageID,
		}
		_, err := group.DeleteEssenceMsg(data)
		if err != nil {
			return err
		}
		return nil
	case false:
		data := types.SetEssenceMsgData{
			MessageID: MessageID,
		}
		_, err := group.SetEssenceMsg(data)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("remove error")
}

// GroupSendNotice send group notice.
func GroupSendNotice(GroupID int64, Content string, Image string) error {
	data := types.SendGroupNoticeData{
		GroupID: GroupID,
		Content: Content,
		Image:   Image,
	}
	_, err := group.SendGroupNotice(data)
	if err != nil {
		return err
	}
	return nil
}

// GroupKick kick group member.
func GroupKick(GroupID int64, UserID int64, RejectAddRequest bool) error {
	data := types.SetGroupKickData{
		GroupID: GroupID,
		UserID:  UserID,
		Reject:  RejectAddRequest,
	}
	_, err := group.SetGroupKick(data)
	if err != nil {
		return err
	}
	return nil
}

// LeaveGroup leave group.
func LeaveGroup(GroupID int64) error {
	data := types.SetGroupLeaveData{
		GroupID:   GroupID,
		IsDismiss: false,
	}
	_, err := group.SetGroupLeave(data)
	if err != nil {
		return err
	}
	return nil
}

// DismissGroup dismiss group.
func DismissGroup(GroupID int64) error {
	data := types.SetGroupLeaveData{
		GroupID:   GroupID,
		IsDismiss: true,
	}
	_, err := group.SetGroupLeave(data)
	if err != nil {
		return err
	}
	return nil
}

// SendGroupSign send group sign.
func SendGroupSign(GroupID int64) error {
	data := types.SendGroupSignData{
		GroupID: GroupID,
	}
	_, err := group.SendGroupSign(data)
	if err != nil {
		return err
	}
	return nil
}

// SetEssenceMsg set group essence message.
func SetEssenceMsg(MessageID int32, Remove bool) error {
	switch Remove {
	case true:
		data := types.DeleteEssenceMsgData{
			MessageID: MessageID,
		}
		_, err := group.DeleteEssenceMsg(data)
		if err != nil {
			return err
		}
		return nil
	case false:
		data := types.SetEssenceMsgData{
			MessageID: MessageID,
		}
		_, err := group.SetEssenceMsg(data)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("remove error")
}
