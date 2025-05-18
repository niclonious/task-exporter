package main

import (
	"task-exporter/internal/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		panic(1)
	}
	app.Run()
}
