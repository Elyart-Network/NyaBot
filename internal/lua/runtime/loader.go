package runtime

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"log"
	"strings"
)

type CallbackData struct {
	CqCall callback.Full
}

func CqLoader(ctx callback.Full) {
	var data = CallbackData{
		CqCall: ctx,
	}
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
