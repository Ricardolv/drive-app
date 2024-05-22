package queue

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	URL       string
	TopicName string
	Timeout   time.Time
}

type RabbitmqConnection struct {
	config     RabbitMQConfig
	connection *amqp.Connection
}

func (rc *RabbitmqConnection) Publish(message []byte) error {

	channel, err := rc.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	mp := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         message,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return channel.PublishWithContext(ctx,
		"",
		rc.config.TopicName,
		false,
		false,
		mp)
}

func (rc *RabbitmqConnection) Consume() error {

	channel, err := rc.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(rc.config.TopicName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	for d := range msgs {

	}

}
