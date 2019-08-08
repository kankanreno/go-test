package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"net/http"
)

// === Hub ===
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unRegister chan *Client
}

func (hub *Hub) start() {
	for {
		select {
		case conn := <-hub.register:
			hub.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: conn.id + "/A new socket has connected."})
			hub.send(jsonMessage, conn)
		case conn := <-hub.unRegister:
			if _, ok := hub.clients[conn]; ok {
				close(conn.send)
				delete(hub.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: conn.id + "/A socket has disconnected."})
				hub.send(jsonMessage, conn)
			}
		case message := <-hub.broadcast:
			for conn := range hub.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(hub.clients, conn)
				}
			}
		}
	}
}

func (hub *Hub) send(message []byte, ignore *Client) {
	for conn := range hub.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}


// === Client ===
type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

func (c *Client) read() {
	defer func() {
		hub.unRegister <- c
		c.socket.Close()
	}()
	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			hub.unRegister <- c
			c.socket.Close()
			break
		}
		var content map[string]string
		json.Unmarshal(message, &content)
		if user, ok := content["user"]; ok {
			fmt.Print(user)
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		hub.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var hub = Hub{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unRegister: make(chan *Client),
	clients:    make(map[*Client]bool),
}


func main() {
	fmt.Println("Starting application...")
	go hub.start()
	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(":8081", nil)
}

func wsPage(res http.ResponseWriter, req *http.Request) {
	u := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	conn, err := u.Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	newUUID := uuid.NewV4()
	client := &Client{id: newUUID.String(), socket: conn, send: make(chan []byte)}
	hub.register <- client
	go client.read()
	go client.write()
}
