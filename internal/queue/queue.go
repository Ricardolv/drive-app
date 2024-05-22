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

	default:
		log.Fatal("type not implemented")
	}

	queue.config = config

	return
}

type QueueConnection interface {
	Publish([]byte) error
	Consume() error
}

type Queue struct {
	config          any
	queueConnection QueueConnection
}

func (q *Queue) Publish(message []byte) error {
	return q.queueConnection.Publish(message)
}

func (q *Queue) Consume() error {
	return q.queueConnection.Consume()
}
