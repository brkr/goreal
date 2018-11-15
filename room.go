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
	OnUpdate(delta float64)
	//clienttan mesaj geldiginde
	OnMessage()

	Init(gs *GameServer, clients *[]Client)
}

//
type Room struct {
	roomEvents *RoomEvents
	GameServer *GameServer
	clients    *[]Client
}

func (rm *Room) Init(gs *GameServer, clients *[]Client) {
	rm.GameServer = gs
	rm.clients = clients
}

func (rm *Room) OnInit() {}

func (rm *Room) OnUpdate(delta float64) {}

func (rm *Room) OnMessage() {}

func (rm *Room) onJoinRequest(client *Client) bool { return true }
