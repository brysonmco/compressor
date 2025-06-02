package messaging

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Service struct {
	Connection *amqp.Connection
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Connect(
	username string,
	password string,
	host string,
) error {
	connStr := fmt.Sprintf("amqp://%s:%s@%s/", username, password, host)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return err
	}
	s.Connection = conn
	return nil
}

func (s *Service) Close() error {
	return s.Connection.Close()
}

func (s *Service) Consume(queueName string, handler func([]byte) error) error {
	ch, err := s.Connection.Channel()
	if err != nil {
		return err
	}
	// Keep channel open for consuming

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queueName,
		"",
		false, // manual ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			err = handler(d.Body)
			if err != nil {
				log.Printf("Handler error: %v. Nacking message.", err)
				d.Nack(false, false)
				continue
			}
			d.Ack(false)
		}
	}()

	return nil
}
