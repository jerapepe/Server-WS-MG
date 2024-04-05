package websocket

import (
	"Server-WS/pkg/mongo"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	//client holds all current clients in this room
	clients map[*client]bool
	//join is a channel for clients wishing to join the room
	join chan *client
	//leave is a channel for clients wishing to leave the room
	leave chan *client
	//forward is a channel that holds incoming messages
	forward chan []byte
}

// newRoom create a new chat room
func NewRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) Run() {
	for {
		select {
		case client := <-r.join:
			//joining
			r.clients[client] = true
		case client := <-r.leave:
			//leaving
			delete(r.clients, client)
			close(client.recive)
		case msg := <-r.forward:
			//forward message to all clients
			var msgs mongo.Message
			err := json.Unmarshal(msg, &msgs)
			if err != nil {
				log.Println("Error unmarshaling message:", err)
			}
			//valid := mongo.SaveMessageToMongo(msgs)
			//if !valid {
			//	log.Println("Error saving message to mongo")
			//}

			for client := range r.clients {
				select {
				case client.recive <- msg:
				default:
					delete(r.clients, client)
					close(client.recive)
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		recive: make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
