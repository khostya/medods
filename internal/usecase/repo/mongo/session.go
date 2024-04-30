package mongo

import (
	"context"
	"errors"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"medods/internal/entity"
	"medods/internal/model"
)

type Session struct {
	sessions *mongo.Collection
}

func NewSession(database *mongo.Database) Session {
	return Session{sessions: database.Collection("sessions")}
}

func (s Session) Update(ctx context.Context, session entity.Session) error {
	filter := bson.M{"_id": session.ID}

	_, err := s.sessions.ReplaceOne(ctx, filter, session)
	return err
}

func (s Session) Get(ctx context.Context, id bson.ObjectId) (entity.Session, error) {
	filter := bson.M{"_id": id}
	session := new(entity.Session)

	err := s.sessions.FindOne(ctx, filter).Decode(session)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return entity.Session{}, model.ErrNotFound
	}
	return *session, err
}

func (s Session) Create(ctx context.Context, session entity.Session) error {
	_, err := s.sessions.InsertOne(ctx, session)
	return err
}
