package goreal

// game room
type Room interface {
	// client request to join to room
	onJoinRequest(client *Client) bool

	// client joined the room
	onClientJoin(clietn *Client)

	// when room registered
	onInit()
}
