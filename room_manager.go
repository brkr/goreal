package goreal

import (
	"log"
)

type RoomManager struct {
	Path       string
	RoomEvents *RoomEvents
}

func newRoomManager(path string, roomEvents *RoomEvents) *RoomManager {
	return &RoomManager{RoomEvents: roomEvents, Path: path}
}

func (rm *RoomManager) run() {
	log.Println("start {} room manager.", rm.Path)
}
