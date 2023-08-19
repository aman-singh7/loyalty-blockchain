package brand

import "github.com/aman-singh7/loyalty-blockchain/domain/coupon"

type Brand struct {
}

type DomainCouponPriceRequest struct {
	BrandID       int           `json:"uid"`
	Coupon        coupon.Coupon `json:"coupon"`
	Count         int           `json:"count"`
	Discount      int           `json:"discount"`
	TransactionID int           `json:"transactionID"`
}

type DomainCreateCouponRequest struct {
	BrandID       int           `json:"uid"`
	Coupon        coupon.Coupon `json:"coupon"`
	Count         int           `json:"count"`
	Discount      int           `json:"discount"`
	Tokens        int           `json:"tokens"`
	TransactionID int           `json:"transactionID"`
}

type DomainRedeemTokensRequest struct {
	BrandID       int `json:"uid"`
	Tokens        int `json:"tokens"`
	TransactionID int `json:"transactionID"`
}
