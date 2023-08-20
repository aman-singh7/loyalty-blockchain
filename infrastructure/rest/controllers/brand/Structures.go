package brand

import "github.com/aman-singh7/loyalty-blockchain/domain/coupon"

type Brand struct{}

type CouponPriceRequest struct {
	BrandID  string        `json:"uid"`
	Coupon   coupon.Coupon `json:"coupon"`
	Count    int           `json:"count"`
	Discount int           `json:"discount"`
}

type CreateCouponRequest struct {
	BrandID  string        `json:"uid"`
	Coupon   coupon.Coupon `json:"coupon"`
	Count    int           `json:"count"`
	Discount int           `json:"discount"`
	Tokens   int           `json:"tokens"`
}

type RedeemTokensRequest struct {
	BrandID string `json:"uid"`
	Tokens  int    `json:"tokens"`
}
