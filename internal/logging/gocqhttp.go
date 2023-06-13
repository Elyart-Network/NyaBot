package lua

import (
	"github.com/Elyart-Network/NyaBot/internal/logging/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/plugin"
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
	gocqhttp.Message(ctx)
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
