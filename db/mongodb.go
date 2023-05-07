package db

import (
	"context"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	db *mongo.Client
}

func MongoDB() *MongoClient {
	// Connect to MongoDB
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
	return &MongoClient{db: conn}
}

func (m *MongoClient) Insert(ctx context.Context, collection string, content any) {
	if m.db == nil {
		return
	}
	dbn := config.Get("logging.mongo_db").(string)
	conn := m.db.Database(dbn).Collection(collection)
	_, err := conn.InsertOne(ctx, content)
	if err != nil {
		logger.Errorf("MongoDB", "Error while inserting message into MongoDB", err)
	}
}

func (m *MongoClient) Disconnect(ctx context.Context) error {
	return m.db.Disconnect(ctx)
}
