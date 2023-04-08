package gocqhttp

import (
	"github.com/Elyart-Network/NyaBot/pkg/fastlib/fastcq"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"log"
)

type ActionFunc struct{}

func (c *ActionFunc) SetGroupInfo(GroupID int64, GroupName string, Avatar string, UserId int64, Card string, SpecialTitle string, Type int) {
	err := fastcq.SetGroupInfo(GroupID, GroupName, Avatar, UserId, Card, SpecialTitle, Type)
	if err != nil {
		log.Println("[Lua] SetGroupInfo error: ", err)
	}
}

func (c *ActionFunc) SetGroupAdmin(GroupID int64, UserId int64, Enable bool) {
	err := fastcq.SetGroupAdmin(GroupID, UserId, Enable)
	if err != nil {
		log.Println("[Lua] SetGroupAdmin error: ", err)
	}
}

func (c *ActionFunc) GroupBan(GroupID int64, UserID int64, Duration uint32, DeBan bool) {
	err := fastcq.GroupBan(GroupID, UserID, Duration, DeBan)
	if err != nil {
		log.Println("[Lua] GroupBan error: ", err)
	}
}

func (c *ActionFunc) GroupMute(GroupID int64, UnMute bool) {
	err := fastcq.GroupMute(GroupID, UnMute)
	if err != nil {
		log.Println("[Lua] GroupMute error: ", err)
	}
}

func (c *ActionFunc) GroupEssenceMsg(MessageID int32, Remove bool) {
	err := fastcq.GroupEssenceMsg(MessageID, Remove)
	if err != nil {
		log.Println("[Lua] GroupEssenceMsg error: ", err)
	}
}

func (c *ActionFunc) GroupSendNotice(GroupID int64, Content string, Image string) {
	err := fastcq.GroupSendNotice(GroupID, Content, Image)
	if err != nil {
		log.Println("[Lua] GroupSendNotice error: ", err)
	}
}

func (c *ActionFunc) GroupKick(GroupID int64, UserID int64, RejectAddRequest bool) {
	err := fastcq.GroupKick(GroupID, UserID, RejectAddRequest)
	if err != nil {
		log.Println("[Lua] GroupKick error: ", err)
	}
}

func (c *ActionFunc) LeaveGroup(GroupID int64) {
	err := fastcq.LeaveGroup(GroupID)
	if err != nil {
		log.Println("[Lua] LeaveGroup error: ", err)
	}
}

func (c *ActionFunc) DismissGroup(GroupID int64) {
	err := fastcq.DismissGroup(GroupID)
	if err != nil {
		log.Println("[Lua] DeleteMsg error: ", err)
	}
}

func Action(L *lua.LState) {
	var ActionFunc = &ActionFunc{}
	L.SetGlobal("CqAct", luar.New(L, ActionFunc))
}
