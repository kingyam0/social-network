package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 512                 // Maximum message size allowed from peer.
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub
type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn // The websocket connection
	Send   chan []byte     // Buffered channel of outbound messages
	UserId string          // The user id of the client
}

/* -------- ReadPump is ran in a per-connection goroutine. -------- */
/* --- The application ensures that there is at most one reader --- */
/* -- on a connection by executing all reads from this goroutine -- */
// Pumps messages from the websocket connection to the hub
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		// fmt.Println(string(message))
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Hub.Broadcast <- message
	}
}

/* -------- WritePump is called for each connection. It is -------- */
/* --------- ensured that there is at most one writer to a -------- */
/* ---- connection by executing all writes from this goroutine ---- */
// Pumps messages from the hub to the websocket connection
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			// c.Conn.WriteJSON(message)
			w.Write(message)
			// Add queued chat messages to the current websocket message
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Handles websocket requests from the peer
func (data *DB) ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// TODO: userID need to be replace by the userId of the current user from the cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		log.Println("Error reading cookie:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("Cookie value:", cookie.Value)
	fmt.Println("cookie: ", strings.Split(cookie.String(), "&")[1])
	client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256), UserId: strings.Split(cookie.String(), "&")[1]}
	// fmt.Println("ServeWs", client.UserId, client.Conn)
	client.Hub.Register <- client
	// Allow collection of memory referenced by the caller by doing all work in new goroutines
	go client.WritePump()
	go client.ReadPump()
}

// Hub maintains the set of active clients and broadcasts messages to the clients
type Hub struct {
	Clients    map[string]*Client // Registered clients
	Broadcast  chan []byte        // Inbound messages from the clients
	Register   chan *Client       // Register requests from the clients
	Unregister chan *Client       // Unregister requests from clients
	Database   *DB
}

func NewHub(db *DB) *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
		Database:   db,
	}
}
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.UserId] = client
		case client := <-h.Unregister:
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("USERID:   ", client.UserId)
			fmt.Println()
			fmt.Println("CLient hub: ", h.Clients)
			fmt.Println()
			if _, ok := h.Clients[client.UserId]; ok {
				delete(h.Clients, client.UserId)
				fmt.Println("logging out user")
				close(client.Send)
			}
		case message := <-h.Broadcast:
			msg_bytes := []byte(message)
			msg := &Message{}
			err := json.Unmarshal(msg_bytes, msg)
			// h.Database.GetNotifications(msg.MessageRecipient)
			if err != nil {
				fmt.Println(err)
			}
			// h.Database.SaveChat(*msg) saving the chat can be move here
			fmt.Println(" recipient UserID:", msg.Recipient, "Message sender: ", msg.SenderID)
			if _, valid := h.Clients[msg.Recipient]; valid {
				h.Clients[msg.Recipient].Send <- message
			}
		}
	}
}
func (h *Hub) LogConns() {
	for {
		fmt.Println(len(h.Clients), "clients connected")
		for userId := range h.Clients {
			fmt.Printf("client %v have %v connections\n", userId, len(h.Clients))
		}
		fmt.Println()
		fmt.Println()
		time.Sleep(1 * time.Second)
		break
	}
}
