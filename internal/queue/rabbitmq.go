package queue

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	URL       string
	TopicName string
	Timeout   time.Duration
}

type RabbitmqConnection struct {
	config     RabbitMQConfig
	connection *amqp.Connection
}

func newRabbitMQConnection(config RabbitMQConfig) (rc *RabbitmqConnection, err error) {
	rc.config = config
	rc.connection, err = amqp.Dial(config.URL)

	return rc, err
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

func (rc *RabbitmqConnection) Consume(queueResp chan<- Message) error {

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

	for response := range msgs {
		message := Message{}
		message.Unmarshal(response.Body)

		queueResp <- message
	}

	return nil
}
