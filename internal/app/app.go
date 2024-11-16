package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/MaKYaro/sso/internal/app/grpc-app"
	"github.com/MaKYaro/sso/internal/config"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	gRPCPort int,
	DBConfig *config.DBConnectionConfig,
	tokenTTL time.Duration,
) *App {
	// init storage

	// init auth service

	gRPCApp := grpcapp.New(log, gRPCPort)

	return &App{
		GRPCServer: gRPCApp,
	}
}
