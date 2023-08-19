package utils

import "math/big"

func BigInt(value int) *big.Int {
	return big.NewInt(int64(value))
}