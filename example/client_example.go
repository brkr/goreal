package main

import (
	"github.com/brkr/goreal"
	"log"
	"time"
)

func main() {

	log.Println("client started")

	gameClient := goreal.GameClient{ConnectionString: ":1111"}
	gameClient.OnMessage(func(message []byte) {
		log.Println(string(message))
	})
	gameClient.Connect()

	ticker := time.NewTicker(time.Second * 15)
	defer ticker.Stop()

	for {
		select {

		case _ = <-ticker.C:
			log.Println("tick")
			err := gameClient.SendMessage("send from client")
			if err != nil {
				log.Println("write:", err)
				return
			}

		case <-time.After(time.Second):
		}
		return
	}
}
