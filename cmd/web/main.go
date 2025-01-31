package main

import (
	"github.com/codedbyshoe/grit/internal/app"
)

func main() {
	app := app.NewApplication()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
