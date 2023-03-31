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

type GoCqInvoker func(callback cqcall.CallbackFull)
