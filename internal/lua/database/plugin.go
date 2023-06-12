package database

import (
	"github.com/Elyart-Network/NyaBot/pkg/database"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type DBPluginFunc struct{}

func (p *DBPluginFunc) SetConfig(PName string, CKey string, CValue any) {
	database.SetConfig(PName, "Lua", CKey, CValue)
}

func (p *DBPluginFunc) GetConfig(PName string) map[string]any {
	return database.GetConfig(PName, "Lua")
}

func (p *DBPluginFunc) DelConfig(PName string, CKey string) {
	database.DelConfig(PName, "Lua", CKey)
}

func (p *DBPluginFunc) DelPlugin(PName string) {
	database.DelPlugin(PName, "Lua")
}

func DBPlugin(L *lua.LState) {
	var DBPluginFunc = &DBPluginFunc{}
	L.SetGlobal("DBPlug", luar.New(L, DBPluginFunc))
}
