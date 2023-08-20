package utils

import (
	"math/rand"
	"math/big"
)

type TransactionGen struct {
	count int64
}

func NewTransactionGen() *TransactionGen {
	return &TransactionGen{
		count: 0,
	}
}

func (s *TransactionGen) GetTransactionId() *big.Int {
	// TODO: store transaction id in db
	s.count++;
	return big.NewInt(rand.Int63() + s.count)
}