package handlers

import (
	"context"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/db"
	"github.com/Elyart-Network/NyaBot/internal/logging/types"
	"github.com/Elyart-Network/NyaBot/logger"
)

func Message(msg types.Message) {
	mdb := db.MongoDB()
	if mdb == nil {
		return
	}
	ctx := context.Background()
	defer func() {
		if err := mdb.Disconnect(ctx); err != nil {
			logger.Errorf("MongoDB", "Error while disconnecting from MongoDB", err)
		}
	}()
	dbn := config.Get("logging.mongo_db").(string)
	conn := mdb.Database(dbn).Collection("messages")
	_, err := conn.InsertOne(ctx, msg)
	if err != nil {
		logger.Errorf("MongoDB", "Error while inserting message into MongoDB", err)
	}
}
