package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"shipping-management/internal/infrastructure/config"
	"time"
)

func NewMongoDatabase(ctx context.Context, cfg *config.Config) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx,
		options.Client().
			SetMaxConnIdleTime(15*time.Second).
			ApplyURI(cfg.MongoDBUri).
			SetAppName(cfg.AppName),
	)

	if err != nil {
		return nil, err
	}

	return client.Database(cfg.MongoDBDatabase), nil
}
