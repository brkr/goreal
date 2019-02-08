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
// new room manager
func newRoomManager(path string, roomEvents RoomEvents) *RoomManager {
	config := &RoomConfig{MaxUser: 1000, SimulationTick: 10}
	return &RoomManager{RoomEvents: roomEvents, Path: path, Config: config, clients: make(map[*Client]bool)}
}

func (rm *RoomManager) OnInit(gs *GameServer) {
	log.Println("room init", rm.Path)
	rm.RoomEvents.Init(gs, rm.clients, rm.Config)
}

func (rm *RoomManager) CanJoinTheRoom(client *Client) bool {
	// todo check if client already join

	_, ok := rm.clients[client]
	if ok {
		log.Println("Client is already join.")
		return true
	}

	canJoin := rm.RoomEvents.OnJoinRequest(client)
	return canJoin
}
func (rm *RoomManager) AddClientToRoom(client *Client) {

	rm.clients[client] = true

	// listen client's message
	client.ListenMessage(rm)

	// send join information to room
	rm.RoomEvents.OnClientJoin(client)
}

func (rm *RoomManager) RemoveClientFromRoom(client *Client) {

	_, ok := rm.clients[client]

	if !ok {
		log.Println("client is not in room!")
		return
	}

	rm.RoomEvents.OnLeave(client)

	client.RemoveListener(rm)

	delete(rm.clients, client)


}

func (rm *RoomManager) ReceiveMessage(client *Client, message []byte) {
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
