package main

import (
	"task-exporter/internal/api"
)

func main() {
	api, err := api.New()
	if err != nil {
		panic(1)
	}
	api.Run()
}
