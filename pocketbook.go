package pocketbook

import (
	"github.com/dlbarduzzi/pocketbook/core"
)

// Ensures that the PocketBook implements the App interface.
var _ core.App = (*PocketBook)(nil)

type PocketBook struct {
	core.App
}

// Config is the PocketBook initialization config struct.
type Config struct {
	ServerPort   int
	DatabaseURL  string
	MaxOpenConns int
	MaxIdleConns int
}

func New() *PocketBook {
	return NewWithConfig(Config{})
}

func NewWithConfig(config Config) *PocketBook {
	pb := &PocketBook{}

	pb.App = core.NewBaseApp(core.BaseAppConfig{
		ServerPort:   config.ServerPort,
		DatabaseURL:  config.DatabaseURL,
		MaxOpenConns: config.MaxOpenConns,
		MaxIdleConns: config.MaxOpenConns,
	})

	return pb
}

func (pb *PocketBook) Start() error {
	return pb.Execute()
}

func (pb *PocketBook) Execute() error {
	if err := pb.Bootstrap(); err != nil {
		return err
	}
	return nil
}
