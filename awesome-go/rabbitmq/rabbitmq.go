package rabbitmq

import (
	"github.com/streadway/amqp"
)

var (
	defaultUrl  = "amqp://guest:guest@localhost:5672/"
	connections map[string]*amqp.Connection
)

type RabbitMQ struct {
}

func New() *RabbitMQ {
	connections = map[string]*amqp.Connection{}
	r := &RabbitMQ{}
	return r
}

func (r *RabbitMQ) connection(url string) *amqp.Connection {
	if url == "" {
		url = defaultUrl
	}
	if conn, ok := connections[url]; ok {
		return conn
	}
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil
	}
	connections[url] = conn
	return conn
}
func (r *RabbitMQ) GetChannel() (channel *amqp.Channel, err error) {
	conn := r.connection("")
	channel, err = conn.Channel()
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (r *RabbitMQ) Send(ch *amqp.Channel, exchange, name, route, message string) error {
	_, err := ch.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		return err
	}
	err = ch.Publish(exchange, route, false, false, amqp.Publishing{
		Body: []byte(message),
	})
	if err != nil {
		return err
	}
	return nil
}
