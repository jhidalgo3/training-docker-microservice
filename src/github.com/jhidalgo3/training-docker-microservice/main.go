package main

import (
	"fmt"

	"github.com/jhidalgo3/training-docker-microservice/config"
	"github.com/jhidalgo3/training-docker-microservice/server"
)

func main() {
	fmt.Printf("Commit %v\n", config.GetCommit())
	fmt.Printf("Version %v\n", config.GetVersion())
	server.StartServer()
}
