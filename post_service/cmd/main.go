package main

import (
	app "post_service/internal/app"

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
