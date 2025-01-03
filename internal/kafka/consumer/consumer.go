package consumer

import (
	"github.com/augustus281/trackingcoin/global"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(groupID string, topics []string) *Consumer {
	consumer := &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:     []string{"localhost:9092"},
			GroupID:     groupID,
			GroupTopics: topics,
		}),
	}
	return consumer
}

func (c *Consumer) Start(ctx *gin.Context, handler func(topic string, message []byte) error) error {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			global.Logger.Error("error reading messages ", zap.Error(err))
			return err
		}
		global.Logger.Info("message received ",
			zap.String("topic", msg.Topic),
			zap.String("msgKey", string(msg.Key)),
			zap.String("value", string(msg.Value)),
		)
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
