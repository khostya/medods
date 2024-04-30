package session

import (
	"github.com/globalsign/mgo/bson"
	"medods/internal/entity"
	"medods/internal/lib/auth"
	"medods/internal/model"
	"time"
)

type Session struct {
	tokenManager          auth.TokenManager
	accessTTL, refreshTTL time.Duration
}

func New(tokenManager auth.TokenManager, accessTTL, refreshTTL time.Duration) Session {
	return Session{
		tokenManager: tokenManager,
		accessTTL:    accessTTL,
		refreshTTL:   refreshTTL,
	}
}

func (s Session) Create(userID string) (entity.Session, model.Tokens, error) {
	return s.create(userID, bson.NewObjectId())
}

func (s Session) Refresh(userID string, ID bson.ObjectId) (entity.Session, model.Tokens, error) {
	return s.create(userID, ID)
}

func (s Session) create(userID string, ID bson.ObjectId) (entity.Session, model.Tokens, error) {
	var (
		err       error
		tokens    model.Tokens
		expiresIn = time.Now().Add(s.accessTTL)
		session   = entity.Session{UserID: userID, ID: ID}
	)

	accessToken, err := s.tokenManager.NewJWT(session.ID.Hex(), session.UserID, expiresIn)
	if err != nil {
		return session, tokens, err
	}

	session.RefreshToken = s.tokenManager.NewRefreshToken()
	session.ExpiresIn = time.Now().Add(s.refreshTTL)
	tokens = model.NewTokens(accessToken, session.RefreshToken)

	return session, tokens, nil
}
