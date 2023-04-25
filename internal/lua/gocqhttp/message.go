package gocqhttp

import (
	"github.com/Elyart-Network/NyaBot/pkg/fastlib/fastcq"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqcode"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"log"
	"strconv"
)

type MessageFunc struct{}

func (c *MessageFunc) SendMsg(Message string, Id int64, IsGroup bool) {
	_, err := fastcq.SendMsg(Message, Id, IsGroup)
	if err != nil {
		log.Println("[Lua] SendMsg error: ", err)
	}
}

func (c *MessageFunc) Reply(Message string, Id int64, IsGroup bool, To int) {
	ReplyCodeData := types.ReplyData{ID: strconv.Itoa(To)}
	ReplyCode := cqcode.Reply(ReplyCodeData)
	_, err := fastcq.SendMsg(ReplyCode+Message, Id, IsGroup)
	if err != nil {
		log.Println("[Lua] SendMsgWithAt error: ", err)
	}
}

func (c *MessageFunc) SendPic(Url string, Type string, Id int64, IsGroup bool) {
	PicData := types.ImageData{
		File: "pic." + Type,
		Url:  Url,
	}
	_, err := fastcq.SendMsg(cqcode.Image(PicData), Id, IsGroup)
	if err != nil {
		log.Println("[Lua] SendPic error: ", err)
	}
}

var ForwardMessages []interface{}

func (c *MessageFunc) SetIdForward(MessageID string) {
	data := fastcq.GenIdForward(MessageID)
	ForwardMessages = append(ForwardMessages, data)
}

func (c *MessageFunc) SetCustomForward(Name string, Id string, Content string) {
	data := fastcq.GenCustomForward(Name, Id, Content)
	ForwardMessages = append(ForwardMessages, data)
}

func (c *MessageFunc) SendForwardMsg(Id int64, IsGroup bool) {
	_, err := fastcq.SendForwardMsg(ForwardMessages, Id, IsGroup)
	if err != nil {
		log.Println("[Lua] SendForwardMsg error: ", err)
	}
	ForwardMessages = nil
}

func (c *MessageFunc) DeleteMsg(MessageID int32) {
	err := fastcq.DeleteMsg(MessageID)
	if err != nil {
		log.Println("[Lua] DeleteMsg error: ", err)
	}
}

func Message(L *lua.LState) {
	var MessageFunc = &MessageFunc{}
	L.SetGlobal("CqMsg", luar.New(L, MessageFunc))
}
