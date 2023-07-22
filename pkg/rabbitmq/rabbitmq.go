package rabbitmq

import (
	ampq "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*ampq.Channel, error) {
	conn, err := ampq.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func Consume(ch *ampq.Channel, out chan ampq.Delivery) error {
	msgs, err := ch.Consume(
		"order",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}
	return nil
}
