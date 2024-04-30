package auth

import (
	"github.com/go-chi/chi/v5"
	"medods/internal/lib/auth"
	"medods/internal/usecase"
)

type Router struct {
	useCase      usecase.Auth
	tokenManager auth.TokenManager
}

func New(r chi.Router, useCases usecase.UseCases) {
	router := Router{useCase: useCases.Auth, tokenManager: useCases.Dependencies.TokenManager}

	r.Group(func(r chi.Router) {
		r.Get("/access", router.access)
		r.Post("/refresh", router.refresh)
	})
}
