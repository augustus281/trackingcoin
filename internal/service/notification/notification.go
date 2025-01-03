package notification

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/dto"
	"github.com/augustus281/trackingcoin/internal/util"
)

const (
	NotificationTopic = "notification_topic"
)

type INotifyService interface {
	Send(ctx *gin.Context, subject, message string) error
	Receive(ctx *gin.Context) error
}

type notifyService struct{}

func NewNotifyService() INotifyService {
	return &notifyService{}
}

func (s *notifyService) Send(ctx *gin.Context, subject, message string) error {
	writer := kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    NotificationTopic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	notify := dto.Notification{
		Subject: subject,
		Message: message,
	}

	msg, err := json.Marshal(notify)
	if err != nil {
		global.Logger.Error("failed to marshal notification", zap.Error(err))
		return err
	}

	err = writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(subject),
			Value: msg,
		})
	if err != nil {
		global.Logger.Error("error to write message to kafka ", zap.Error(err))
		return err
	}
	return nil
}

func (s *notifyService) Receive(ctx *gin.Context) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "notification-consumer-group",
		Topic:   NotificationTopic,
	})
	defer reader.Close()

	ctxTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	msg, err := reader.ReadMessage(ctxTimeout)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			global.Logger.Error("context timeout exceeded while reading Kafka message", zap.Error(err))
			return err
		}
		global.Logger.Error("error reading Kafka message", zap.Error(err))
		return err
	}

	var notification dto.Notification
	if err := json.Unmarshal(msg.Value, &notification); err != nil {
		global.Logger.Error("error unmarshalling Kafka message", zap.Error(err))
		return err
	}

	emails, err := global.Db.Queries.GetAllEmails(ctx)
	if err != nil {
		global.Logger.Error("failed to get all emails", zap.Error(err))
		return err
	}

	for _, email := range emails {
		notification.To = email
		if err := util.SendNotificationEmail(notification); err != nil {
			global.Logger.Error("error sending email",
				zap.String("email", email),
				zap.Error(err),
			)
		} else {
			global.Logger.Info("email sent successfully", zap.String("email", email))
		}
	}

	return nil
}
