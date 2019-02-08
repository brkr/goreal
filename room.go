package goreal

// game room
type RoomEvents interface {
	// client request to join to room
	OnJoinRequest(client *Client) bool

	// client joined the room
	OnClientJoin(client *Client)

	// when room created
	OnInit()
	// called each update patch
	OnUpdate(delta float64)
	//clienttan mesaj geldiginde
	OnMessage(client *Client, message []byte)

	Init(gs *GameServer, clients map[*Client]bool, config *RoomConfig)

	// client leave from the room.
	OnLeave(client *Client)

	// todo OnDisconnect()
}

//
type Room struct {
	Name string
	roomEvents *RoomEvents
	GameServer *GameServer
	Clients    map[*Client]bool
	Config     *RoomConfig
}

func (r *Room) Init(gs *GameServer, clients map[*Client]bool, config *RoomConfig) {
	r.GameServer = gs
	r.Clients = clients
	r.Config = config
}

func (rm *Room) OnInit() {}

func (rm *Room) OnUpdate(delta float64) {}

func (rm *Room) OnMessage(client *Client, message []byte) {}

func (rm *Room) OnJoinRequest(client *Client) bool { return true }

func (rm *Room) OnClientJoin(client *Client) {}

func (rm *Room) OnLeave(client *Client) {}

// tum clientlara broadcast mesaj gonderir.
func (rm *Room) BroadcastMessage(message string) {

	if len(rm.Clients) == 0 {
		return
	}

	for k := range rm.Clients {
		k.SendMessageStr(message)
	}
}



