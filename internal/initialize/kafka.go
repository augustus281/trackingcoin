package initialize

import (
	"github.com/augustus281/trackingcoin/internal/kafka"
	"github.com/augustus281/trackingcoin/internal/kafka/consumer"
	"github.com/augustus281/trackingcoin/internal/kafka/producer"
)

func InitKafka() {
	kafkaConfig := kafka.NewKafka([]string{"localhost:9092"}, "tracking-group", []string{"topic1", "topic2"})

	producer := producer.NewProducer(kafkaConfig.Brokers)
	defer producer.Close()

	consumer := consumer.NewConsumer(kafkaConfig.Brokers, kafkaConfig.GroupID, kafkaConfig.Topics)
	defer consumer.Close()
}
