package app

import (
	"github.com/go-chi/chi/v5"
	"medods/config"
	"medods/internal/httpserver/handlers"
	"medods/internal/usecase"
	"medods/pkg/httpserver"
)

func newHttpServer(http config.HTTP, dependencies usecase.Dependencies) *httpserver.Server {
	r := chi.NewRouter()

	useCases := usecase.NewUseCases(dependencies)
	handlers.NewRouter(r, useCases)

	return httpserver.New(r,
		httpserver.Port(http.Port),
		httpserver.MaxHeaderBytes(http.MaxHeaderBytes),
		httpserver.IdleTimeout(http.IdleTimeout),
		httpserver.WriteTimeout(http.WriteTimeout),
		httpserver.ReadTimeout(http.ReadTimeout),
	)
}
