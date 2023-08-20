package brand

import (
	"math/big"

	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
)

type Brand struct{}

type CouponPriceRequest struct {
	BrandID  string        `json:"uid"`
	Coupon   coupon.Coupon `json:"coupon"`
	Count    big.Int       `json:"count"`
	Discount big.Int       `json:"discount"`
}

type CreateCouponRequest struct {
	BrandID  string        `json:"uid"`
	Coupon   coupon.Coupon `json:"coupon"`
	Count    big.Int       `json:"count"`
	Discount big.Int       `json:"discount"`
	Tokens   big.Int       `json:"tokens"`
}

type RedeemTokensRequest struct {
	BrandID string  `json:"uid"`
	Tokens  big.Int `json:"tokens"`
}
