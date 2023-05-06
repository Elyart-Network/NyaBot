package actions

import (
	"context"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/db/drivers"
	"github.com/Elyart-Network/NyaBot/logger"
)

func MongoInsert(collection string, content any) {
	mdb := drivers.MongoDB()
	if mdb == nil {
		return
	}
	dbctx := context.Background()
	defer func() {
		if err := mdb.Disconnect(dbctx); err != nil {
			logger.Errorf("MongoDB", "Error while disconnecting from MongoDB", err)
		}
	}()
	dbn := config.Get("logging.mongo_db").(string)
	conn := mdb.Database(dbn).Collection(collection)
	_, err := conn.InsertOne(dbctx, content)
	if err != nil {
		logger.Errorf("MongoDB", "Error while inserting message into MongoDB", err)
	}
}
