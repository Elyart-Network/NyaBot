package database

import (
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type DBPluginFunc struct{}

func (p *DBPluginFunc) SetConfig(PName string, CKey string, CValue any) {

}

func (p *DBPluginFunc) GetConfig() {

}

func (p *DBPluginFunc) DelConfig() {

}

func DBPlugin(L *lua.LState) {
	var DBPluginFunc = &DBPluginFunc{}
	L.SetGlobal("DBPlug", luar.New(L, DBPluginFunc))
}
