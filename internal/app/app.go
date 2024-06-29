package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"sso/internal/config"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	dbConn config.DBConnConfig,
	tokenTTL time.Duration,
) *App {
	// TODO: init storage

	// TODO: init auth service

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{grpcApp}
}
