package gocqhttp

import "github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"

type MessageData struct {
	Sender    Sender        `bson:"sender"`
	Content   string        `bson:"content"`
	MessageID int32         `bson:"message_id"`
	TimeStamp int64         `bson:"timestamp"`
	Raw       callback.Full `bson:"raw"`
}

type Sender struct {
	UserIdentity string `bson:"user_identity"`
	UserCard     string `bson:"user_card"`
	UserNick     string `bson:"user_nick"`
}
