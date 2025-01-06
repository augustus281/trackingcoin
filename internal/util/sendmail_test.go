package util

import (
	"bytes"
	"html/template"
	"log"
	"testing"
	"time"

	"github.com/augustus281/trackingcoin/internal/dto"
	"gopkg.in/gomail.v2"
)

func TestSendHTMLMail(t *testing.T) {
	listing := dto.EmailListing{
		Subject:       "Coin listing",
		RecipientName: "huy.nguyen28012002@hcmut.edu.vn",
		ListingData: dto.ListingResponse{
			Data: []dto.CryptoData{
				{
					Name:        "Bitcoin",
					Symbol:      "BTC",
					LastUpdated: time.Now(),
					Quote: dto.QuoteCryptop{
						USD: dto.QuoteUSD{
							Price: 98830.11719584813,
						},
					},
				},
				{
					Name:        "Ethereum",
					Symbol:      "ETH",
					LastUpdated: time.Now(),
					Quote: dto.QuoteCryptop{
						USD: dto.QuoteUSD{
							Price: 3660.8926992061274,
						},
					},
				},
			},
		},
	}
	tmpl, err := template.New("email").Parse(ListingTemplate)
	if err != nil {
		log.Fatal(err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, listing); err != nil {
		log.Fatal(err)
	}

	message := gomail.NewMessage()
	message.SetHeader("From", "theflash28012002@gmail.com")
	message.SetHeader("To", "huy.nguyen28012002@hcmut.edu.vn")
	message.SetHeader("Subject", "Coin listing")
	message.SetBody("text/html", body.String())
	dialer := gomail.NewDialer("host", 587, "theflash28012002@gmail.com", "password")

	if err := dialer.DialAndSend(message); err != nil {
		log.Fatal(err)
	}
}
