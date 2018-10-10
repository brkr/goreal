package goreal

import (
	"log"
)

//Game server manager.
type GameServer struct {
	Port  int16
	rooms map[string]*Room
}

// create a new server
func NewGameServer(port int16) *GameServer {
	server := &GameServer{Port: port}
	server.rooms = make(map[string]*Room)
	server.init()
	return server
}

func (gs *GameServer) init() {
	log.Println("game server init")
	log.Println(gs.rooms)
}

// add room to server
func (gs *GameServer) RegisterRoom(path string, room Room) {
	// gs.rooms[path] = room
	if room == nil {
		log.Println("room instance not null")
		return
	}
	room.OnInit()
	log.Println("room added", len(gs.rooms))
}
