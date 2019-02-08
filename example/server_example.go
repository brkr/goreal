package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brkr/goreal"
)

func main() {
	gameServer := goreal.NewGameServer(1111)
	log.Println(gameServer.Port)
	lobby := &Lobby{}

	gameServer.InitClient = func(w http.ResponseWriter, r *http.Request, client *goreal.Client) {
		log.Println("new user handler")
		gameServer.JoinRoom("lobby", client)
	}

	gameServer.RegisterRoom("lobby", lobby)
	gameServer.Start()

}

type GameState struct {
	counter int32
	Player1 *goreal.Client
	Player2 *goreal.Client
}

type Lobby struct {
	goreal.Room
	State *GameState
}

func (l *Lobby) OnJoinRequest(client *goreal.Client) bool {
	log.Println("lobby onJoinRequest")

	if len(l.Clients) >= 1 {
		log.Println("room is full")
		return false
	}

	return true
}

func (l *Lobby) OnInit() {
	l.Room.Config.SimulationTick = 1
	l.State = &GameState{}
}

func (l *Lobby) OnMessage(client *goreal.Client, message []byte)  {
	log.Printf("Lobby: Message : %s", string(message))
}

func (l *Lobby) OnClientJoin(client *goreal.Client) {
	log.Println("lobby onClientJoin")

	if l.State.Player1 == nil {
		client.SendMessage([]byte("You're player 1."))
		l.State.Player1 = client
	} else if l.State.Player2 == nil {
		client.SendMessage([]byte("You're player 1."))
		l.State.Player2 = client
	}
}

func (l *Lobby) OnUpdate(delta float64) {
	//log.Println("update game simulation delta time: ", delta)

	l.State.counter++
	l.BroadcastMessage(fmt.Sprintf("Selam. State %d", l.State.counter))

	if l.State.counter % 25 == 0 && l.State.Player1 != nil  {
		l.GameServer.
	}
}

func NewLobby() *Lobby {
	return &Lobby{}
}


