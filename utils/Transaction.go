package utils

import (
	"math/rand"
	"math/big"
)

type Service struct {
	count int64
}

func NewService() *Service {
	return &Service{
		count: 0,
	}
}

func (s *Service) GetTransactionId(key interface{}, salt int64) *big.Int {
	// TODO: store transaction id in db
	s.count++;
	return big.NewInt(rand.Int63() + s.count)
}