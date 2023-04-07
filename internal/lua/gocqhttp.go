package lua

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/plugin"
	"log"
)

// Plugin a struct for all functions below.
type Plugin struct{}

// Info set plugin info, `Name` has to be unique!
func (p *Plugin) Info() plugin.InfoStruct {
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
func (p *Plugin) Message(callback callback.Full) {
	log.Println("Message")
}

// Request process request event from callback. (required)
func (p *Plugin) Request(callback callback.Full) {
	log.Println("Request")
}

// Notice process notice event from callback. (required)
func (p *Plugin) Notice(callback callback.Full) {}

// MetaEvent process meta event from callback. (required)
func (p *Plugin) MetaEvent(callback callback.Full) {}

// init register plugin and depends to plugin manager (frame).
func init() {
	plugin.CqRegister(&Plugin{})
}
