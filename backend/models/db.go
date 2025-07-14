package models

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		slog.Error("Failed to connect to MongoDB", "error", err)
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		slog.Error("Failed to ping MongoDB", "error", err)
		return err
	}

	MongoClient = client
	slog.Info("Connected to MongoDB", "uri", uri)
	return nil
}

func GetCollection(dbName, collName string) *mongo.Collection {
	return MongoClient.Database(dbName).Collection(collName)
}
