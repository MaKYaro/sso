package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	authgrpc "github.com/MaKYaro/sso/internal/grpc/auth"
	"google.golang.org/grpc"
)

// App is gRPC app instance
type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// New creates new gRPC server app
func New(
	log *slog.Logger,
	port int,
) *App {
	gRPCServer := grpc.NewServer()

	// Connect serverAPI handler to gRPCServer
	authgrpc.Register(gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// MustRun starts gRPC app and panics if any error occurs
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run starts gRPC app
func (a *App) Run() error {
	const op = "internal.app.grpcapp.Run"

	a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	address := fmt.Sprintf(":%d", a.port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.Info(
		"gRPC server is running",
		slog.String("address", address),
	)

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop method makes graceful stop of the app
func (a *App) Stop() {
	const op = "internal.app.grpcapp.Stop"

	a.log.With(slog.String("op", op))
	a.log.Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
