package goreal

import (
	"log"
	"net/http"
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

func (gs *GameServer) Start() {
	log.Println("server starting..")
	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("ws request detected")
		//client connection starting
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

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
	log.Println("room added", len(gs.rooms))
	roomObj.Init(gs)
	roomObj.OnInit()
}
