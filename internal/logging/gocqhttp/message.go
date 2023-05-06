package gocqhttp

import (
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/db/actions"
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
	switch msg.Raw.SubType {
	case "private":
		collection = "cq_messages." + msg.Raw.SubType + "." + strconv.FormatInt(msg.Raw.UserID, 10)
	case "group":
		collection = "cq_messages." + msg.Raw.SubType + "." + strconv.FormatInt(msg.Raw.GroupID, 10)
	default:
		collection = "cq_messages." + msg.Raw.SubType
	}
	// Connect to Cache and insert message
	cacheNum := config.Get("logging.cache_num").(int)
	if cacheNum > 0 {
		actions.CacheInsert(collection, msg)
	}
	// Connect to MongoDB and insert message
	enableMDB := config.Get("logging.external").(bool)
	if enableMDB {
		actions.MongoInsert(collection, msg)
	}
}
