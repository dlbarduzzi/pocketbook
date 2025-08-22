package pocketbook

import (
	"fmt"

	"github.com/dlbarduzzi/pocketbook/core"
)

// Ensures that the PocketBook implements the App interface.
var _ core.App = (*PocketBook)(nil)

type PocketBook struct {
	core.App
}

type Config struct {
	isDev bool
}

func New() *PocketBook {
	return NewWithConfig(Config{
		isDev: true,
	})
}

func NewWithConfig(config Config) *PocketBook {
	pb := &PocketBook{}

	if config.isDev {
		fmt.Println("Running in dev mode")
	}

	return pb
}

func (pb *PocketBook) Start() error {
	return pb.Execute()
}

func (pb *PocketBook) Execute() error {
	fmt.Println("Executing...")
	return nil
}
