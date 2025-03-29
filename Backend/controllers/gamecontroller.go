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
	Subprotocols: []string{"json"},
}

// Host สร้างห้องใหม่
func HostGame(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket Upgrade Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "WebSocket Upgrade Error"})
		return
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		conn.Close()
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		conn.Close()
		return
	}

	roomID := utils.GenerateRoomID()

	roomsMutex.Lock()
	rooms[roomID] = &models.Room{
		ID:          roomID,
		Host:        conn,
		Players:     []*websocket.Conn{},
		HostName:    username.(string),
		PlayerNames: make(map[*websocket.Conn]string), // กำหนดค่าเริ่มต้นให้ PlayerNames
	}
	roomsMutex.Unlock()

	conn.WriteJSON(gin.H{"room_id": roomID, "host": username})
	fmt.Println("Room created:", roomID)

	go handleMessages(conn, roomID)
}

// เข้าร่วมห้อง
func JoinGame(c *gin.Context) {
	roomID := c.Query("room_id")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	username, exists := c.Get("username")
	if !exists {
		conn.Close()
		return
	}

	roomsMutex.Lock()
	room, exists := rooms[roomID]
	roomsMutex.Unlock()
	if !exists {
		conn.WriteJSON(gin.H{"error": "Room not found"})
		conn.Close()
		return
	}

	room.Mutex.Lock()
	// ตรวจสอบจำนวนผู้เล่น (รวม host แล้วต้องไม่เกิน 3)
	totalPlayers := len(room.Players) + 1 // +1 เพราะมี host
	if totalPlayers >= 3 {
		conn.WriteJSON(gin.H{"error": "Room is full (max 3 players)"})
		conn.Close()
		room.Mutex.Unlock()
		return
	}

	room.Players = append(room.Players, conn)
	room.PlayerNames[conn] = username.(string) // เก็บ username ของผู้เล่นใน map
	totalPlayers++                             // อัปเดตจำนวนผู้เล่นหลังจากเพิ่ม
	room.Mutex.Unlock()

	// ส่งข้อความยืนยันไปยังผู้เล่นที่เข้าร่วม
	conn.WriteJSON(gin.H{"message": "Joined room successfully", "username": username})

	// สร้างรายชื่อผู้เล่น
	playersList := getPlayersList(room)

	// Broadcast รายชื่อผู้เล่นให้ทุกคนในห้อง
	broadcastPlayerList(roomID, playersList, nil)

	// ถ้าครบ 3 คน ส่งข้อความว่าเริ่มเกมได้
	if totalPlayers == 3 {
		broadcastStartGame(roomID)
	}

	go handleMessages(conn, roomID)
}

// ดึงรายชื่อผู้เล่นในห้อง
func getPlayersList(room *models.Room) []string {
	var playersList []string
	playersList = append(playersList, room.HostName) // เพิ่มชื่อ host

	room.Mutex.Lock()
	for _, conn := range room.Players {
		username, exists := room.PlayerNames[conn]
		if exists {
			playersList = append(playersList, username)
		} else {
			playersList = append(playersList, "unknown") // ในกรณีที่ไม่พบ username (ไม่น่าจะเกิดขึ้น)
		}
	}
	room.Mutex.Unlock()

	return playersList
}

// Broadcast รายชื่อผู้เล่นให้ทุกคนในห้อง
func broadcastPlayerList(roomID string, playersList []string, sender *websocket.Conn) {
	roomsMutex.Lock()
	room, exists := rooms[roomID]
	roomsMutex.Unlock()
	if !exists {
		return
	}

	message := gin.H{"event": "player_list", "players": playersList}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	// ส่งไปยัง host
	if room.Host != sender {
		room.Host.WriteJSON(message)
	}
	// ส่งไปยังผู้เล่นทุกคน
	for _, peer := range room.Players {
		if peer != sender {
			peer.WriteJSON(message)
		}
	}
}

// Broadcast ข้อความว่าเริ่มเกมได้
func broadcastStartGame(roomID string) {
	roomsMutex.Lock()
	room, exists := rooms[roomID]
	roomsMutex.Unlock()
	if !exists {
		return
	}

	message := gin.H{"event": "game_ready", "message": "Room is full. Game can start!"}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	// ส่งไปยัง host
	room.Host.WriteJSON(message)
	// ส่งไปยังผู้เล่นทุกคน
	for _, peer := range room.Players {
		peer.WriteJSON(message)
	}
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

	// ลบ username ออกจาก PlayerNames เมื่อผู้เล่นออก
	delete(room.PlayerNames, conn)
}