package containers

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) NewContainer(
	downloadUrl string,
) {

}
