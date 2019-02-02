package goreal

import (
	"log"
	"time"

	"github.com/kutase/go-gameloop"
)

type RoomManager struct {
	Path       string
	RoomEvents RoomEvents
	clients    map[*Client]bool
	Config     *RoomConfig
}

type RoomConfig struct {
	MaxUser        int
	SimulationTick int
}

func newRoomManager(path string, roomEvents RoomEvents) *RoomManager {
	config := &RoomConfig{MaxUser: 1000, SimulationTick: 30}
	return &RoomManager{RoomEvents: roomEvents, Path: path, Config: config, clients: make(map[*Client]bool)}
}

func (rm *RoomManager) OnInit(gs *GameServer) {
	log.Println("room init", rm.Path)
	rm.RoomEvents.Init(gs, rm.clients, rm.Config)
}

func (rm *RoomManager) OnClientJoin(client *Client) bool {
	canJoin := rm.RoomEvents.OnJoinRequest(client)
	if !canJoin {
		log.Println("Room is not available for client")
		return false
	}

	rm.clients[client] = true

	log.Printf("len=%d", len(rm.clients))

	rm.RoomEvents.OnClientJoin(client)

	//register for listen to client messages
	client.ListenMessage(rm)

	return true
}

func (rm *RoomManager) ReceiveMessage(client *Client, message []byte){
	rm.RoomEvents.OnMessage(client, message)
}

func (rm *RoomManager) run() {
	log.Println("start {} room manager.", rm.Path)
	rm.RoomEvents.OnInit()

	duration := rm.Config.SimulationTick

	gl := gameLoop.New(time.Duration(duration), func(delta float64) {
		rm.RoomEvents.OnUpdate(delta)
	})

	gl.Start()

}
