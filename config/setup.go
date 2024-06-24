package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoConn *mongo.Client
	ctx       context.Context
)

func InitMongoDBConnection(mongoURI, dbName string) *mongo.Client {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to db")

	mongoConn = client
	return mongoConn
}

func GetMongoClient() *mongo.Client {
	return mongoConn
}

func GetContext() context.Context {
	return ctx
}
