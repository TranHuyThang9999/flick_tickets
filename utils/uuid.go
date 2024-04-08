package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenUuid() string {
	id := uuid.New()
	return id.String()
}
func GetCurrentTimestamp() int64 {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	return timestamp
}

// GenerateUniqueKey tạo một key có độ dài bằng nhau từ Int64
func GenerateUniqueKey() int64 {
	var length int = 9
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Tạo key với độ dài bằng nhau từ Int63
	key := int64(0)
	for i := 0; i < length; i++ {
		key = key*10 + int64(seededRand.Intn(9)) + 1
	}

	return key
}
