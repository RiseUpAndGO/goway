package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

const (
	connectTimeout         = 10 * time.Second
	serverSelectionTimeout = 5 * time.Second
)

func Connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017",
		os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		os.Getenv("MONGO_CONTAINER_NAME"))

	clientOpts := options.Client().
		ApplyURI(uri).
		SetConnectTimeout(connectTimeout).
		SetServerSelectionTimeout(serverSelectionTimeout)

	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	return client, nil
}
