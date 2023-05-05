package db

import (
	"context"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDB() *mongo.Client {
	enable := config.Get("logging.external").(bool)
	if !enable {
		return nil
	}
	dsn := config.Get("logging.mongo_uri").(string)
	ctx := context.Background()
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		logger.Errorf("MongoDB", "Failed to connect to MongoDB", err)
	}
	err = conn.Ping(ctx, nil)
	if err != nil {
		logger.Errorf("MongoDB", "Failed to ping MongoDB", err)
	}
	return conn
}
