package initialize

import (
	"ecom-project/global"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ() {
	conn, err := amqp.Dial("amqps://akpwryeb:UiSkc1jO9GXYtkjH2PjOVRGRtW9pdmSg@octopus.rmq3.cloudamqp.com/akpwryeb")
	if err != nil {
		checkErrorPanic(err, "Failed to connect to RabbitMQ")
	}
	defer conn.Close()
	global.RabbitMQConn = conn
	global.Logger.Info("RabbitMQ Cloud connected successfully")
}
