package plugin

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/webhook"
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

type WhPlugin interface {
	CommonInfo
	Receive(callback webhook.Data)
}

type CqPlugin interface {
	CommonInfo
	Message(callback callback.Full)
	Request(callback callback.Full)
	Notice(callback callback.Full)
	MetaEvent(callback callback.Full)
}
