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
}

//
type Room struct {
	room    *RoomEvents
	clients *[]Client
}

func (rm *Room) OnInit() {}

func (rm *Room) OnUpdate() {}

func (rm *Room) OnMessage() {}

func (rm *Room) onJoinRequest(client *Client) bool { return true }
