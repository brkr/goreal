package goreal

import (
	"log"
)

//Game server manager.
type GameServer struct {
	Port  int16
	rooms map[string]*RoomEvents
}

// create a new server
func NewGameServer(port int16) *GameServer {
	server := &GameServer{Port: port}
	server.rooms = make(map[string]*RoomEvents)
	server.init()
	return server
}

func (gs *GameServer) init() {
	log.Println("game server init")
	log.Println(gs.rooms)
}

// add room to server
func (gs *GameServer) RegisterRoom(path string, room interface{}) {
	// gs.rooms[path] = room
	if room == nil {
		log.Println("room instance not null")
		return
	}
	roomObj, ok := room.(RoomEvents)
	if !ok {
		log.Printf("wrong room type")
		return
	}

	gs.rooms[path] = &roomObj

	roomObj.OnInit()
	log.Println("room added", len(gs.rooms))
}
