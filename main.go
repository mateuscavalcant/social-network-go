package main

import (
	"log"
	"social-network-go/api"
)

func main() {
	err := api.Handler()
	if err != nil {
		log.Fatal("Error initializing server: ", err)
	}

}
