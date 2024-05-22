package queue

const (
	RABBITMQ QueueType = iota
)

type QueueType int

func New(config any, queueType QueueType) *Queue {
	q := new(Queue)

	switch queueType {
	case RABBITMQ:
		fmt.Println("nao implementado")
	default:
		log.Fatal("type not implemented")
	}

	q.config = config
	queueType = queueType

	return q
}

type QueueConnection interface {
	Publish([]byte) error
	Consume() error
}

type Queue struct {
	config any
	queueConnection QueueConnection
}

func (q *Queue) Publish(message []byte) error {
	return q.queueConnection.Publish(message)
}

func (q *Queue) Consume() error {
	return q.queueConnection.Consume()
}




