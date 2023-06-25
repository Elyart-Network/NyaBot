package gocqhttp

import (
	"github.com/Elyart-Network/NyaBot/pkg/fastlib/fastcq"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type RequestFunc struct{}

func (c *RequestFunc) FriendReq(Flag string, Approve bool, Remark string) {
	err := fastcq.FriendReq(Flag, Approve, Remark)
	if err != nil {
		log.Warning("[JSVM] FriendReq error: ", err)
	}
}

func (c *RequestFunc) GroupReq(Flag string, Type string, Approve bool, Reason string) {
	err := fastcq.GroupReq(Flag, Type, Approve, Reason)
	if err != nil {
		log.Warning("[JSVM] GroupReq error: ", err)
	}
}

func Request(L *lua.LState) {
	var RequestFunc = &RequestFunc{}
	L.SetGlobal("CqReq", luar.New(L, RequestFunc))
}
