package main

import (
	"log"

	"github.com/brkr/goreal"
)

func main() {
	gameServer := goreal.NewGameServer(1111)
	log.Println(gameServer.Port)
	lobby := &Lobby{}
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
	l.Room.Config.SimulationTick = 10
}

func (l *Lobby) onClientJoin(client *goreal.Client) {
	log.Println("lobby onClientJoin")
}

func (l *Lobby) OnUpdate(delta float64) {
	log.Println("update game simulation delta time: ", delta)

}

func NewLobby() *Lobby {
	return &Lobby{}
}
