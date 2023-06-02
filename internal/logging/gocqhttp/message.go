package gocqhttp

import (
	"context"
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/data/actions"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"strconv"
)

func Message(ctx callback.Full) {
	// Create message object
	sender := Sender{
		UserIdentity: strconv.FormatInt(ctx.Sender.UserID, 10),
		UserCard:     ctx.Sender.Card,
		UserNick:     ctx.Sender.Nickname,
	}
	msg := MessageData{
		Sender:    sender,
		Content:   ctx.MessageData,
		MessageID: ctx.MessageID,
		TimeStamp: ctx.Time,
		Raw:       ctx,
	}
	var collection string
	switch msg.Raw.MessageType {
	case "private":
		collection = "cq_messages." + msg.Raw.MessageType + "." + strconv.FormatInt(msg.Raw.UserID, 10)
	case "group":
		collection = "cq_messages." + msg.Raw.MessageType + "." + strconv.FormatInt(msg.Raw.GroupID, 10)
	default:
		collection = "cq_messages." + msg.Raw.MessageType
	}
	con := context.Background()

	loggingComponent := actions.Logging{}
	// Encode raw message to JSON
	content, _ := json.Marshal(msg.Raw)
	loggingComponent.Cache(con, collection, content)

	// Connect to MongoDB and insert message
	loggingComponent.Mongo(con, collection, content)
}
