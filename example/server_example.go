package main

import (
	"fmt"
	"github.com/brkr/goreal"
	"log"
	"net/http"
	"sync/atomic"
)

func main() {
	gameServer := goreal.NewGameServer(1111)
	log.Println(gameServer.Port)
	//lobby := &Lobby{}

	var roomNumber uint64

	gameServer.InitClient = func(w http.ResponseWriter, r *http.Request, client *goreal.Client) {

		isJoinRoom := gameServer.JoinRoom("lobby", client)
		if !isJoinRoom {
			// user not in
			roomName := fmt.Sprintf("lobby-%d", roomNumber)
			atomic.AddUint64(&roomNumber, 1)
			log.Printf("room count : %d", roomNumber)
			log.Println(roomName)
			room := &Lobby{}
			gameServer.RegisterRoom(roomName, room)
			gameServer.JoinRoom(roomName, client)

		}
	}

	//gameServer.RegisterRoom("lobby", lobby)
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

func (l *Lobby) OnMessage(client *goreal.Client, message []byte) {
	log.Printf("Lobby: Message : %s", string(message))
}

func (l *Lobby) OnClientJoin(client *goreal.Client) {
	log.Println("lobby onClientJoin")
	client.SendMessage([]byte("Welcome to game room"))
	if l.State.Player1 == nil {
		client.SendMessage([]byte("You're player 1."))
		l.State.Player1 = client
	} else if l.State.Player2 == nil {
		client.SendMessage([]byte("You're player 2."))
		l.State.Player2 = client
	}
}

func (l *Lobby) OnLeave(client *goreal.Client) {
	client.SendMessageStr("You're kicked from room.")
	if l.State.Player1 == client {
		log.Println("Player 1 leave from room")
		l.State.Player1 = nil
	} else if l.State.Player2 == client {
		log.Println("Player 2 leave from room")
		l.State.Player2 = nil
	}
}

func (l *Lobby) OnUpdate(delta float64) {
	//log.Println("update game simulation delta time: ", delta)

	l.State.counter++
	l.BroadcastMessage(fmt.Sprintf("User Count %d .State %d", len(l.Clients), l.State.counter))

	if l.State.counter%5 == 0 && l.State.Player1 != nil {

		//l.RoomOperation.Kick(l.State.Player1)
		//panic("unknown panic")
	}
}

func NewLobby() *Lobby {
	return &Lobby{}
}
