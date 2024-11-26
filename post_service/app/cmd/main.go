package main

import (
	server "facelessbook/post_service/app/internal"

	log "github.com/sirupsen/logrus"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatal()
		return
	}
	server.Run()
}
