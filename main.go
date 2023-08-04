package main

import (
	"log"
	"social-network-go/api"
)

func main() {
	err := api.InitServer()
	if err != nil {
		log.Fatal("Error initializing server: ", err)
	}

}
