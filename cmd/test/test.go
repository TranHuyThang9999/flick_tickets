package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
)

// Encrypt sử dụng AES để mã hóa dữ liệu với khóa đã cho.
func Encrypt(data []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Tạo một vector khởi tạo (IV) ngẫu nhiên
	iv := make([]byte, aes.BlockSize)

	// Tạo chế độ CBC với khối mã hóa và vector khởi tạo
	mode := cipher.NewCBCEncrypter(block, iv)

	// Thêm padding vào dữ liệu
	blockSize := aes.BlockSize
	data = pkcs7Pad(data, blockSize)

	// Mã hóa dữ liệu
	ciphertext := make([]byte, len(data))
	mode.CryptBlocks(ciphertext, data)

	// Chuyển đổi mã hóa thành chuỗi base64
	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)

	return ciphertextBase64, nil
}

// Decrypt sử dụng AES để giải mã dữ liệu đã được mã hóa với khóa đã cho.
func Decrypt(ciphertextBase64 string, key []byte) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Tạo một vector khởi tạo (IV) ngẫu nhiên
	iv := make([]byte, aes.BlockSize)

	// Tạo chế độ CBC với khối mã hóa và vector khởi tạo
	mode := cipher.NewCBCDecrypter(block, iv)

	// Giải mã dữ liệu
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// Xóa padding từ dữ liệu giải mã
	plaintext, err = pkcs7Unpad(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// Hàm pkcs7Unpad xóa padding từ dữ liệu theo chuẩn PKCS7
func pkcs7Unpad(data []byte) ([]byte, error) {
	padding := int(data[len(data)-1])
	if padding < 1 || padding > aes.BlockSize {
		return nil, fmt.Errorf("Invalid padding")
	}

	return data[:len(data)-padding], nil
}

// Hàm pkcs7Pad thêm padding vào dữ liệu theo chuẩn PKCS7
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func main() {
	// Khóa AES (32 byte)
	key := []byte("your-256-bit-key-32-byte-long-wx")
	fakeKey := []byte("your-256-bit-key-32-byte-long-wap")
	// Đoạn mã dữ liệu cần mã hóa
	plaintext := []byte("Hello, World!")

	// Mã hóa dữ liệu
	ciphertext, err := Encrypt(plaintext, key)
	if err != nil {
		log.Fatal("Lỗi khi mã hóa dữ liệu:", err)
	}

	// In ra mã hóa chuỗi base64
	log.Println("Mã hóa base64:", ciphertext)

	// Giải mã dữ liệu
	decrypted, err := Decrypt(ciphertext, fakeKey)
	if err != nil {
		log.Fatal("Lỗi khi giải mã dữ liệu:", err)
	}

	// In ra dữ liệu giải mã
	log.Println("Dữ liệu giải mã:", string(decrypted))
}
