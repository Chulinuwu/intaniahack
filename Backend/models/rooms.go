package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Room represents a game room/lobby
type Room struct {
	ID          string
	Host        *websocket.Conn
	Players     []*websocket.Conn
	HostName    string
	PlayerNames map[*websocket.Conn]string
	Topic       string
	Mutex       sync.Mutex
	GameState   *GameState
}
