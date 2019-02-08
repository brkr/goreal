package goreal

import (
	"log"
)

// Hub maintains the set of active Clients and broadcasts messages to the
// Clients.
type Hub struct {
	// Registered Clients.
	clients map[*Client]bool

	// Inbound messages from the Clients.
	broadcast chan []byte

	// Register requests from the Clients.
	register chan *Client

	// Unregister requests from Clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run(gs *GameServer) {
	for {
		select {
		case client := <-h.register:
			log.Println("client register")
			h.clients[client] = true

		case client := <-h.unregister:
			// client was disconnected
			log.Println("client unregister")

			gs.DisconnectClient(client)

			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			// send broadcast message to all Clients
			log.Println("message from client")
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
