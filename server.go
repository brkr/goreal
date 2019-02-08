package goreal

import (
	"fmt"
	"log"
	"net/http"
)

//Game server manager.
type GameServer struct {
	// running port
	Port int16
	// all rooms in game server
	rooms map[string]*RoomManager
	// handle client's first connection
	InitClient func(w http.ResponseWriter, r *http.Request, client *Client)
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
		client := serveWs(hub, w, r)

		if gs.InitClient != nil {
			gs.InitClient(w, r, client)
		}
	})

	// connection string addr:port
	connectionString := fmt.Sprintf(":%d", gs.Port)

	err := http.ListenAndServe(connectionString, nil)
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
		panic("wrong room type")
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

// change the client's room
func (gs *GameServer) JoinRoom(roomName string, client *Client) bool {
	// todo check if room created
	// todo if client is another room, must be leave current room

	room, ok := gs.rooms[roomName]
	if !ok {
		log.Printf("Room not found! %s", roomName)
		return false
	}

	canJoin := room.CanJoinTheRoom(client)
	if !canJoin {
		log.Println("Room is not available for client")
		return false
	}

	// send client to room information
	client.SendMessageStr("{\"join_room\":\"" + roomName + "\"}")
	room.AddClientToRoom(client)

	return true
}

func (gs *GameServer) LeaveFromRoom(client *Client, room *Room )  {

	client.RemoveListener(room)


}
