package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	// Nội dung cho mã QR
	content := "Hello, World!"

	// Tạo mã QR
	qrCode, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	}

	// Lưu mã QR vào file PNG
	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, qrCode.Image(256))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mã QR đã được tạo và lưu vào file qrcode.png.")
}
