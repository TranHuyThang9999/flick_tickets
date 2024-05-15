package utils

import (
	"bytes"
	"encoding/json"
	"flick_tickets/common/log"
	"flick_tickets/configs"
	"flick_tickets/core/entities"
	"fmt"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/skip2/go-qrcode"
	"gopkg.in/gomail.v2"
)

// GeneratesQrCodeAndSendQrWithEmail tạo mã QR từ token và gửi mã QR đến địa chỉ email.
func GeneratesQrCodeAndSendQrWithEmail(addressEmail string, req *entities.OrderSendTicketToEmail, title, token string) error {
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
		log.Infof("url qr code ", resp.URL)
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
		err := SendEmail(email, title, req, attachment)
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

func SendEmail(to, title string, req *entities.OrderSendTicketToEmail, attachment []byte) error {
	infomations := configs.Get()

	username := infomations.FromEmail
	password := infomations.PasswordEmail
	smtpHost := infomations.SmtpHost
	port, err := strconv.Atoi(infomations.SmtpPort)

	if err != nil {
		return err
	}

	// Create a new message.
	message := gomail.NewMessage()
	message.SetHeader("From", username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", title)

	// Read HTML template file
	renderHtml, err := os.ReadFile("api/public/send/index.html")
	if err != nil {
		log.Error(err, "error reading file")
		return err
	}

	// Unmarshal cinemaName JSON into a map
	var cinemaData map[string]string
	err = json.Unmarshal([]byte(req.CinemaName), &cinemaData)
	if err != nil {
		log.Infof("****************  : ", req.CinemaName)
		log.Error(err, "error parsing cinema info JSON")
		return err
	}

	// Set HTML body using template and request data
	body := string(renderHtml)
	body = strings.ReplaceAll(body, "{{id_ve}}", strconv.Itoa(int(req.ID)))
	body = strings.ReplaceAll(body, "{{ten_phim}}", req.MoviceName)
	body = strings.ReplaceAll(body, "{{vi_tri_ghe}}", req.Seats)
	body = strings.ReplaceAll(body, "{{gia_ve}}", strconv.FormatFloat(req.Price, 'f', -1, 64))
	body = strings.ReplaceAll(body, "{{thoi_gian_chieu}}", ConvertTimestampToDateTime(int64(req.MovieTime)))
	body = strings.ReplaceAll(body, "{{tai_dap}}", cinemaData["cinema_name"])
	body = strings.ReplaceAll(body, "{{description}}", cinemaData["description"])
	body = strings.ReplaceAll(body, "{{conscious}}", cinemaData["conscious"])
	body = strings.ReplaceAll(body, "{{district}}", cinemaData["district"])
	body = strings.ReplaceAll(body, "{{commune}}", cinemaData["commune"])
	body = strings.ReplaceAll(body, "{{address_details}}", cinemaData["address_details"])

	message.SetBody("text/html", body)

	// Attach the image to the message.
	message.Attach("QRCode.png", gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(attachment)
		if err != nil {
			log.Error(err, "error sending qr :%s")
			return err
		}
		return nil
	}))

	// Create a new SMTP client.
	dialer := gomail.NewDialer(smtpHost, port, username, password)

	// Sending email.
	if err := dialer.DialAndSend(message); err != nil {
		log.Error(err, "error sending email")
		return err
	}

	log.Info("Email sent successfully")
	return nil
}
func SendOtpToEmail(email, title string, OTP int64) error {
	infomations := configs.Get()

	username := infomations.FromEmail
	password := infomations.PasswordEmail
	smtpHost := infomations.SmtpHost
	port, err := strconv.Atoi(infomations.SmtpPort)
	if err != nil {
		return err
	}

	// Create a new message.
	message := gomail.NewMessage()
	message.SetHeader("From", username)
	message.SetHeader("To", email)
	message.SetHeader("Subject", title)

	// HTML content with h2 tag and CSS styling
	htmlBody := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
		<style>
		h2 {
			color: #333;
			font-family: Arial, sans-serif;
			font-size: 24px;
		}
		</style>
		</head>
		<body>
		<h2>%d</h2>
		</body>
		</html>
	`, OTP)

	// Set the email body with MIME type text/html
	message.SetBody("text/html", htmlBody)

	// Create a new SMTP dialer.
	dialer := gomail.NewDialer(smtpHost, port, username, password)

	// Sending email.
	if err := dialer.DialAndSend(message); err != nil {
		log.Error(err, "error sending email")
		return err
	}

	log.Info("Email sent successfully")
	return nil
}

func SendPasswordToEmail(email, title string, passwordInit string) error {
	infomations := configs.Get()

	username := infomations.FromEmail
	password := infomations.PasswordEmail
	smtpHost := infomations.SmtpHost
	port, err := strconv.Atoi(infomations.SmtpPort)
	if err != nil {
		return err
	}

	// Create a new message.
	message := gomail.NewMessage()
	message.SetHeader("From", username)
	message.SetHeader("To", email)
	message.SetHeader("Subject", title)
	message.SetBody("text/plain", passwordInit) // Set the email body as plain text

	// Create a new SMTP dialer.
	dialer := gomail.NewDialer(smtpHost, port, username, password)

	// Sending email.
	if err := dialer.DialAndSend(message); err != nil {
		log.Error(err, "error sending email")
		return err
	}

	log.Info("Email sent successfully")
	return nil
}

// Hàm để lấy tên người dùng từ email
func GetUsernameFromEmail(email string) string {
	// Tách email thành hai phần bởi dấu '@'
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		log.Error(fmt.Errorf("error"), "Invalid email format")
		return email
	}
	return parts[0]
}
