package runtime

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/webhook"
)

type CallbackData struct {
	WhCall webhook.Data
	CqCall callback.Full
}
