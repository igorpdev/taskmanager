package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect(uri string, minPoolSize, maxPoolSize uint64, timeout time.Duration) error {
	clientOptions := options.Client().
		ApplyURI(uri).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize).
		SetConnectTimeout(timeout)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	Client = client
	return nil
}

func Disconnect() error {
	if Client == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return Client.Disconnect(ctx)
}

func GetCollection(collectionName string) *mongo.Collection {
	if Client == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return Client.Database("taskmanager").Collection(collectionName)
}
