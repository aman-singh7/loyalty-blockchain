package utils

import "math/big"

func BigInt(value int) *big.Int {
	return big.NewInt(int64(value))
}

func ToInt(val *big.Int) int {
	return int(val.Int64())
}
