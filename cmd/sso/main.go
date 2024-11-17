package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/MaKYaro/sso/internal/app"
	"github.com/MaKYaro/sso/internal/config"
	log "github.com/MaKYaro/sso/internal/logger"
)

func main() {

	// load config
	cfg := config.MustLoad()

	// init logger
	log := log.New(cfg.Env)
	log.Debug("debug messages are enabled")

	// init application
	application := app.New(log, cfg.GRPCServer.Port, &cfg.DBConnection, cfg.TokenTTL)

	// start application
	log.Info(
		"starting application",
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.GRPCServer.Port),
	)
	log.Debug("config", slog.Any("", cfg))

	go application.GRPCServer.MustRun()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("application stopped")
}
