package queue

import (
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)


type RabbitMQConfig struct {
	URL string
	TopicName string
	Timeout time.Time
}

type RabbitmqConnection struct {
	config RabbitMQConfig
	connection *amqp.Connection
}