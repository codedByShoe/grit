package main

import "github.com/codedbyshoe/grit/internal/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
