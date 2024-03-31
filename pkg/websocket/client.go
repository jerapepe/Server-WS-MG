package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user
type client struct {
	//socket is the web socket for this client
	socket *websocket.Conn
	//recive is a channel to recive messages from other clients
	recive chan []byte
	//room is the room this client is chatting in
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.recive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error writing to socket:", err)
			return
		}
	}
}
