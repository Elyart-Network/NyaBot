package lua

import (
	"github.com/Elyart-Network/NyaBot/internal/lua/runtime"
	"github.com/Elyart-Network/NyaBot/pkg/plugin"
	"github.com/Elyart-Network/NyaBot/pkg/webhook"
)

// WhPlugin a struct for all functions below.
type WhPlugin struct{}

// Info set plugin info, `Name` has to be unique!
func (p *WhPlugin) Info() plugin.InfoStruct {
	return plugin.InfoStruct{
		Name:        "wh_example",
		Version:     "",
		Author:      "",
		Description: "",
		License:     "",
		Homepage:    "",
		Repository:  "",
		Type:        "Webhook",
	}
}

// Receive process webhook data from callback. (required)
func (p *WhPlugin) Receive(callback webhook.Data) {
	runtime.WhLoader(callback)
}

// init register plugin and depends to plugin manager (frame).
func init() {
	plugin.WhRegister(&WhPlugin{})
}
