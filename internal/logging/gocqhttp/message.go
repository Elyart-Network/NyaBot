package gocqhttp

import (
	"context"
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/db"
	"github.com/Elyart-Network/NyaBot/logger"
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
	// Connect to Cache and insert message
	cacheNum := config.Get("logging.cache_num").(int)
	externalCache := config.Get("cache.external").(bool)
	if cacheNum > 0 {
		switch externalCache {
		case true:
			content, _ := json.Marshal(msg.Raw)
			r := db.Redis(0)
			clen := r.LLen(con, collection)
			if clen.Val() >= int64(cacheNum) {
				r.LPopCount(con, collection, 1)
			}
			r.RPush(con, collection, content)
			err := r.Close()
			if err != nil {
				logger.Warningf("Redis", "Failed to disconnect from Redis", err)
			}
		case false:
			// TODO
		}
	}
	// Connect to MongoDB and insert message
	enableMDB := config.Get("logging.external").(bool)
	if enableMDB {
		m := db.MongoDB()
		m.Insert(con, collection, msg)
		err := m.Disconnect(con)
		if err != nil {
			logger.Warningf("MongoDB", "Failed to disconnect from MongoDB", err)
		}
	}
}
