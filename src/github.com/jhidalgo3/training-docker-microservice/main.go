package main

import "github.com/jhidalgo3/training-docker-microservice/server"
import "fmt"

var Version string

func main() {
	fmt.Printf("%s", Version)
	server.StartServer()
}
