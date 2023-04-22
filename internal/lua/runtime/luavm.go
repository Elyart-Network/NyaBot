package runtime

import (
	"github.com/Elyart-Network/NyaBot/internal/lua/common"
	"github.com/Elyart-Network/NyaBot/internal/lua/gocqhttp"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"log"
)

func LVM(path string, data CallbackData) {
	L := lua.NewState()
	defer L.Close()
	// send callback data to the lua script
	L.SetGlobal("Callback", luar.New(L, data))
	// register modules to the table
	common.HttpReq(L)
	gocqhttp.Message(L)
	gocqhttp.Request(L)
	gocqhttp.Action(L)
	gocqhttp.GetData(L)
	// load and run the script
	if err := L.DoFile(path); err != nil {
		log.Println("Lua script error: ", err)
	}
}
