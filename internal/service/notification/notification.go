package notification

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/dto"
	"github.com/augustus281/trackingcoin/internal/kafka/consumer"
	"github.com/augustus281/trackingcoin/internal/kafka/producer"
)

const (
	NotificationTopic = "notification_topic"
)

type INotifyService interface {
	Send(ctx *gin.Context, message string) error
	Receive(ctx *gin.Context, userID int) error
}

type notifyService struct {
	producer *producer.Producer
	consumer *consumer.Consumer
}

func NewNotifyService() INotifyService {
	return &notifyService{
		producer: producer.NewProducer(global.KafkaConsumer.Config().Brokers),
		consumer: consumer.NewConsumer("localhost:9092", "consumer-notification-group-", NotificationTopic),
	}
}

func (s *notifyService) Send(ctx *gin.Context, message string) error {
	listEmail, err := global.Db.Queries.GetAllEmails(ctx)
	if err != nil {
		global.Logger.Error("failed to get all emails ", zap.Error(err))
		return err
	}

	noti := dto.Notification{
		From:    "theflash28012002@gmail.com",
		To:      listEmail,
		Message: message,
	}

	_, err = json.Marshal(noti)
	if err != nil {
		global.Logger.Error("failed to marshal message", zap.Error(err))
		return err
	}
	return nil
}

func (s *notifyService) Receive(ctx *gin.Context, userID int) error {
	return nil
}
