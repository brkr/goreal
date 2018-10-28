package goreal

// game room
type RoomEvents interface {
	// client request to join to room
	onJoinRequest(client *Client) bool

	// client joined the room
	// onClientJoin(clietn *Client)

	// when room created
	OnInit()
	// called each update patch
	OnUpdate()
	//clienttan mesaj geldiginde
	OnMessage()

	Init(gs *GameServer)
}

//
type Room struct {
	room       *RoomEvents
	GameServer *GameServer
	clients    *[]Client
}

func (rm *Room) Init(gs *GameServer) {
	rm.GameServer = gs
}

func (rm *Room) OnInit() {}

func (rm *Room) OnUpdate() {}

func (rm *Room) OnMessage() {}

func (rm *Room) onJoinRequest(client *Client) bool { return true }
