package main

import (
	"log"
	"social-network-go/server"
)

func main() {
	err := server.InitServer()
	if err != nil {
		log.Fatal("Error initializing server: ", err)
	}

}
