package runtime

import (
	"github.com/Elyart-Network/NyaBot/internal/config"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/webhook"
	"strings"
)

type CallbackData struct {
	WhCall webhook.Data
	CqCall callback.Full
}

var luaDir = config.Get("plugin.lua_script_dir").(string)

func LoadScript(data CallbackData) {
	GetScripts()
	go func() {
		for _, script := range scripts {
			if strings.HasSuffix(script.FileName, ".lua") && script.Enable {
				LVM(luaDir+"/"+script.FileName, data)
			} else if script.Name == "DEFAULT" {
				continue
			}
		}
	}()
}

func WhLoader(ctx webhook.Data) {
	var data = CallbackData{
		WhCall: ctx,
	}
	LoadScript(data)
}

func CqLoader(ctx callback.Full) {
	var data = CallbackData{
		CqCall: ctx,
	}
	LoadScript(data)
}
