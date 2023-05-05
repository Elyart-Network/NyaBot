package gocqhttp

import (
	"github.com/Elyart-Network/NyaBot/logger"
	"github.com/Elyart-Network/NyaBot/pkg/fastlib/fastcq"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type GetDataFunc struct{}

func (c *GetDataFunc) GetGroupMembers(GroupID int64) []types.GroupMemberInfoObject {
	members, err := fastcq.GetGroupMembers(GroupID)
	if err != nil {
		logger.Warningf("Lua", "GetGroupMembers error", err)
	}
	return members.Data
}

func GetData(L *lua.LState) {
	var GetDataFunc = &GetDataFunc{}
	L.SetGlobal("CqGet", luar.New(L, GetDataFunc))
}
