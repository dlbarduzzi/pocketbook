package main

import (
	"github.com/dlbarduzzi/pocketbook"
)

func main() {
	app := pocketbook.New()

	if err := app.Start(); err != nil {
		panic(err)
	}
}
