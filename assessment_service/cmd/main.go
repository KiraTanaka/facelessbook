package main

import (
	server "assessment_service/internal"

	log "github.com/sirupsen/logrus"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatal()
		return
	}
	//server.Run()
	server.GRPCServer.MustRun()
}
