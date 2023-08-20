package brand

import (
	"math/big"

	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
	"github.com/ethereum/go-ethereum/common"
)

type Brand struct{}

type CouponPriceRequest struct {
	BrandAddress  common.Address `json:"brandAddress"`
	Coupon        coupon.Coupon  `json:"coupon"`
	Count         big.Int        `json:"count"`
	Discount      big.Int        `json:"discount"`
	TransactionID big.Int        `json:"transactionID"`
}

type CreateCouponRequest struct {
	BrandAddress  common.Address `json:"brandAddress"`
	Coupon        coupon.Coupon  `json:"coupon"`
	Count         big.Int        `json:"count"`
	Discount      big.Int        `json:"discount"`
	Tokens        big.Int        `json:"tokens"`
	TransactionID big.Int        `json:"transactionID"`
}

type RedeemTokensRequest struct {
	BrandAddress  common.Address `json:"brandAddress"`
	Tokens        big.Int        `json:"tokens"`
	TransactionID big.Int        `json:"transactionID"`
}
