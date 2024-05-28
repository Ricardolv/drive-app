package queue

import (
	"fmt"
	"log"
	"reflect"
)

const (
	RABBITMQ QueueType = iota
)

type QueueType int

func New(config any, queueType QueueType) (queue *Queue, err error) {
	rt := reflect.TypeOf(config)

	switch queueType {
	case RABBITMQ:

		if rt.Name() != "RabbitMQConfig" {
			return nil, fmt.Errorf("Config need's to be of type RabbitMQConfig")
		}

		conn, err := newRabbitMQConnection(config.(RabbitMQConfig))
		if err != nil {
			return nil, err
		}

		queue.queueConnection = conn

	default:
		log.Fatal("type not implemented")
	}

	return
}

type QueueConnection interface {
	Publish([]byte) error
	Consume(chan<- Message) error
}

type Queue struct {
	queueConnection QueueConnection
}

func (q *Queue) Publish(message []byte) error {
	return q.queueConnection.Publish(message)
}

func (q *Queue) Consume(queueResp chan<- Message) error {
	return q.queueConnection.Consume(queueResp)
}
