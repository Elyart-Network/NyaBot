package runtime

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/webhook"
	"log"
	"strings"
)

type CallbackData struct {
	WhCall webhook.Data
	CqCall callback.Full
}

func LoadScript(data CallbackData) {
	GetScripts()
	go func() {
		for _, script := range scripts {
			if strings.HasSuffix(script.FileName, ".lua") && script.Enable {
				LVM("scripts/"+script.FileName, data)
			} else if script.Name == "DEFAULT" {
				continue
			} else {
				log.Println("[Lua] Script " + script.Name + " is disabled!")
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
