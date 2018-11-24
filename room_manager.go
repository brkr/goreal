package goreal

import (
	"log"
	"time"

	"github.com/kutase/go-gameloop"
)

type RoomManager struct {
	Path       string
	RoomEvents RoomEvents
	clients    *[]Client
	Config     *RoomConfig
}

type RoomConfig struct {
	MaxUser        int
	SimulationTick int
}

func newRoomManager(path string, roomEvents RoomEvents) *RoomManager {
	config := &RoomConfig{MaxUser: 1000, SimulationTick: 15}
	return &RoomManager{RoomEvents: roomEvents, Path: path, Config: config}
}

func (rm *RoomManager) OnInit(gs *GameServer) {
	log.Println("room init", rm.Path)
	rm.RoomEvents.Init(gs, rm.clients, rm.Config)
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
