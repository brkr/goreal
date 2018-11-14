package main

import (
	"fmt"
	"log"
	"time"

	"github.com/brkr/goreal"
)

func main() {

	gameServer := goreal.NewGameServer(1111)
	log.Println(gameServer.Port)
	lobby := &Lobby{}
	log.Println(lobby)
	gameServer.RegisterRoom("lobby", lobby)
	gameServer.Start()
}

type Lobby struct {
	goreal.Room
}

func (l *Lobby) onJoinRequest(client *goreal.Client) bool {
	log.Println("lobby onJoinRequest")

	return true
}

func (l *Lobby) OnInit() {
	l.Room.OnInit()
	log.Println("lobby init")
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for t := range ticker.C {
		log.Println("Tick at", t)
		i = i + 1
		roomName := fmt.Sprintf("room-%d", i)

		log.Println("room name %s ", roomName)
		log.Println(l.GameServer)

		// l.GameServer.RegisterRoom(roomName, &Lobby{})
	}

	log.Println("denemeee")
}

func (l *Lobby) onClientJoin(clietn *goreal.Client) {
	log.Println("lobby onClientJoin")
}

func NewLobby() *Lobby {
	return &Lobby{}
}
