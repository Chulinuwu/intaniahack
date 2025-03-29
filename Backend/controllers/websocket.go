package controllers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Room struct {
	Host  *websocket.Conn
	Peers []*websocket.Conn
	mu    sync.Mutex
}

var rooms = make(map[string]*Room)
var roomsMutex sync.Mutex

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket Upgrade Error:", err)
		return
	}

	roomID := c.Param("roomID")

	roomsMutex.Lock()
	room, exists := rooms[roomID]
	if !exists {
		room = &Room{Host: conn, Peers: []*websocket.Conn{}}
		rooms[roomID] = room
		fmt.Println("Room created:", roomID)
	} else {
		room.mu.Lock()
		room.Peers = append(room.Peers, conn)
		room.mu.Unlock()
		fmt.Println("User joined room:", roomID)
	}
	roomsMutex.Unlock()

	go handleMessages(conn, roomID)
}

func handleMessages(conn *websocket.Conn, roomID string) {
	defer func() {
		conn.Close()
		removeConnection(roomID, conn)
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket Read Error:", err)
			break
		}

		broadcast(roomID, message, conn)
	}
}

func broadcast(roomID string, message []byte, sender *websocket.Conn) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	room, exists := rooms[roomID]
	if !exists {
		return
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	for _, peer := range room.Peers {
		if peer != sender {
			peer.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func removeConnection(roomID string, conn *websocket.Conn) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	room, exists := rooms[roomID]
	if !exists {
		return
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	newPeers := []*websocket.Conn{}
	for _, peer := range room.Peers {
		if peer != conn {
			newPeers = append(newPeers, peer)
		}
	}
	room.Peers = newPeers
}
