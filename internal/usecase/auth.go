package usecase

import (
	"context"
	"errors"
	"github.com/globalsign/mgo/bson"
	"medods/internal/lib/session"
	"medods/internal/model"
	"medods/internal/usecase/repo/mongo"
	"medods/pkg/hash"
)

type Auth struct {
	session            session.Session
	sessionRepo        mongo.Session
	refreshTokenHasher hash.Bcrypt
}

func NewAuthUseCase(sessionRepo mongo.Session, session session.Session, refreshTokenHasher hash.Bcrypt) Auth {
	return Auth{
		sessionRepo:        sessionRepo,
		session:            session,
		refreshTokenHasher: refreshTokenHasher,
	}
}

func (uc Auth) Access(ctx context.Context, userID string) (model.Tokens, error) {
	session, tokens, err := uc.session.Create(userID)
	if err != nil {
		return model.Tokens{}, err
	}

	hashedRefreshToken, err := uc.refreshTokenHasher.Hash(session.RefreshToken)
	if err != nil {
		return model.Tokens{}, err
	}

	session.RefreshToken = hashedRefreshToken
	err = uc.sessionRepo.Create(ctx, session)

	return tokens, err
}

func (uc Auth) Refresh(ctx context.Context, id bson.ObjectId, refreshToken, userID string) (model.Tokens, error) {
	session, err := uc.sessionRepo.Get(ctx, id)
	if err != nil {
		return model.Tokens{}, err
	}

	if userID != session.UserID {
		return model.Tokens{}, errors.New("invalid token")
	}

	if !uc.refreshTokenHasher.Equals(session.RefreshToken, refreshToken) {
		return model.Tokens{}, errors.New("bad refresh token")
	}

	session, tokens, err := uc.session.Refresh(session.UserID, session.ID)
	if err != nil {
		return model.Tokens{}, err
	}

	session.RefreshToken, err = uc.refreshTokenHasher.Hash(session.RefreshToken)
	if err != nil {
		return model.Tokens{}, err
	}

	err = uc.sessionRepo.Update(ctx, session)
	return tokens, err
}
