package plugin

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqcall"
)

type InfoStruct struct {
	Name        string
	Version     string
	Author      string
	Description string
	License     string
	Homepage    string
	Repository  string
	Type        string
}

type CommonInfo interface {
	Info() InfoStruct
}

var cqCallback cqcall.CallbackFull

type CqPlugin interface {
	CommonInfo
	Message(callback cqcall.CallbackFull)
	Request(callback cqcall.CallbackFull)
	Notice(callback cqcall.CallbackFull)
	MetaEvent(callback cqcall.CallbackFull)
}
