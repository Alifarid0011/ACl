package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

var (
	mongoClient *mongo.Client
	mongoOnce   sync.Once
)

// InitMongoClient initializes and returns a singleton MongoDB client.
// It ensures connection is alive and reconnects if needed.
func InitMongoClient(uri string) *mongo.Client {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)

		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
		}
		mongoClient = client
	})
	return mongoClient
}
