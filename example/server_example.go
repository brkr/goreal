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
	l.Room.Config.SimulationTick = 1
}

func (l *Lobby) OnMessage(client *goreal.Client, message []byte)  {
	log.Printf("Lobby: Message : %s", string(message))
}

func (l *Lobby) OnClientJoin(client *goreal.Client) {
	log.Println("lobby onClientJoin")
	client.SendMessage([]byte("hello world"))
}

func (l *Lobby) OnUpdate(delta float64) {
	//log.Println("update game simulation delta time: ", delta)

	l.BroadcastMessage([] byte("selam herkese"))

}

func NewLobby() *Lobby {
	return &Lobby{}
}
