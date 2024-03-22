package utils

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to string, otp int64) error {
	// Sender data.
	from := "tranhuythang9999@gmail.com"
	password := "nvkq qdrq ecpa bapz"

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Email subject and message.
	subject := "OTP Verification"
	body := fmt.Sprintf("Your OTP is: %d", otp)

	// Construct the message with subject.
	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
