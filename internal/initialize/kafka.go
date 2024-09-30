package initialize

import (
	"ecom-project/global"
	"log"

	"github.com/segmentio/kafka-go"
)

// InitKafka produces a new Kafka connection.

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "otp-auth-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatal("Failed to close Kafka producer: ", err)
	}
}
