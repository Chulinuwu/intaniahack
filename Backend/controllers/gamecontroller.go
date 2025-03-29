package controllers

import (
	"backend-go/models"
	"backend-go/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
    c.Writer.WriteHeader(http.StatusSwitchingProtocols)

    tokenString := c.GetHeader("Authorization")
    topic := c.Query("topic")
    if tokenString == "" {
        // ถ้าไม่มีใน header ให้รอข้อความแรก
        _, msg, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("WebSocket Read Error:", err)
            conn.Close()
            return
        }

        var tokenData struct {
            Authorization string `json:"Authorization"`
            Topic         string `json:"topic"`
        }
        if err := json.Unmarshal(msg, &tokenData); err != nil {
            conn.WriteJSON(gin.H{"error": "Invalid token format"})
            conn.Close()
            return
        }

        tokenString = strings.TrimPrefix(tokenData.Authorization, "Bearer ")
        if tokenData.Topic != "" {
            topic = tokenData.Topic
        }
    } else {
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
    }

    if tokenString == "" {
        conn.WriteJSON(gin.H{"error": "Token missing"})
        conn.Close()
        return
    }

    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return JwtKey, nil
    })
    if err != nil || !token.Valid {
        conn.WriteJSON(gin.H{"error": "Invalid token"})
        conn.Close()
        return
    }

    username := claims.Username
    roomID := utils.GenerateRoomID()
    if topic == "" {
        topic = "General Knowledge" // Default topic
    }

    roomsMutex.Lock()
    rooms[roomID] = &models.Room{
        ID:          roomID,
        Host:        conn,
        Players:     []*websocket.Conn{},
        HostName:    username,
        PlayerNames: make(map[*websocket.Conn]string),
        Topic:       topic,
    }
    roomsMutex.Unlock()

    conn.WriteJSON(gin.H{"room_id": roomID, "host": username, "topic": topic})
    fmt.Println("Room created:", roomID, "Topic:", topic)

    go handleMessages(conn, roomID)
}

// เข้าร่วมห้อง
func JoinGame(c *gin.Context) {
    roomID := c.Query("room_id")
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        return
    }

    // อ่านข้อความแรกจาก WebSocket เพื่อรับ token
    _, msg, err := conn.ReadMessage()
    if err != nil {
        fmt.Println("WebSocket Read Error:", err)
        conn.Close()
        return
    }

    var tokenData struct {
        Authorization string `json:"Authorization"`
    }
    if err := json.Unmarshal(msg, &tokenData); err != nil {
        conn.Close()
        return
    }

    tokenString := strings.TrimPrefix(tokenData.Authorization, "Bearer ")
    if tokenString == "" {
        conn.Close()
        return
    }

    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return JwtKey, nil
    })
    if err != nil || !token.Valid {
        conn.Close()
        return
    }

    username := claims.Username

    roomsMutex.Lock()
    room, exists := rooms[roomID]
    roomsMutex.Unlock()
    if !exists {
        conn.WriteJSON(gin.H{"error": "Room not found"})
        conn.Close()
        return
    }

    room.Mutex.Lock()
    totalPlayers := len(room.Players) + 1
    if totalPlayers >= 3 {
        conn.WriteJSON(gin.H{"error": "Room is full (max 3 players)"})
        conn.Close()
        room.Mutex.Unlock()
        return
    }

    room.Players = append(room.Players, conn)
    room.PlayerNames[conn] = username
    totalPlayers++
    room.Mutex.Unlock()

    conn.WriteJSON(gin.H{"message": "Joined room successfully", "username": username})

    playersList := getPlayersList(room)
    broadcastPlayerList(roomID, playersList, nil)

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

    message := gin.H{
        "event": "player_list",
        "players": playersList,
        "host": room.HostName, 
		"topic":   room.Topic,
    }

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

        // Parse ข้อความเพื่อตรวจสอบ event
        var msgData struct {
            Event  string `json:"event"`
            RoomID string `json:"room_id"`
        }
        if err := json.Unmarshal(message, &msgData); err == nil {
            if msgData.Event == "start_game" {
                broadcast(roomID, message, conn) // ส่งต่อ start_game ไปยังทุกคน
                continue
            }
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
