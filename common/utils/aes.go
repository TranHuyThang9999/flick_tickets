package utils

import (
	"bytes"
	"crypto/aes"
	"fmt"
)

// Hàm pkcs7Unpad xóa padding từ dữ liệu theo chuẩn PKCS7
func Pkcs7Unpad(data []byte) ([]byte, error) {
	padding := int(data[len(data)-1])
	if padding < 1 || padding > aes.BlockSize {
		return nil, fmt.Errorf("Invalid padding")
	}

	return data[:len(data)-padding], nil
}

// Hàm pkcs7Pad thêm padding vào dữ liệu theo chuẩn PKCS7
func Pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
