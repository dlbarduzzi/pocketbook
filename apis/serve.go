package apis

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dlbarduzzi/pocketbook/core"
)

func Serve(app core.App) error {
	logger := app.Logger()
	router := baseRouter(app)

	handler, err := router.buildHandler()
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.ServerPort()),
		Handler:      handler,
		IdleTimeout:  time.Second * 60,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}

	shutdownErr := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		inSignal := <-quit

		logger.Info(
			"server received shutdown signal",
			slog.String("signal", inSignal.String()),
		)

		logger.Info("server running shutdown tasks...")
		app.OnServerShutdown()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			shutdownErr <- err
		}

		shutdownErr <- nil
	}()

	logger.Info("server starting", slog.Int("port", app.ServerPort()))

	err = server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErr
	if err != nil {
		return err
	}

	logger.Info("server stopped", slog.Int("port", app.ServerPort()))

	return nil
}
