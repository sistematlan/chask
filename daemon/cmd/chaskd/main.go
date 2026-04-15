package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/sistematlan/chask/daemon/internal/api"
)

const socketPath = "/var/run/chaskd.sock"

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	slog.Info("starting chaskd", "socket", socketPath)

	server, err := api.NewServer(socketPath)
	if err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}

	// Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		slog.Info("shutting down chaskd")
		server.Close()
		os.Exit(0)
	}()

	if err := server.Serve(); err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}
}
