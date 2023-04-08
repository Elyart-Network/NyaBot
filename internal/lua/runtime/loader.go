package runtime

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"strings"
)

type CallbackData struct {
	CqCall callback.Full
}

func CqLoader(ctx callback.Full) {
	var data = CallbackData{
		CqCall: ctx,
	}
	scripts, _ := GetScripts()
	go func() {
		for _, script := range scripts {
			if strings.HasSuffix(script, ".lua") {
				LVM("scripts/"+script, data)
			}
		}
	}()
}
