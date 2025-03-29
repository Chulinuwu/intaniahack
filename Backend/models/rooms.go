package models

import (
	"sync"
	"github.com/gorilla/websocket"
)

type Room struct {
	ID          string
	Host        *websocket.Conn
	Players     []*websocket.Conn
	Mutex       sync.Mutex
	HostName    string
	PlayerNames map[*websocket.Conn]string // เพิ่มฟิลด์นี้เพื่อเก็บ username ของผู้เล่น
}

