package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const (
	letterBytes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func GenerateRoomID() string {
	return uuid.New().String()[:6] // ใช้ 6 ตัวแรกให้สั้นลง
}

// GenerateGameID creates a random 10-character game ID
func GenerateGameID() string {
	b := make([]byte, 10)
	for i, cache, remain := 0, src.Int63(), letterIdxMax; i < 10; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
