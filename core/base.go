package core

import "fmt"

type BaseAppConfig struct {
	DatabaseURL  string
	MaxOpenConns int
	MaxIdleConns int
}

// Ensures that the BaseApp implements the App interface.
var _ App = (*BaseApp)(nil)

type BaseApp struct {
	config *BaseAppConfig
}

func NewBaseApp(config BaseAppConfig) *BaseApp {
	app := &BaseApp{
		config: &config,
	}

	if app.config.DatabaseURL == "" {
		fmt.Println("[WARN] - config DatabaseURL is unset")
	}

	if app.config.MaxOpenConns < 1 {
		fmt.Println("[WARN] - config MaxOpenConns is unset")
	}

	if app.config.MaxIdleConns < 1 {
		fmt.Println("[WARN] - config MaxIdleConns is unset")
	}

	return app
}

func (app *BaseApp) Bootstrap() error {
	fmt.Println("[INFO] - bootstrapping app")
	return nil
}
