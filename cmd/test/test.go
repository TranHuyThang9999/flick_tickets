package main

import (
	"flick_tickets/common/log"
	"io"
	"net/http"
	"os"

	"gopkg.in/gomail.v2"
)

func main() {
	log.LoadLogger() // Initialize the logger
	//thang := "tranhuythang9999@gmail.com"
	thuy := "thuynguyen151387@gmail.com"
	url := "http://localhost:1234/manager/shader/huythang/411373416.png"
	SendEmail(thuy, url)
}

func SendEmail(from, imagePath string) error {
	// Sender data.
	username := "tranhuythang9999@gmail.com"
	password := "nvkq qdrq ecpa bapz"

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	// Download the image from the URL.
	response, err := http.Get(imagePath)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Tạo tệp tin tạm để lưu dữ liệu hình ảnh.
	tempFile, err := os.CreateTemp("", "image.png")
	if err != nil {
		return err
	}

	defer os.Remove(tempFile.Name())
	// Ghi dữ liệu hình ảnh vào tệp tin tạm.
	if _, err := io.Copy(tempFile, response.Body); err != nil {
		return err
	}
	tempFile.Close()

	// Create a new message.
	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", from)
	message.SetHeader("Subject", "Example Subject") // Đặt giá trị tiêu đề

	// Attach the image to the message.
	message.Attach(tempFile.Name())

	// Create a new SMTP client.
	dialer := gomail.NewDialer(smtpHost, smtpPort, username, password)

	// Sending email.
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}

//http://localhost:1234/manager/shader/huythang/411373416.png
