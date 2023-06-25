package lua

import (
	"context"
	"github.com/Elyart-Network/NyaBot/internal/lua/runtime"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/plugin"
)

// CqPlugin a struct for all functions below.
type CqPlugin struct{}

// Info set plugin info, `Name` has to be unique!
func (p *CqPlugin) Info() plugin.InfoStruct {
	return plugin.InfoStruct{
		Name:        "cq_internal_lua",
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
func (p *CqPlugin) Message(ctx context.Context, callback callback.Full) {
	runtime.CqLoader(callback)
}

// Request process request event from callback. (required)
func (p *CqPlugin) Request(ctx context.Context, callback callback.Full) {
	runtime.CqLoader(callback)
}

// Notice process notice event from callback. (required)
func (p *CqPlugin) Notice(ctx context.Context, callback callback.Full) {
	runtime.CqLoader(callback)
}

// MetaEvent process meta event from callback. (required)
func (p *CqPlugin) MetaEvent(ctx context.Context, callback callback.Full) {}

// init register plugin and depends to plugin manager (frame).
func init() {
	plugin.CqRegister(&CqPlugin{})
}
