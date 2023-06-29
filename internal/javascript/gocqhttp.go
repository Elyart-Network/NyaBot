package javascript

import (
	"context"
	"github.com/Elyart-Network/NyaBot/internal/javascript/runtime"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/plugin"
)

// Plugin a struct for all functions below.
type Plugin struct{}

// Info set plugin info, `Name` has to be unique!
func (p *Plugin) Info() plugin.InfoStruct {
	return plugin.InfoStruct{
		Name:        "cq_internal_javascript",
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
func (p *Plugin) Message(ctx context.Context, callback callback.Full) {
	runtime.CqLoader(callback)
}

// Request process request event from callback. (required)
func (p *Plugin) Request(ctx context.Context, callback callback.Full) {
	runtime.CqLoader(callback)
}

// Notice process notice event from callback. (required)
func (p *Plugin) Notice(ctx context.Context, callback callback.Full) {
	runtime.CqLoader(callback)
}

// MetaEvent process meta event from callback. (required)
func (p *Plugin) MetaEvent(ctx context.Context, callback callback.Full) {
	runtime.CqLoader(callback)
}

// init register plugin and depends to plugin manager (frame).
func init() {
	plugin.CqRegister(&Plugin{})
}
