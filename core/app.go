package core

import "log/slog"

type App interface {
	// Logger returns the default app logger.
	Logger() *slog.Logger

	// Bootstrap initializes the application.
	Bootstrap() error

	// ServerPort returns the TCP port for the HTTP server to listen on.
	ServerPort() int

	// OnServerShutdown run jobs when the server receives a shutdown signal.
	OnServerShutdown()
}
