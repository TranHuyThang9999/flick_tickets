package utils

import (
	"math/rand"
	"time"
)

// GenerateUniqueKey tạo một key có độ dài bằng nhau từ Int64
func GenerateUniqueKey() int64 {
	var length = 7
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Tạo key với độ dài bằng nhau từ Int63
	key := int64(0)
	for i := 0; i < length; i++ {
		key = key*10 + int64(seededRand.Intn(9)) + 1
	}

	return key
}
func GenerateTimestamp() int {
	timeNow := time.Now()
	return int(timeNow.Unix())
}
