package handler

import (
	"log"
	"social-network-go/handler"
	
)

func Handler() {
	err := handler.InitServer()
	if err != nil {
		log.Fatal("Error initializing server: ", err)
	}
}