package main

import (
	app "api_gateway/internal/app"

	log "github.com/sirupsen/logrus"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal()
		return
	}
	app.Run()
}
