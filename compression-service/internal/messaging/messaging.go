package messaging

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
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
