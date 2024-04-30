package app

import (
	"errors"
	"fmt"
	"log/slog"
	"medods/config"
	"medods/pkg/httpserver"
	"medods/pkg/mongo"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	httpServer *httpserver.Server
	db         mongo.Mongo
}

func New(cfg *config.Config) (App, error) {
	db, err := openDB(cfg.MONGO)
	if err != nil {
		return App{}, err
	}

	dependencies := newDependencies(cfg, db.DB)
	httpServer := newHttpServer(cfg.HTTP, dependencies)

	return App{httpServer: httpServer, db: db}, nil
}

func (app App) Run() error {
	app.httpServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var err error
	select {
	case s := <-interrupt:
		err = errors.New("app - Run - signal: " + s.String())
	case err = <-app.httpServer.Notify():
		err = fmt.Errorf("app - Run - httpServer.Notify: %w", err)
	}

	return err
}

func (app App) Shutdown() error {
	httpErr := app.httpServer.Shutdown()
	if httpErr != nil {
		slog.Error(fmt.Sprintf("app - Run - httpServer.Shutdown: %s", httpErr))
	}
	dbErr := app.db.Disconnect()
	if dbErr != nil {
		slog.Error(fmt.Sprintf("app - Run - db.Close: %s", dbErr))
	}
	return errors.Join(httpErr, dbErr)
}
