package goreal

import (
	"log"

	"github.com/kutase/go-gameloop"
)

type RoomManager struct {
	Path       string
	RoomEvents RoomEvents
	clients    *[]Client
}

func newRoomManager(path string, roomEvents RoomEvents) *RoomManager {
	return &RoomManager{RoomEvents: roomEvents, Path: path}
}

func (rm *RoomManager) OnInit(gs *GameServer) {
	log.Println("room init", rm.Path)
	rm.RoomEvents.OnInit()
	rm.RoomEvents.Init(gs, rm.clients)
}

func (rm *RoomManager) run() {
	log.Println("start {} room manager.", rm.Path)

	gl := gameLoop.New(1, func(delta float64) {
		log.Println("tick:", delta)
		rm.RoomEvents.OnUpdate(delta)
	})

	gl.Start()

}
