package notification

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/segmentio/kafka-go"
	"gopkg.in/gomail.v2"
)

var ctx = context.Background()

func TestSendNotification(t *testing.T) {
	notification := Notification{
		Recipient: "huy.nguyen28012002@hcmut.edu.vn",
		Subject:   "Hello from Kafka-go!",
		Body:      "This is a test email sent via Kafka and Kafka-go.",
	}
	writer := kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "topic1",
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	message, err := json.Marshal(notification)
	if err != nil {
		t.Fatalf("Error marshaling notification: %v", err)
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(notification.Recipient),
			Value: message,
		})
	if err != nil {
		t.Fatalf("Error writing to Kafka: %v", err)
	}

	t.Logf("Notification sent to Kafka: %+v", notification)
}

type Notification struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

func TestReceiveNotification(t *testing.T) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "topic1",
		GroupID: "groupID",
	})
	defer reader.Close()

	for {
		select {
		case <-ctx.Done():
			log.Println("Test timed out")
			return
		default:
			message, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}

			var notification Notification
			err = json.Unmarshal(message.Value, &notification)
			if err != nil {
				log.Println("Error unmarshalling message:", err)
				continue
			}

			log.Println("Received notification:", notification)

			err = sendEmail(notification)
			if err != nil {
				log.Println("Error sending email:", err)
			} else {
				log.Println("Email sent successfully to:", notification.Recipient)
			}
			return
		}
	}
}

func sendEmail(notification Notification) error {
	message := gomail.NewMessage()
	message.SetHeader("From", "theflash28012002@gmail.com")
	message.SetHeader("To", notification.Recipient)
	message.SetHeader("Subject", notification.Subject)
	message.SetBody("text/html", notification.Body)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "theflash28012002@gmail.com", "")
	return dialer.DialAndSend(message)
}
