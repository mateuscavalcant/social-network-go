package main

import (
	"log"
	"social-network-go/handler"
)

func Main() {
	err := handler.Handler()
	if err != nil {
		log.Fatal("Error initializing server: ", err)
	}

}
