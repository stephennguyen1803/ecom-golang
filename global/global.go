package global

import (
	"database/sql"
	"ecom-project/pkg/logger"
	"ecom-project/pkg/settings"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Mdb           *gorm.DB
	Mdbc          *sql.DB
	Config        settings.Config
	Logger        *logger.LoggerZap
	Redis         *redis.Client
	KafkaProducer *kafka.Writer
	RabbitMQConn  *amqp.Connection
	RabbitMQCh    *amqp.Channel
)
