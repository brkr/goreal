package goreal

import (
	"log"
	"net/http"
)

//Game server manager.
type GameServer struct {
	Port  int16
	rooms map[string]*RoomManager
}

// create a new server
func NewGameServer(port int16) *GameServer {
	server := &GameServer{Port: port}
	server.rooms = make(map[string]*RoomManager)
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
func (gs *GameServer) RegisterRoom(path string, roomInterface interface{}) {
	// gs.rooms[path] = room
	if roomInterface == nil {
		log.Println("room instance not null")
		return
	}

	roomEvents, ok := roomInterface.(RoomEvents)

	if !ok {
		log.Printf("wrong room type")
		return
	}

	if _, ok := gs.rooms[path]; ok {
		log.Println("Room already created. {}", path)
		return
	}

	roomManager := newRoomManager(path, roomEvents)
	gs.rooms[path] = roomManager
	gs.bootstrapRoom(roomManager)
}

func (gs *GameServer) bootstrapRoom(roomManager *RoomManager) {
	roomManager.OnInit(gs)
	go roomManager.run()
}
