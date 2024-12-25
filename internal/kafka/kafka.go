package kafka

type Kafka struct {
	Brokers []string
	GroupID string
	Topics  []string
}

func NewKafka(brokers []string, groupID string, topics []string) *Kafka {
	return &Kafka{
		Brokers: brokers,
		GroupID: groupID,
		Topics:  topics,
	}
}
