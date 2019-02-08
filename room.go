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

	Init(gs *GameServer, clients map[*Client]bool, config *RoomConfig, operation interface{})

	// client leave from the room.
	OnLeave(client *Client)

	// client disconnected from server
	OnDisconnect(client *Client)
}

type RoomOperation interface {
	// kick client from the room
	Kick(client *Client)
}

//
type Room struct {
	Name string
	GameServer *GameServer
	Clients    map[*Client]bool
	Config     *RoomConfig
	RoomOperation  RoomOperation
}

func (r *Room) Init(gs *GameServer, clients map[*Client]bool, config *RoomConfig, operation interface{}) {
	r.GameServer = gs
	r.Clients = clients
	r.Config = config

	o := operation.(RoomOperation)
	r.RoomOperation = o
}

func (rm *Room) OnInit() {}

func (rm *Room) OnUpdate(delta float64) {}

func (rm *Room) OnMessage(client *Client, message []byte) {}

func (rm *Room) OnJoinRequest(client *Client) bool { return true }

func (rm *Room) OnClientJoin(client *Client) {}

func (rm *Room) OnLeave(client *Client) {}
func (rm *Room) OnDisconnect(client *Client) {}

// tum clientlara broadcast mesaj gonderir.
func (rm *Room) BroadcastMessage(message string) {

	if len(rm.Clients) == 0 {
		return
	}

	for k := range rm.Clients {
		k.SendMessageStr(message)
	}
}



