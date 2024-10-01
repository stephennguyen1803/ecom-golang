package initialize

import (
	"ecom-project/global"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ() {
	// RabbitMq config initialization
	rabbitMQConfig := global.Config.RabbitMQ
	// Connect to RabbitMQ Cloud
	url := fmt.Sprintf("amqps://%s:%s@%s/%s", rabbitMQConfig.User, rabbitMQConfig.Password, rabbitMQConfig.Host, rabbitMQConfig.VHost)
	//url := amqps://akpwryeb:UiSkc1jO9GXYtkjH2PjOVRGRtW9pdmSg@octopus.rmq3.cloudamqp.com/akpwryeb
	conn, err := amqp.Dial(url)
	if err != nil {
		checkErrorPanic(err, "Failed to connect to RabbitMQ")
	}
	global.RabbitMQConn = conn
	global.Logger.Info("RabbitMQ Cloud connected successfully")
	ch, err := conn.Channel()
	if err != nil {
		checkErrorPanic(err, "Failed to open a channel")
	}
	global.RabbitMQCh = ch
}
