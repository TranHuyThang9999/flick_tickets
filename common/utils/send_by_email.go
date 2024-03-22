package utils

import (
	"bytes"
	"fmt"
	"image/png"
	"io"

	"github.com/skip2/go-qrcode"
	"gopkg.in/gomail.v2"
)

// GeneratesQrCodeAndSendQrWithEmail tạo mã QR từ token và gửi mã QR đến địa chỉ email.
func GeneratesQrCodeAndSendQrWithEmail(addressEmail, token string) error {
	// Tạo mã QR từ token
	qrCode, err := qrcode.New(token, qrcode.Medium)
	if err != nil {
		return fmt.Errorf("error generating QR code: %v", err)
	}

	// Tạo buffer để lưu trữ hình ảnh mã QR dưới dạng byte
	qrCodeBuffer := new(bytes.Buffer)
	err = png.Encode(qrCodeBuffer, qrCode.Image(256))
	if err != nil {
		return fmt.Errorf("error encoding QR code to byte buffer: %v", err)
	}

	// Gửi hình ảnh mã QR đến dịch vụ khác
	// Kênh để nhận lỗi từ goroutine gửi hình ảnh mã QR đến dịch vụ khác
	serviceErrCh := make(chan error)

	// Sử dụng goroutine để gửi hình ảnh mã QR đến dịch vụ khác
	go func(qrCodeBytes []byte) {
		resp := setByCurlImageQrToService(qrCodeBytes)
		if resp.Result.Code != 0 {
			serviceErrCh <- fmt.Errorf("failed to send QR image to service: %s", resp.Result.Message)
		} else {
			serviceErrCh <- nil
		}
	}(qrCodeBuffer.Bytes())

	// Kênh để nhận lỗi từ goroutine gửi hình ảnh mã QR đến địa chỉ email
	emailErrCh := make(chan error)

	// Sử dụng goroutine để gửi hình ảnh mã QR đến địa chỉ email
	go func(email string, attachment []byte) {
		err := SendEmail(email, attachment)
		if err != nil {
			emailErrCh <- fmt.Errorf("error sending email: %v", err)
		} else {
			emailErrCh <- nil
		}
	}(addressEmail, qrCodeBuffer.Bytes())
	// Đợi cho cả hai goroutine hoàn thành và kiểm tra lỗi
	serviceErr := <-serviceErrCh
	emailErr := <-emailErrCh

	if serviceErr != nil {
		return serviceErr
	}

	if emailErr != nil {
		return emailErr
	}
	return nil
}

func SendEmail(to string, attachment []byte) error {
	//	infomations := *configs.Get()
	// Sender data.
	username := "tranhuythang9999@gmail.com"
	password := "nvkq qdrq ecpa bapz"

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	// Create a new message.
	message := gomail.NewMessage()
	message.SetHeader("From", username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", "Example Subject") // Đặt giá trị tiêu đề

	// Attach the image to the message.
	message.Attach("QRCode.png", gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(attachment)
		return err
	}))

	// Create a new SMTP client.
	dialer := gomail.NewDialer(smtpHost, smtpPort, username, password)

	// Sending email.
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
