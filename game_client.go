package goreal

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// game client
type GameClient struct {
	ConnectionString string //localhost:8080
	conn             *websocket.Conn
	listen           func(message []byte)
}

func (gc *GameClient) Connect() {
	// isletim sisteminden gele
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{Scheme: "ws", Host: gc.ConnectionString, Path: "/ws"}
	log.Printf("connecting to %s", u.String())
	header := http.Header{}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), header)

	gc.conn = conn

	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// listen connection
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			//log.Printf("recv: %s", message)
			if gc.listen != nil {
				gc.listen(message)
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

	log.Println("connection complete")
}

// Send message to server
func (gc *GameClient) SendMessage(message string) error {
	return gc.conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func (gc *GameClient) OnMessage(f func(message []byte)) {
	gc.listen = f
}
