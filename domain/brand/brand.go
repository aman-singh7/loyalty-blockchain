package brand

import (
	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
	"github.com/ethereum/go-ethereum/common"
)

type Brand struct{}

type CouponPriceRequest struct {
	BrandAddress  common.Address `json:"brandAddress"`
	Coupon        coupon.Coupon  `json:"coupon"`
	Count         int            `json:"count"`
	Discount      int            `json:"discount"`
	TransactionID int            `json:"transactionID"`
}

type CreateCouponRequest struct {
	BrandAddress  common.Address `json:"brandAddress"`
	Coupon        coupon.Coupon  `json:"coupon"`
	Count         int            `json:"count"`
	Discount      int            `json:"discount"`
	Tokens        int            `json:"tokens"`
	TransactionID int            `json:"transactionID"`
}

type RedeemTokensRequest struct {
	BrandAddress  common.Address `json:"brandAddress"`
	Tokens        int            `json:"tokens"`
	TransactionID int            `json:"transactionID"`
}
