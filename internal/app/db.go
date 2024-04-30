package app

import (
	"medods/config"
	"medods/pkg/mongo"
)

func openDB(cfg config.MONGO) (mongo.Mongo, error) {
	return mongo.New(cfg)
}
