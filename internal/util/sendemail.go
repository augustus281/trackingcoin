package util

import (
	"bytes"
	"html/template"

	"go.uber.org/zap"
	"gopkg.in/gomail.v2"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/dto"
)

func SendNotificationEmail(notification dto.Notification) error {
	message := gomail.NewMessage()
	message.SetHeader("From", global.Config.SMTP.Username)
	message.SetHeader("To", notification.To)
	message.SetHeader("Subject", notification.Subject)
	message.SetBody("text/html", notification.Message)

	dialer := gomail.NewDialer(global.Config.SMTP.Host, 587, global.Config.SMTP.Username, global.Config.SMTP.Password)
	return dialer.DialAndSend(message)
}

func SendMailHTML(to string, listing dto.EmailListing) error {
	tmpl, err := template.New("email").Parse(ListingTemplate)
	if err != nil {
		global.Logger.Error("failed to parse template", zap.Error(err))
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, listing); err != nil {
		global.Logger.Error("failed to execute template", zap.Error(err))
		return err
	}
	message := gomail.NewMessage()
	message.SetHeader("From", global.Config.SMTP.Username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", "Coin listing")
	message.AddAlternative("text/html", ListingTemplate)
	dialer := gomail.NewDialer(global.Config.SMTP.Host, 587, global.Config.SMTP.Username, global.Config.SMTP.Password)

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	global.Logger.Info("send mail html successfully")
	return nil
}
