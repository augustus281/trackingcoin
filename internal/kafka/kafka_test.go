package kafka

import (
	"fmt"
	"testing"

	"github.com/augustus281/trackingcoin/global"
	gomail "gopkg.in/mail.v2"
)

func TestSendEmail(t *testing.T) {
	message := gomail.NewMessage()

	message.SetHeader("From", "theflash28012002@gmail.com")
	message.SetHeader("To", "huy.nguyen28012002@hcmut.edu.vn")
	message.SetHeader("Subject", "Hello from the Mailtrap team")

	htmlContent := `
		<!DOCTYPE html>
		<html>
		<body>
			<h1 style="color:blue;">Hello from Golang!</h1>
			<p>This is an example of an email with <strong>HTML</strong> content.</p>
			<a href="https://example.com">Visit our website</a>
		</body>
		</html>
	`
	message.SetBody("text/html", htmlContent)

	// Set up the SMTP dialer
	dialer := gomail.NewDialer("smtp.gmail.com", 587, global.Config.SMTP.Username, global.Config.SMTP.Password)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
