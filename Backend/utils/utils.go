package utils

import (
    "github.com/google/uuid"
)

func GenerateRoomID() string {
    return uuid.New().String()[:6] // ใช้ 6 ตัวแรกให้สั้นลง
}