package initialize

import (
	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/kafka"
	"github.com/augustus281/trackingcoin/internal/kafka/consumer"
	"github.com/augustus281/trackingcoin/internal/kafka/producer"
)

func InitKafka() {
	kafkaConfig := kafka.NewKafka([]string{"localhost:9092"}, "tracking-group", []string{"topic1", "topic2"})

	producer := producer.NewProducer(kafkaConfig.Brokers)
	global.Logger.Info("init producer successfully")
	defer producer.Close()

	consumer := consumer.NewConsumer("localhost:9092", kafkaConfig.GroupID, kafkaConfig.Topics[0])
	global.Logger.Info("init consumer successfully")
	defer consumer.Close()
}
