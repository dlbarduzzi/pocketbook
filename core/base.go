package core

// Ensures that the BaseApp implements the App interface.
var _ App = (*BaseApp)(nil)

type BaseApp struct{}

func (app *BaseApp) Bootstrap() error {
	return nil
}
