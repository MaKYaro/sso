package main

import (
	"log/slog"

	"github.com/MaKYaro/sso/internal/config"
	log "github.com/MaKYaro/sso/internal/logger"
)

func main() {

	// load config
	cfg := config.MustLoad()

	// init logger
	logger := log.New(cfg.Env)
	logger.Debug("debug messages are enabled")

	// start application
	logger.Info(
		"starting application",
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.GRPCServer.Port),
	)
	logger.Debug("config", slog.Any("", cfg))

}
