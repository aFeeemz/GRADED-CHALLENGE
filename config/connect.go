package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectionDatabase(ctx context.Context) (*mongo.Collection, error) {
	mongiURI := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongiURI))
	if err != nil {
		return nil, err
	}

	collection := client.Database("rentfield_gc").Collection("transactions")
	return collection, nil
}
