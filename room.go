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

	// when before room runing
	Init(gs *GameServer, clients *[]Client, config *RoomConfig)
}

//
type Room struct {
	roomEvents *RoomEvents
	GameServer *GameServer
	Clients    *[]Client
	Config     *RoomConfig
}

func (r *Room) Init(gs *GameServer, clients *[]Client, config *RoomConfig) {
	r.GameServer = gs
	r.Clients = clients
	r.Config = config
}

func (rm *Room) OnInit() {}

func (rm *Room) OnUpdate(delta float64) {}

func (rm *Room) OnMessage() {}

func (rm *Room) onJoinRequest(client *Client) bool { return true }
