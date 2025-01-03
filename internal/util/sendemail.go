package util

import (
	"gopkg.in/gomail.v2"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/dto"
)

func SendNotificationEmail(notification dto.Notification) error {
	message := gomail.NewMessage()
	message.SetHeader("From", "theflash28012002@gmail.com")
	message.SetHeader("To", notification.To)
	message.SetHeader("Subject", notification.Subject)
	message.SetBody("text/html", notification.Message)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "theflash28012002@gmail.com", global.Config.SMTP.Password)
	return dialer.DialAndSend(message)
}
