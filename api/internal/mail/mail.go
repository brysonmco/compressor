package mail

import "errors"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SendVerificationEmail(
	email string,
	firstName string,
	lastName string,
	verificationCode string,
) error {
	return errors.New("not implemented")
}
