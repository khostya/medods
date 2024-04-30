package main

import (
	"log/slog"
	"medods/config"
	"medods/internal/app"
	"os"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	app, err := app.New(cfg)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer app.Shutdown()

	if err := app.Run(); err != nil {
		slog.Error(err.Error())
	}
}
