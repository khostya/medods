package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "medods/docs"
	"medods/internal/httpserver/handlers/auth"
	"medods/internal/usecase"
	"net/http"
)

// @title       medods
// @version     1.0
// @BasePath    /
func NewRouter(h *chi.Mux, useCases usecase.UseCases) {
	h.Use(cors.AllowAll().Handler)
	h.Use(middleware.Recoverer)
	h.Use(middleware.RequestID)

	h.Get("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	swaggerHandler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)
	h.Get("/swagger/*", swaggerHandler)

	h.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Location", "/swagger/index.html")
		writer.WriteHeader(http.StatusTemporaryRedirect)
	})

	h.Group(func(r chi.Router) {
		auth.New(r, useCases)
	})
}
