// Package Example is an example plugin.
package example

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqcall"
	"github.com/Elyart-Network/NyaBot/pkg/plugin"
	"log"
)

// Plugin a struct for all functions below.
type Plugin struct{}

// Info set plugin info, `Name` has to be unique!
func (p *Plugin) Info() plugin.InfoStruct {
	return plugin.InfoStruct{
		Name:        "CqExample",
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
func (p *Plugin) Message(callback cqcall.CallbackFull) {
	log.Println("Message")
}

// Request process request event from callback. (required)
func (p *Plugin) Request(callback cqcall.CallbackFull) {
	log.Println("Request")
}

// Notice process notice event from callback. (required)
func (p *Plugin) Notice(callback cqcall.CallbackFull) {
	log.Println("Notice")
}

// MetaEvent process meta event from callback. (required)
func (p *Plugin) MetaEvent(callback cqcall.CallbackFull) {
	log.Println("MetaEvent")
}

// init register plugin and depends to plugin manager (frame).
func init() {
	plugin.CqRegister(&Plugin{})
}