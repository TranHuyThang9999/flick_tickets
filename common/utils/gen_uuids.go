package utils

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
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
func GeneratePassword() string {
	var length = 5
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Tạo key với độ dài bằng nhau từ Int63
	key := int64(0)
	for i := 0; i < length; i++ {
		key = key*10 + int64(seededRand.Intn(9)) + 1
	}

	ranStr := make([]rune, length)

	// Tạo chuỗi ngẫu nhiên với chữ cái in hoa
	for i := 0; i < length; i++ {
		ranStr[i] = rune(65 + rand.Intn(25))
	}

	keyinit := fmt.Sprintf("%d%s", key, string(ranStr))
	shuff := []rune(keyinit)
	rand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})

	// Biến các ký tự thành chữ cái in hoa
	for i := 0; i < len(shuff); i++ {
		shuff[i] = unicode.ToUpper(shuff[i])
	}

	return string(shuff)
}
func GenerateOtp() int64 {
	var length = 6
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
func ConvertTimestampToDateTime(timestamp int64) string {
	// Chuyển đổi timestamp thành đối tượng thời gian
	t := time.Unix(timestamp, 0)

	// Định dạng lại ngày tháng theo ý muốn
	formattedDateTime := t.Format("2006-01-02 15:04:05")

	return formattedDateTime
}
func GenerateTimestampExpiredAt(expiredAt int) *int {
	timeNow := time.Now()

	// Cộng thêm thời gian xác định vào thời điểm hiện tại
	expirationTime := timeNow.Add(time.Duration(expiredAt) * time.Minute)

	timestamp := int(expirationTime.Unix())
	return &timestamp
}
