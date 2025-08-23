package core

import (
	"fmt"
	"log/slog"
	"net/url"

	"github.com/dlbarduzzi/pocketbook/tools/registry"
)

const (
	defaultServerPort   int = 8000
	defaultMaxOpenConns int = 10
	defaultMaxIdleConns int = 5
)

type BaseAppConfig struct {
	ServerPort   int
	DatabaseURL  string
	MaxOpenConns int
	MaxIdleConns int
}

// Ensures that the BaseApp implements the App interface.
var _ App = (*BaseApp)(nil)

type BaseApp struct {
	logger *slog.Logger
	config *BaseAppConfig
}

func NewBaseApp(config BaseAppConfig) *BaseApp {
	app := &BaseApp{
		config: &config,
	}

	if app.config.ServerPort < 1 {
		app.config.ServerPort = defaultServerPort
	}

	if app.config.MaxOpenConns < 1 {
		app.config.MaxOpenConns = defaultMaxOpenConns
	}

	if app.config.MaxIdleConns < 1 {
		app.config.MaxIdleConns = defaultMaxIdleConns
	}

	return app
}

func (app *BaseApp) Logger() *slog.Logger {
	if app.logger == nil {
		return slog.Default()
	}
	return app.logger
}

func (app *BaseApp) Bootstrap() error {
	if err := app.initRegistry(); err != nil {
		return err
	}

	if err := app.initLogger(); err != nil {
		return err
	}

	return nil
}

func (app *BaseApp) ServerPort() int {
	return app.config.ServerPort
}

func (app *BaseApp) OnServerShutdown() {
	fmt.Println("...Implement shutdown tasks...")
}

func (app *BaseApp) initRegistry() error {
	reg, err := registry.NewRegistry()
	if err != nil {
		return fmt.Errorf("failed to initialize registry - %v", err)
	}

	serverPort := reg.GetInt("PB_SERVER_PORT")
	if serverPort > 0 {
		app.config.ServerPort = serverPort
	}

	databaseURL := reg.GetString("PB_DATABASE_URL")
	if databaseURL != "" {
		app.config.DatabaseURL = databaseURL
	}

	maxOpenConns := reg.GetInt("PB_MAX_OPEN_CONNS")
	if maxOpenConns > 0 {
		app.config.MaxOpenConns = maxOpenConns
	}

	maxIdleConns := reg.GetInt("PB_MAX_IDLE_CONNS")
	if maxIdleConns > 0 {
		app.config.MaxIdleConns = maxIdleConns
	}

	return app.validateRegistry()
}

func (app *BaseApp) initLogger() error {
	app.logger = slog.Default()
	return nil
}

func (app *BaseApp) validateRegistry() error {
	if _, err := url.ParseRequestURI(app.config.DatabaseURL); err != nil {
		return fmt.Errorf("invalid database url - %v", err)
	}
	return nil
}
