package main

import (
	"log"

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
	log.Println("lobby init")

}

func (l *Lobby) onClientJoin(clietn *goreal.Client) {
	log.Println("lobby onClientJoin")
}

func (l *Lobby) OnUpdate(delta float64) {
	log.Println("update game simulation delta time: ", delta)
}

func NewLobby() *Lobby {
	return &Lobby{}
}
