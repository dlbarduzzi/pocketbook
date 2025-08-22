package core

type App interface {
	// Bootstrap initializes the application.
	Bootstrap() error
}
