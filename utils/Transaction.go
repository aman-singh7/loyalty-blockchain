package utils

type Service struct {
	transactionId int
}

func NewService() *Service {
	return &Service{
		transactionId: 0,
	}
}

func (s *Service) GetTransactionId() int {
	// TODO: store transaction id in db
	s.transactionId++
	return s.transactionId
}