package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"medods/config"
)

type Mongo struct {
	client *mongo.Client
	DB     *mongo.Database
}

func New(cfg config.MONGO) (Mongo, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.URL))
	if err != nil {
		return Mongo{}, err
	}
	db := client.Database(cfg.Database)
	return Mongo{client: client, DB: db}, nil
}

func (m Mongo) Disconnect() error {
	return m.client.Disconnect(context.Background())
}
