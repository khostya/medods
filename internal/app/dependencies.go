package app

import (
	"go.mongodb.org/mongo-driver/mongo"
	"medods/config"
	"medods/internal/lib/auth"
	"medods/internal/lib/session"
	"medods/internal/usecase"
	repo "medods/internal/usecase/repo/mongo"
	"medods/pkg/hash"
)

func newDependencies(cfg *config.Config, db *mongo.Database) usecase.Dependencies {
	repositories := repo.NewRepositories(db)

	refreshTokenHasher := hash.NewBcrypt(cfg.JWT.RefreshTokenCost)
	passwordHasher := hash.NewBcrypt(cfg.Auth.PasswordCostBcrypt)

	tokenManager := auth.NewManager(cfg.JWT.SigningKey)
	session := session.New(tokenManager, cfg.Auth.AccessTokenTTL, cfg.Auth.RefreshTokenTTL)

	return usecase.Dependencies{
		Repos:              repositories,
		TokenManager:       tokenManager,
		PasswordHasher:     passwordHasher,
		Session:            session,
		RefreshTokenHasher: refreshTokenHasher,
	}
}
