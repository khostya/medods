package usecase

import (
	"medods/internal/lib/auth"
	"medods/internal/lib/session"
	"medods/internal/usecase/repo/mongo"
	"medods/pkg/hash"
)

type UseCases struct {
	Auth         Auth
	Dependencies Dependencies
}

type Dependencies struct {
	Repos              mongo.Repositories
	TokenManager       auth.TokenManager
	PasswordHasher     hash.Bcrypt
	Session            session.Session
	RefreshTokenHasher hash.Bcrypt
}

func NewUseCases(dependencies Dependencies) UseCases {
	repos := dependencies.Repos
	return UseCases{
		Dependencies: dependencies,
		Auth:         NewAuthUseCase(repos.Session, dependencies.Session, dependencies.RefreshTokenHasher),
	}
}
