package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	Session Session
}

func NewRepositories(client *mongo.Database) Repositories {
	return Repositories{
		Session: NewSession(client),
	}
}
