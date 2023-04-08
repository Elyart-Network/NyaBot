package gocqhttp

import (
	"github.com/Elyart-Network/NyaBot/pkg/fastlib/fastcq"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"log"
)

type MessageFunc struct{}

func (c *MessageFunc) SendMsg(Message string, Id int64, IsGroup bool) {
	_, err := fastcq.SendMsg(Message, Id, IsGroup)
	if err != nil {
		log.Println("[Lua] SendMsg error: ", err)
	}
}

func (c *MessageFunc) Reply(Message string, Id int64, IsGroup bool, To string) {
	ReplyCodeData := types.ReplyData{ID: To}
	ReplyCode := gocqhttp.Reply(ReplyCodeData)
	_, err := fastcq.SendMsg(ReplyCode+Message, Id, IsGroup)
	if err != nil {
		log.Println("[Lua] SendMsgWithAt error: ", err)
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
