package main

import (
	"log/slog"
	"os"

	"oss-backend/internal/bootstrap"
	"oss-backend/internal/config"
	"oss-backend/internal/logger"
)

func main() {
	deps, err := bootstrap.Up()
	if err != nil {
		panic(err)
	}

	slog.SetDefault(slog.New(logger.NewLogger(os.Stdout, config.Get().Log.Level)))

	slog.Info("Starting HTTP server on " + deps.Config.Address)

	if err = deps.HTTPServer.Start(); err != nil {
		panic(err)
	}
}
