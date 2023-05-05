package lua

import (
	"github.com/Elyart-Network/NyaBot/internal/logging/handlers"
	"github.com/Elyart-Network/NyaBot/internal/logging/types"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/plugin"
	"strconv"
)

// CqPlugin a struct for all functions below.
type CqPlugin struct{}

// Info set plugin info, `Name` has to be unique!
func (p *CqPlugin) Info() plugin.InfoStruct {
	return plugin.InfoStruct{
		Name:        "cq_internal_logging",
		Version:     "",
		Author:      "",
		Description: "",
		License:     "",
		Homepage:    "",
		Repository:  "",
		Type:        "GoCqHttp",
	}
}

// Message process message event from callback. (required)
func (p *CqPlugin) Message(ctx callback.Full) {
	sender := types.Sender{
		UserIdentity: strconv.FormatInt(ctx.Sender.UserID, 10),
		UserName:     ctx.Sender.Card,
		UserNick:     ctx.Sender.Nickname,
	}
	message := types.Message{
		Sender:    sender,
		Content:   ctx.MessageData,
		TimeStamp: ctx.Time,
		RawMsg:    ctx.RawMessage,
		Raw:       ctx,
	}
	handlers.Message(message)
}

// Request process request event from callback. (required)
func (p *CqPlugin) Request(ctx callback.Full) {}

// Notice process notice event from callback. (required)
func (p *CqPlugin) Notice(ctx callback.Full) {}

// MetaEvent process meta event from callback. (required)
func (p *CqPlugin) MetaEvent(ctx callback.Full) {}

// init register plugin and depends to plugin manager (frame).
func init() {
	plugin.CqRegister(&CqPlugin{})
}
