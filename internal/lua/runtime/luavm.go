package runtime

import (
	"github.com/Elyart-Network/NyaBot/internal/lua/common"
	"github.com/Elyart-Network/NyaBot/internal/lua/gocqhttp"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func LVM(path string, data CallbackData) {
	L := lua.NewState()
	defer L.Close()
	// send callback data to the lua script
	L.SetGlobal("Callback", luar.New(L, data))
	// register modules to the table
	common.HttpReq(L)
	common.System(L)
	gocqhttp.Message(L)
	gocqhttp.Request(L)
	gocqhttp.Action(L)
	gocqhttp.GetData(L)
	// load and run the script
	if err := L.DoFile(path); err != nil {
		log.Errorf("[Lua] Error running lua script: %v", err)
	}
}
