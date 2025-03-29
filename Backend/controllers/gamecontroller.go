package controllers

import (
	"backend-go/models"
	"backend-go/utils"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var rooms = make(map[string]*models.Room) // เก็บข้อมูลห้องที่สร้าง
var roomsMutex sync.Mutex                 // ป้องกันปัญหาการเข้าถึงข้อมูลพร้อมกัน

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Host สร้างห้องใหม่
func HostGame(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        fmt.Println("WebSocket Upgrade Error:", err)
        return
    }

    // ตรวจสอบว่ามี token หรือไม่
    username, exists := c.Get("username")
    if !exists {
        conn.Close()
        return
    }

    roomID := utils.GenerateRoomID()

    roomsMutex.Lock()
    rooms[roomID] = &models.Room{
        ID:      roomID,
        Host:    conn,
        Players: []*websocket.Conn{},
        HostName: username.(string),  // เก็บชื่อผู้เล่นที่โฮสต์
    }
    roomsMutex.Unlock()

    conn.WriteJSON(gin.H{"room_id": roomID, "host": username})
    fmt.Println("Room created:", roomID)

    go handleMessages(conn, roomID)
}

func JoinGame(c *gin.Context) {
    roomID := c.Query("room_id")
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        return
    }

    // ตรวจสอบว่ามี token หรือไม่
    username, exists := c.Get("username")
    if !exists {
        conn.Close()
        return
    }

    room, exists := rooms[roomID]
    if !exists {
        conn.WriteJSON(gin.H{"error": "Room not found"})
        conn.Close()
        return
    }

    room.Mutex.Lock()
    room.Players = append(room.Players, conn)
    room.Mutex.Unlock()

    conn.WriteJSON(gin.H{"message": "Joined room successfully", "username": username})
}



// จัดการข้อความจาก WebSocket
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

// กระจายข้อความไปยังผู้เล่นทุกคนในห้อง ยกเว้นคนส่ง
func broadcast(roomID string, message []byte, sender *websocket.Conn) {
	roomsMutex.Lock()
	room, exists := rooms[roomID]
	roomsMutex.Unlock()
	if !exists {
		return
	}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	for _, peer := range room.Players {
		if peer != sender {
			peer.WriteMessage(websocket.TextMessage, message)
		}
	}
}

// ลบผู้เล่นออกจากห้องเมื่อ disconnect
func removeConnection(roomID string, conn *websocket.Conn) {
	roomsMutex.Lock()
	room, exists := rooms[roomID]
	roomsMutex.Unlock()
	if !exists {
		return
	}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	newPlayers := []*websocket.Conn{}
	for _, peer := range room.Players {
		if peer != conn {
			newPlayers = append(newPlayers, peer)
		}
	}
	room.Players = newPlayers
}
