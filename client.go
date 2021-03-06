package goreal

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// client message's listeners
	listeners map[ClientListener]bool

	data map[string]interface{}
}

type ClientListener interface {
	// message from client
	ReceiveMessage(client *Client, message []byte)
}

// send message to client
func (c *Client) SendMessage(message []byte) {
	c.send <- message
}

func (c *Client) SendMessageStr(message string) {
	c.SendMessage([]byte(message))
}

//
func (c *Client) ListenMessage(listener interface{}) {

	clientListener, ok := listener.(ClientListener)

	if ok {
		c.listeners[clientListener] = true
	}
}

func (c *Client) RemoveListener(listener interface{}) {
	clientListener, ok := listener.(ClientListener)

	if ok {
		delete(c.listeners, clientListener)
	}
}

func (c *Client) CloseConnection() {
	c.hub.unregister <- c
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {

		c.hub.unregister <- c
		//log.Println("client disconnecting...")
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		//log.Println("message : ", message)
		for k, _ := range c.listeners {
			k.ReceiveMessage(c, message)
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			//n := len(c.send)
			//for i := 0; i < n; i++ {
			//	w.Write(newline)
			//	w.Write(<-c.send)
			//}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) *Client {
	// disable check origin
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return nil
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), listeners: make(map[ClientListener]bool), data: make(map[string]interface{})}
	client.hub.register <- client
	log.Println("serve ws starting..")

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()

	return client
}

// put data
func (c *Client) Put(key string, data interface{}) {
	c.data[key] = data
}

func (c *Client) Get(key string) interface{} {
	return c.data[key]
}
