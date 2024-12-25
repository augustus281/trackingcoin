package producer

import (
	"context"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	"github.com/augustus281/trackingcoin/global"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Balancer: &kafka.Hash{},
		},
	}
}

func (p *Producer) Publish(topic string, key, message []byte) error {
	err := p.writer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Key:   key,
		Value: message,
	})
	if err != nil {
		global.Logger.Error("publish message to topic failed ", zap.Error(err))
		return err
	}
	global.Logger.Info("message published to top", zap.String("topic", topic))
	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}