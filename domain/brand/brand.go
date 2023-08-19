package brand

import "github.com/aman-singh7/loyalty-blockchain/domain/coupon"

type Brand struct {
}

type CouponPriceRequest struct {
	BrandID       int           `json:"uid"`
	Coupon        coupon.Coupon `json:"coupon"`
	Count         int           `json:"count"`
	Discount      int           `json:"discount"`
	TransactionID int           `json:"transactionID"`
}

type CreateCouponRequest struct {
	BrandID       int           `json:"uid"`
	Coupon        coupon.Coupon `json:"coupon"`
	Count         int           `json:"count"`
	Discount      int           `json:"discount"`
	Tokens        int           `json:"tokens"`
	TransactionID int           `json:"transactionID"`
}

type RedeemTokensRequest struct {
	BrandID       int `json:"uid"`
	Tokens        int `json:"tokens"`
	TransactionID int `json:"transactionID"`
}
