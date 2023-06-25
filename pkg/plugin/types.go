package plugin

import (
	"context"
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
	Message(ctx context.Context, callback callback.Full)
	Request(ctx context.Context, callback callback.Full)
	Notice(ctx context.Context, callback callback.Full)
	MetaEvent(ctx context.Context, callback callback.Full)
}
